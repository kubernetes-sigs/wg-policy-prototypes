# Trivy Adapter
The image vulnerability adapter runs an image vulnerability scanner tool called [trivy](https://github.com/aquasecurity/trivy) to scan a given kubernetes pod or workload to produces a namespace policy report based on the [Policy Report Custom Resource Definition](https://github.com/kubernetes-sigs/wg-policy-prototypes/tree/master/policy-report)

## Installing
Build the trivy-adapter binary: `make trivy-adapter` and install `make install`

Type `trivy-adapter` command in your terminal to confirm if the image vulnerability adapter is installed already


## Running

**Prerequisites**: 
* Kubernetes Cluster: To run the Kubernetes cluster locally, tools like [kind](https://kind.sigs.k8s.io/) or [minikube](https://minikube.sigs.k8s.io/docs/start/) can be used. Here are the steps to run the image vulnerability adapter with getting a `kind` cluster up and running.

### Steps

#### Common steps
```sh
# 1. Create a Kubernetes cluster
kind create cluster --name mycluster --image <kindest/node:stpecify version tag>

# 2. Create a policy report CustomResourceDefinition
kubectl create -f kubernetes/crd/wgpolicyk8s.io_policyreports.yaml

# 3. Create resources in the cluster that trivy-adapter will use.
trivy-adapter init

# 4. Create a pod of image openzipkin/zipkin:latest called zipkin
kubectl create deployment zipkin --image openzipkin/zipkin:latest
```
**Note**:
* You can also deploy your pod that has more than one containers.

#### Steps to run image vulnerability adapter on your pod or workload
```sh
# 5. Check your pods and namespaces
kubectl get pod --all-namespaces

# 6. Scan default pods
trivy-adapter scan policyreports zipkin-uewcde32cs9-ui

# 7. Check policyreports created through the custom resource
kubectl get policyreports
```
**Note**:
* You should add Flags `-yaml` or `json` to view the full policy report of each pod.
```sh
kubectl get policyreports pod-zipkin-uewcde32cs9-ui -o yaml
```
#### Roadmap:
* CronJob for periodic scan on the pods or workloads to create and update the namespace policy report.
* Extend scan to a cluster wide scan to scan all workloads and pods to create or update the cluster policy report periodically using CronJob.
