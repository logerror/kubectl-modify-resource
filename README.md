# kubectl-modify-resource

A kubectl plugin to update deployment resource (cpu and mem)

## Installation/Upgrade

```
curl -sL https://github.com/logerror/kubectl-modify-resource/releases/download/1.0.0/kubectl-modify_resource-1.0.0-linux-amd64.tar.gz -o kubectl-modify_resource-1.0.0-linux-amd64.tar.gz
tar -xzvf kubectl-modify_resource-1.0.0-linux-amd64.tar.gz
mv kubectl-modify_resource /usr/local/bin
chmod +x /usr/local/bin/kubectl-modify_resource
```

For other OS/Arch:

```
replace download url with this file :

kubectl-modify_resource-1.0.0-darwin-amd64.tar.gz
kubectl-modify_resource-1.0.0-darwin-arm64.tar.gz
kubectl-modify_resource-1.0.0-linux-amd64.tar.gz
kubectl-modify_resource-1.0.0-linux-arm64.tar.gz
kubectl-modify_resource-1.0.0-windows-amd64.zip

```

## Usage

```
kubectl modify_resource --kubeconfig=xx/you-kubeconfig-path --deployment=deploy-test --namespace=k8s-namespace --cpu-request=200m
 
kubectl modify_resource --help                                                                                                                                                    
  -cpu-limit string
        CPU limit
  -cpu-request string
        CPU request
  -deployment string
        Deployment Name
  -kubeconfig string
        kubeconfig path
  -memory-limit string
        Mem limit
  -memory-request string
        Mem request
  -namespace string
        namespace

```