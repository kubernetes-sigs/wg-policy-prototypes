# Falco adapter
The falco adapter runs as an output for falcosidekick when enabled, the adapter receives falco events (Falco ships with a default set of rules that check the kernel for unusual behavior) and produces N+1 reports (1 cluster-wide policy report and N namespace specific reports) based on the [Policy Report Custom Resource Definition](https://github.com/kubernetes-sigs/wg-policy-prototypes/tree/master/policy-report)

## Running

**Prerequisites**: 
* To run the Kubernetes cluster locally, tools like [kind](https://kind.sigs.k8s.io/) or [minikube](https://minikube.sigs.k8s.io/docs/start/) can be used. Here are the steps to run the falco adapater with a `kind` cluster.
```sh
kind create cluster --config=kind-config.yaml
```
Check out an example of [kind-config.yaml](https://gist.github.com/anushkamittal20/0e21b237b6ff98773675edf4e58be96a/)
```sh
helm repo add falcosecurity https://falcosecurity.github.io/charts

helm repo update 
```
```sh
kubectl create -f https://github.com/kubernetes-sigs/wg-policy-prototypes/raw/master/policy-report/crd/v1alpha2/wgpolicyk8s.io_clusterpolicyreports.yaml

kubectl create -f https://github.com/kubernetes-sigs/wg-policy-prototypes/raw/master/policy-report/crd/v1alpha2/wgpolicyk8s.io_policyreports.yaml
```
```sh
helm install falco falcosecurity/falco --set falcosidekick.enabled=true --set falcosidekick.policyreport.enabled=true falcosidekick.policyreport.kubeconfig=~/.kube/config falcosidekick.policyreport.failthreshold=3 falcosidekick.policyreport.maxevents=10
```
Above can be configured according to specifications
```sh
kubectl port-forward svc/falco-falcosidekick 2801
```
5 can be configured according to specifications

## Understanding the config options
 * Once falcosidekick is enabled in policyreport output we have the following configurations available
 ~~~
1. Enabled = to enable policyreport output to create/update policyreports
2. Kubeconfig = address to the file (default- ~/.kube/config)
3. FailThreshold = events with priority above this threshold are mapped to Fail in PolicyReportSummary and "high" severity; rest are mapped to Warn in PolicyReportSummary (and severity "low" if event priority is below the threshold and "medium if it is equal to threshold" )
4. MaxEvents = this specifies the maximum number of events in any of the N+1 reports; once the events go above this number the report start self pruning
5. PruneByPriority = while pruning by default the events that came initially will be deleted (FIFO); by enabling this config the events that came initially of low priority are deleted before initial events of higher priority
~~~

## To get Report summary
```sh
kubectl get clusterpolicyreports

kubectl get policyreports --all-namespaces
```

## To get Reports
```sh
kubectl get clusterpolicyreports -o yaml 

kubectl get policyreports --all-namespaces -o yaml

```
To get reports in a separate yaml file you can use ` kubectl get clusterpolicyreports  -o yaml > res.yaml` or `kubectl get policyreports --all-namespaces -o yaml > res.yaml`



