# Policy Report

This is a proposal for a Policy Report Custom Resource Definition (CRD) that can be used as a common way to provide policy results to Kubernetes cluster administrators and users, using native tools.

See the [proposal](https://docs.google.com/document/d/1nICYLkYS1RE3gJzuHOfHeAC25QIkFZfgymFjgOzMDVw/edit#) for background and details.

## Installing

Add the CRDs to your cluster:

```console
kubectl create -f https://github.com/kubernetes-sigs/wg-policy-prototypes/raw/master/policy-report/crd/policy.kubernetes.io_policyreports.yaml
```

Create a sample policy report resource:

```console
kubectl create -f https://github.com/kubernetes-sigs/wg-policy-prototypes/raw/master/policy-report/samples/sample-cis-k8s.yaml
```

View policy report resources:

```console
kubectl get policyreports
```

## Building 

```console
make
```

## Contributing  

The Policy Report CRDs definitions are in the `api` folder and defined as Golang types with comments using the syntax of the [kubebuilder controller-gen](https://book.kubebuilder.io/reference/controller-gen.html) tool that can generate Kubernetes YAMLs. 

To update, edit the Golang definitions and then run `make` to generate the Kubernetes OpenAPI schema for the CRDs.

Definitions are provided for both cluster-wide and namespaced policy report resources. 