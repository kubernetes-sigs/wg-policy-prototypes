package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/trivy-adapter/pkg/plugin"
	"github.com/kubernetes-sigs/wg-policy-prototypes/policy-report/trivy-adapter/pkg/imgvuln"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	ext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	extapi "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	clusterRoleImgvuln        = "imgvuln"
	clusterRoleBindingImgvuln = "imgvuln"
)

var (
	namespace = &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: imgvuln.NamespaceName,
			Labels: labels.Set{
				"app.kubernetes.io/managed-by": "imgvuln",
			},
		},
	}
	serviceAccount = &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: imgvuln.ServiceAccountName,
			Labels: labels.Set{
				"app.kubernetes.io/managed-by": "imgvuln",
			},
		},
		AutomountServiceAccountToken: pointer.BoolPtr(false),
	}
	clusterRole = &rbacv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterRoleImgvuln,
			Labels: labels.Set{
				"app.kubernetes.io/managed-by": "imgvuln",
			},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{
					"",
				},
				Resources: []string{
					"nodes",
					"namespaces",
					"pods",
				},
				Verbs: []string{
					"list",
					"get",
				},
			},
			{
				APIGroups: []string{
					"apps",
				},
				Resources: []string{
					"deployments",
					"statefulsets",
					"daemonsets",
					"replicationcontrollers",
					"replicasets",
				},
				Verbs: []string{
					"list",
					"get",
				},
			},
			{
				APIGroups: []string{
					"batch",
				},
				Resources: []string{
					"jobs",
					"cronjobs",
				},
				Verbs: []string{
					"list",
					"get",
				},
			},
		},
	}
	clusterRoleBinding = &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterRoleBindingImgvuln,
			Labels: labels.Set{
				"app.kubernetes.io/managed-by": "imgvuln",
			},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     clusterRoleImgvuln,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      imgvuln.ServiceAccountName,
				Namespace: imgvuln.NamespaceName,
			},
		},
	}
)

type Installer struct {
	buildInfo     imgvuln.BuildInfo
	client        client.Client
	clientset     kubernetes.Interface
	clientsetext  extapi.ApiextensionsV1Interface
	configManager imgvuln.ConfigManager
}

// NewInstaller constructs an Installer with the given imgvuln.ConfigManager and kubernetes.Interface.
func NewInstaller(
	buildInfo imgvuln.BuildInfo,
	// TODO Get rid of kubernetes.Interface and ApiextensionsV1Interface and use just client.Client
	clientset kubernetes.Interface,
	clientsetext extapi.ApiextensionsV1Interface,
	client client.Client,
	configManager imgvuln.ConfigManager,
) *Installer {
	return &Installer{
		buildInfo:     buildInfo,
		clientset:     clientset,
		clientsetext:  clientsetext,
		client:        client,
		configManager: configManager,
	}
}

func (m *Installer) Install(ctx context.Context) error {

	err := m.createNamespaceIfNotFound(ctx, namespace)
	if err != nil {
		return err
	}

	err = m.configManager.EnsureDefault(ctx)
	if err != nil {
		return err
	}

	config, err := m.configManager.Read(ctx)
	if err != nil {
		return err
	}

	pluginResolver := plugin.NewResolver().
		WithBuildInfo(m.buildInfo).
		WithNamespace(imgvuln.NamespaceName).
		WithServiceAccountName(imgvuln.ServiceAccountName).
		WithConfig(config).
		WithClient(m.client)

	vulnerabilityPlugin, pluginContext, err := pluginResolver.GetVulnerabilityPlugin()
	if err != nil {
		return err
	}

	err = vulnerabilityPlugin.Init(pluginContext)
	if err != nil {
		return fmt.Errorf("initializing %s plugin: %w", pluginContext.GetName(), err)
	}

	return m.initRBAC(ctx)

}

func (m *Installer) initRBAC(ctx context.Context) (err error) {
	err = m.createServiceAccountIfNotFound(ctx, serviceAccount)
	if err != nil {
		return
	}

	err = m.createOrUpdateClusterRole(ctx, clusterRole)
	if err != nil {
		return
	}

	err = m.createOrUpdateClusterRoleBinding(ctx, clusterRoleBinding)

	return
}

