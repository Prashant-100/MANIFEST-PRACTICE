
# MANIFEST-PRACTICE

some manifest file

## Deployment

To deploy these yaml/yml files and to work on kubernetes time to time run these commands

```bash
kubectl get nodes

kubectl get pods
kubectl get pod
kubectl get pods -o wide

kubectl describe <podname>

kubectl get nodes --show-labels
kubectl label nodes <ip address> instance=<label_name>

kubectl exec <podname> -it -- sh
kubectl exec -it <podname> -- sh
kubectl exec -it <podname> -n <namespace_name> -- sh
kubectl exec <podname> -n <namespace_name> -it -- sh

kubectl create configmap <configmap_name> --from-literal <key>=<value>
kubectl describe configmap/<configmap_name> -n <namespace_name>
kubectl create configmap <configmap_name> -n <namespace_name> --from-literal <key>=<value>
kubectl create secret genric <secret_name> --from-literal <key>=<value>
```
