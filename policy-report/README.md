# Policy Report

This is a proposal for a Policy Report Custom Resource Definition (CRD) that can be used as a common way to provide policy results to Kubernetes cluster administrators and users, using native tools.

See the [proposal](https://docs.google.com/document/d/1nICYLkYS1RE3gJzuHOfHeAC25QIkFZfgymFjgOzMDVw/edit#) for background and details.

[Policy Report CRD API Reference](https://htmlpreview.github.io/?https://github.com/kubernetes-sigs/wg-policy-prototypes/blob/master/policy-report/docs/index.html)

**Note:** v1beta1 APIs is WIP and will have breaking changes.

## Installing

Add the PolicyReport CRDs to your cluster (v1alpha2):

```console
kubectl create -f https://github.com/kubernetes-sigs/wg-policy-prototypes/raw/master/policy-report/crd/v1alpha2/wgpolicyk8s.io_policyreports.yaml
```

Add the ClusterPolicyReport CRDs to your cluster (v1alpha2):

```console
kubectl create -f https://github.com/kubernetes-sigs/wg-policy-prototypes/raw/master/policy-report/crd/v1alpha2/wgpolicyk8s.io_clusterpolicyreports.yaml
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

NOTE : For generating CRD documentation please follow the steps

```bash
$ git clone https://github.com/M00nF1sh/gen-crd-api-reference-docs.git
$ cd gen-crd-api-reference-docs 
$ go build
$ mv gen-crd-api-reference-docs /usr/local/bin/
$ make generate
```
