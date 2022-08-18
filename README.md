
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
```