func (m *Installer) cleanupRBAC(ctx context.Context) (err error) {
	klog.V(3).Infof("Deleting ClusterRoleBinding %q", clusterRoleBindingImgvuln)
	err = m.clientset.RbacV1().ClusterRoleBindings().Delete(ctx, clusterRoleBindingImgvuln, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	klog.V(3).Infof("Deleting ClusterRole %q", clusterRoleImgvuln)
	err = m.clientset.RbacV1().ClusterRoles().Delete(ctx, clusterRoleImgvuln, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	klog.V(3).Infof("Deleting ServiceAccount %q", imgvuln.NamespaceName+"/"+imgvuln.ServiceAccountName)
	err = m.clientset.CoreV1().ServiceAccounts(imgvuln.NamespaceName).Delete(ctx, imgvuln.ServiceAccountName, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	return nil
}

var (
	cleanupPollingInterval = 2 * time.Second
	cleanupTimeout         = 30 * time.Second
)

func (m *Installer) cleanupNamespace(ctx context.Context) error {
	klog.V(3).Infof("Deleting Namespace %q", imgvuln.NamespaceName)
	err := m.clientset.CoreV1().Namespaces().Delete(ctx, imgvuln.NamespaceName, metav1.DeleteOptions{})
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	for {
		select {
		// This case controls the polling interval
		case <-time.After(cleanupPollingInterval):
			_, err := m.clientset.CoreV1().Namespaces().Get(ctx, imgvuln.NamespaceName, metav1.GetOptions{})
			if errors.IsNotFound(err) {
				klog.V(3).Infof("Deleted Namespace %q", imgvuln.NamespaceName)
				return nil
			}
		// This case caters for polling timeout
		case <-time.After(cleanupTimeout):
			return fmt.Errorf("deleting namespace timed out")
		}
	}
}

func (m *Installer) createNamespaceIfNotFound(ctx context.Context, ns *corev1.Namespace) (err error) {
	_, err = m.clientset.CoreV1().Namespaces().Get(ctx, ns.Name, metav1.GetOptions{})
	switch {
	case err == nil:
		klog.V(3).Infof("Namespace %q already exists", ns.Name)
		return
	case errors.IsNotFound(err):
		klog.V(3).Infof("Creating Namespace %q", ns.Name)
		_, err = m.clientset.CoreV1().Namespaces().Create(ctx, ns, metav1.CreateOptions{})
		return
	}
	return
}

func (m *Installer) createServiceAccountIfNotFound(ctx context.Context, sa *corev1.ServiceAccount) (err error) {
	name := sa.Name
	_, err = m.clientset.CoreV1().ServiceAccounts(imgvuln.NamespaceName).Get(ctx, name, metav1.GetOptions{})
	switch {
	case err == nil:
		klog.V(3).Infof("ServiceAccount %q already exists", imgvuln.NamespaceName+"/"+name)
		return
	case errors.IsNotFound(err):
		klog.V(3).Infof("Creating ServiceAccount %q", imgvuln.NamespaceName+"/"+name)
		_, err = m.clientset.CoreV1().ServiceAccounts(imgvuln.NamespaceName).Create(ctx, sa, metav1.CreateOptions{})
		return
	}
	return
}

func (m *Installer) createOrUpdateClusterRole(ctx context.Context, cr *rbacv1.ClusterRole) (err error) {
	existingRole, err := m.clientset.RbacV1().ClusterRoles().Get(ctx, cr.GetName(), metav1.GetOptions{})
	switch {
	case err == nil:
		klog.V(3).Infof("Updating ClusterRole %q", cr.GetName())
		deepCopy := existingRole.DeepCopy()
		deepCopy.Rules = cr.Rules
		_, err = m.clientset.RbacV1().ClusterRoles().Update(ctx, deepCopy, metav1.UpdateOptions{})
		return
	case errors.IsNotFound(err):
		klog.V(3).Infof("Creating ClusterRole %q", cr.GetName())
		_, err = m.clientset.RbacV1().ClusterRoles().Create(ctx, cr, metav1.CreateOptions{})
		return
	}
	return
}

func (m *Installer) createOrUpdateClusterRoleBinding(ctx context.Context, crb *rbacv1.ClusterRoleBinding) (err error) {
	existingBinding, err := m.clientset.RbacV1().ClusterRoleBindings().Get(ctx, crb.Name, metav1.GetOptions{})
	switch {
	case err == nil:
		klog.V(3).Infof("Updating ClusterRoleBinding %q", crb.GetName())
		deepCopy := existingBinding.DeepCopy()
		deepCopy.RoleRef = crb.RoleRef
		deepCopy.Subjects = crb.Subjects
		_, err = m.clientset.RbacV1().ClusterRoleBindings().Update(ctx, deepCopy, metav1.UpdateOptions{})
		return
	case errors.IsNotFound(err):
		klog.V(3).Infof("Creating ClusterRoleBinding %q", crb.GetName())
		_, err = m.clientset.RbacV1().ClusterRoleBindings().Create(ctx, crb, metav1.CreateOptions{})
		return
	}
	return
}

func (m *Installer) createOrUpdateCRD(ctx context.Context, crd *ext.CustomResourceDefinition) (err error) {
	existingCRD, err := m.clientsetext.CustomResourceDefinitions().Get(ctx, crd.Name, metav1.GetOptions{})

	switch {
	case err == nil:
		klog.V(3).Infof("Updating CRD %q", crd.Name)
		deepCopy := existingCRD.DeepCopy()
		deepCopy.Spec = crd.Spec
		_, err = m.clientsetext.CustomResourceDefinitions().Update(ctx, deepCopy, metav1.UpdateOptions{})
		return
	case errors.IsNotFound(err):
		klog.V(3).Infof("Creating CRD %q", crd.Name)
		_, err = m.clientsetext.CustomResourceDefinitions().Create(ctx, crd, metav1.CreateOptions{})
		return
	}
	return
}