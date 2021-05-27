# policy-report-prototype
Building a prototype of Policy Report Generator. It aims to run a CIS benchmark check with a tool called [kube-bench](https://github.com/aquasecurity/kube-bench) and produce a policy report based on the Custom Resource Definition accordingly.

## Running

**Prerequisites**: 
* Since the policy-report-prototype uses `apps/v1` deployments, the Kubernetes cluster version should be greater than 1.9.
* To run the Kubernetes cluster locally, tools like [kind](https://kind.sigs.k8s.io/) or [minikube](https://minikube.sigs.k8s.io/docs/start/) can be used. In our case, we will be going with [kind](https://kind.sigs.k8s.io/). You can follow the links if kind or minikube aren't installed on your local machine.

### Steps

#### Common steps
```sh
# 1. create a Kubernetes cluster
kind create cluster
    OR
minikube start

# 2. create a CustomResourceDefinition
kubectl create -f kubernetes/crd/v1alpha2/wgpolicyk8s.io_policyreports.yaml
```
#### Steps to run kube-bench adapter in-cluster as a Cron-Job
```sh
# 3. Create Role, Role-Binding and Services
kubectl create -f kubernetes/role.yaml -f kubernetes/rb.yaml -f kubernetes/service.yaml

# 4. Create cron-job
kubectl create -f kubernetes/cron-job.yaml 

# 5. Watch the jobs
kubectl get jobs --watch

# 6. check policyreports created through the custom resource
kubectl get policyreports
```

#### Steps to run kube-bench adapter outside-cluster 
```sh
# 3. Build
make build

# 4. Create policy report using
./policyreport -name="kube-bench" -kube-bench-targets="master,nodes" -yaml="job.yaml" -namespace="default" -category="CIS Benchmarks"

# 5. check policyreports created through the custom resource
kubectl get policyreports
```
**Notes**:
* Flags `-name`,`-namespace`, `-yaml`, `-category` are user configurable and can be changed by changing the variable on the right hand side.
* Flag `-yaml` input is a string that tells the type of `kube-bench` YAML and the strings are matched internally to the path of the job YAMLs located in `pkg/kubebench/jobs`. The user just need to enter the type of yaml. Example:
`-yaml=job.yaml`, `-yaml=job-master.yaml`, `-yaml=job-node.yaml`,etc.
* In order to generate policy report in the form of YAML, step 5 can be written as `kubectl get policyreports -o yaml > res.yaml` which will generate it as `res.yaml` in this case.
