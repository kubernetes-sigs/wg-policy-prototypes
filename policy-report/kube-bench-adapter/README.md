# kube-bench adapter
The kube-bench adapter runs a CIS benchmark check with a tool called [kube-bench](https://github.com/aquasecurity/kube-bench) and produces a cluster-wide policy report based on the [Policy Report Custom Resource Definition](https://github.com/kubernetes-sigs/wg-policy-prototypes/tree/master/policy-report)

## Running

**Prerequisites**: 
* To run the Kubernetes cluster locally, tools like [kind](https://kind.sigs.k8s.io/) or [minikube](https://minikube.sigs.k8s.io/docs/start/) can be used. Here are the steps to run the kube-bench adapater with a `kind` cluster.

### Steps

#### Common steps
```sh
# 1. Create a Kubernetes cluster
kind create cluster

# 2. Create a CustomResourceDefinition
kubectl create -f kubernetes/crd/v1alpha2/wgpolicyk8s.io_clusterpolicyreports.yaml
```
#### Steps to run kube-bench adapter in-cluster as a Cron-Job
```sh
# 3. Create Role, Role-Binding and Services
kubectl create -f kubernetes/role.yaml -f kubernetes/rb.yaml -f kubernetes/service.yaml

# 4. Create cron-job
kubectl create -f kubernetes/cron-job.yaml

# 5. Watch the jobs
kubectl get jobs --watch

# 6. Check policyreports created through the custom resource
kubectl get clusterpolicyreports
```

#### Steps to run kube-bench adapter from outside a cluster 
##### Building
```sh
make build
```
##### Installing
```sh
# Create policy report using
./policyreport -name="kube-bench" -kube-bench-targets="master,node" -yaml="job.yaml" -category="CIS Benchmarks"

# Check policyreports created through the custom resource
kubectl get clusterpolicyreports
```
###### Command Line Arguments
|      Argument         |  Type   |    Default value         | Allowed value  | Usage                                            |
|:--------------------- |:-------:|-------------------------:|:--------------:|:------------------------------------------------:|
| -category             | `string`| CIS Benchmarks           |   Any string name valid for category             | category of the policy report                    |   
| -kube-bench-benchmark | `string`|   cis-1.6                    |    cis-1.5, cis-1.6, gke-1.0, eks-1.0, ack-1.0            | specify the benchmark for kube-bench job         |
| -kube-bench-targets   | `string`(accepts multiple values)| master,node,etcd,policies| 	master, controlplane, node, etcd, policies               | targets for benchmark of kube-bench job          |   
| -kube-bench-version   | `string`|    1.21                    |   Kubernetes Version like 1.20,1.21,etc             | specify the Kubernetes version for kube-bench job|
| -kubebenchImg         | `string`| aquasec/kube-bench:latest|aquasec/kube-bench:(kube-bench-version)                | kube-bench image used as part of this test       |
| -kubeconfig           | `string`| $HOME/.kube/config       |  path to your KUBECONFIG              | absolute path to the kubeconfig file             | 
| -name                 | `string`| kube-bench               |  Any name of string type              | name of policy report                            |
| -yaml                 | `string`| job.yaml                 |  job.yaml, job-master.yaml, job-node.yaml, job-ack.yaml, job-aks.yaml, job-eks-asff.yaml, job-eks.yaml, job-gke.yaml, job-iks.yaml             | name of YAML for kube-bench job                  | 

## Project Maintenance

### Updating the Policy Report CRD

### Updating the kube-bench jobs


**Notes**:
* Flags `-name`, `-yaml`, `-category` are user configurable and can be changed by changing the variable on the right hand side.
* Flag `-yaml` input is a string that tells the type of `kube-bench` YAML and the strings are matched internally to the path of the job YAMLs located in `pkg/kubebench/jobs`. The user just need to enter the type of yaml. Example:
`-yaml=job.yaml`, `-yaml=job-master.yaml`, `-yaml=job-node.yaml`,etc.
* In order to generate policy report in the form of YAML, we can do `kubectl get clusterpolicyreports -o yaml > res.yaml` which will generate it as `res.yaml` in this case.
