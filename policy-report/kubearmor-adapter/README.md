# kubearmor-adapter

This KubeArmor Policy Report adapter converts output received from [KubeArmor](https://github.com/kubearmor/KubeArmor) and produces a policy report based on the [Policy Report Custom Resource Definition](https://github.com/kubernetes-sigs/wg-policy-prototypes/tree/master/policy-report).

## Running

**Pre-requisites**:
* Follow the [Development Guide](https://github.com/kubearmor/KubeArmor/blob/master/contribution/development_guide.md) to install and run KubeArmor on your machine. Using a vagrant environment is highly recommended. Once you have KubeArmor up and running on your machine, you're ready to create and update policy reports.

### Steps:

```sh
# 1. Clone the kubearmor-adapter GitHub repository
git clone https://github.com/haardikdharma10/kubearmor-adapter
```

```sh
# 2. cd into kubearmor-adapter directory
cd kubearmor-adapter
```

```sh
# 3. Create a CustomResourceDefinition
kubectl create -f crd/v1alpha2/wgpolicyk8s.io_policyreports.yaml
```

```sh
# 4. Run main.go program
go run main.go
```
**Note:** Make sure you have the KubeArmor service running in the background before you run this command. If not, you can cd into `KubeArmor/KubeArmor` and run `make clean && make run`. 

```sh
# 5. Open a new terminal window and deploy the multiubuntu microservice.
cd KubeArmor/examples/multiubuntu
kubectl apply -f .
```

```sh
#6. Deploy a security policy for testing
cd KubeArmor/examples/multiubuntu/security-policies
kubectl -n multiubuntu apply -f ksp-group-1-proc-path-block.yaml
```

```sh
#7. Trigger a policy violation by running the following command
kubectl -n [namespace-name] exec -it [pod-name] -- bash -c "/bin/sleep 1"
```
**Note:** In this example, namespace-name is `multiubuntu` and you can get the pod name by running `kubectl get pods -n multiubuntu`. An example pod-name is `ubuntu-1-deployment-5d6b975744-rrkhh`.

Once this command is executed, you'll get the output as below in the terminal window where `main.go` is running:

![image](/assets/create-output.png)

If you can see the output as above, this means that your first policyreport is created. You can now stop running the main program.

### Viewing reports

You can view the summary of the created policyreport by running the following command:
```sh
kubectl get policyreports -n multiubuntu
```

To view the policyreport in yaml format, you can use:
```sh
kubectl get policyreports -n multiubuntu -o yaml
```
To view the report in a separate yaml file you can use:
```sh
kubectl get policyreports -o yaml > res.yaml
```
A new file `res.yaml` will be created in the kubearmor-adapter directory. You can view it by running `cat res.yaml`.

To delete the policyreport, you can use:
```sh
kubectl delete policyreports -n [namespace-name] [policy-report-name]
```

In our example, namespace-name is `multiubuntu` and policy-report-name is `kubearmor-policy-report`.
