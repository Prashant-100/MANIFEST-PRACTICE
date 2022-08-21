
# MANIFEST-PRACTICE

some manifest file

## Deployment

To deploy these yaml/yml files and to work on kubernetes time to time run these commands

```bash
kubectl get nodes

kubectl get pods
kubectl get pod
kubectl get pods -o wide

kubectl describe  pods <podname>

kubectl get pods --all-namespace
kubectl get namespace 
-default
-kube-node-lease
-kube-public
-kube-system
kubectl get pods -n <namespace_name>
kubectl create namespace <namespace_name>
kubectl apply -f <podname>.yaml -n <namespace_name>
kubectl get pods -n <namespace_name>
kubectl delete pods <podname> -n <namespace_name> 
kubectl delete namespace <namespace_name>

kubectl get pv

kubectl get nodes --show-labels
kubectl label nodes <ip address> instance=<label_name>

kubectl taint nodes <ip_address_node> <taint_key>=<taint_value>:<taint_action>
kubectl taint nodes <ip_address_node> <taint_key>- ##for untaint

kubectl exec <podname> -it -- sh
kubectl exec -it <podname> -- sh
kubectl exec -it <podname> -n <namespace_name> -- sh
kubectl exec <podname> -n <namespace_name> -it -- sh

kubectl create configmap <configmap_name> --from-literal <key>=<value>
kubectl describe configmap/<configmap_name> -n <namespace_name>
kubectl create configmap <configmap_name> -n <namespace_name> --from-literal <key>=<value>
kubectl create secret genric <secret_name> --from-literal <key>=<value>

Step1:
kubectl create sa <service_account_name>

Step2: Get the secret token generated for serviceaccount 
kubectl describe sa <service_account_name>
## here we get a token for the secret "<service_account_name>-token-*****"
kubectl get secrets

Step3: Get the token value for the secret
kubectl describe secret <secret token from step2>
kubectl describe secret <service_account_name>-token-*****

Step4: Assign the token value from secret to some variable
<variablename>="<secret_value>"

# Single Command to do 4 steps
kubectl describe secret $(kubectl describe sa <service_account_name> | grep Token | awk '{print $NF}') | grep token: | awk '{print $NF}'

# Now assigning the token value using the above command
<variablename>="$(<above_command>)"

Step5: Set the credential for the serviceaccount
kubectl config set-credentials <service_account_name> --token=$<variablename>

Step6: Create a context for the serviceaccount
kubectl get contexts
kubectl config set-context <context_name> --cluster=kubernetes --user=<service_account_name>
kubectl config use-context <context_name>


kubectl create clusterrole my-cluster-role --verb=get --verb=list --verb=watch --resource=pods --resource=secrets
kubectl create clusterrolebinding my-cluster-role-binding --serviceaccount=default:devops --clusterrole=my-cluster-role-binding 
kubectl get clusterrole

```
