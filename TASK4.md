# Task 4
### Expose the HTTP Server using NodePorts
Expose the HTTP Server via NodePorts so that clients outside the cluster can connect to the HTTP server. 

Any client that can connect to any VMs in a cluster should be able to send requests to the HTTP server using a node port.

### Create a Kubernetes Service
The port-forward command is good for testing the pods directly. But in production, you would want to expose the pod using services.

Instead of relying on the Pods IP addresses which change, Kubernetes provides services as stable endpoint for pods.

The pods that the service exposes are based on a set of labels. 

If Pods have the correct labels, they are automatically picked up and exposed by our services.

<br>
First, add the following snippets into my-go-k8s-deployment.yml file

```
apiVersion: v1
kind: Service                   
metadata:
    name: go-current-time-service   
spec:
    type: NodePort                 # A port is opened on each node in your cluster via Kube proxy.
    ports:                         # Take incoming HTTP requests on port 9090 and forward them to the targetPort of 8080
    - name: http
      port: 8081
      targetPort: 8081
      nodePort: 30007
    selector:
        app: go-current-time         # Map any pod with label `app=go-current-time` to this service
```

create the service:

```
$ kubectl apply -f my-go-k8s-deployment.yml
deployment.apps/go-current-time unchanged
service/go-current-time-service created

$ kubectl get services
NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)             AGE
go-current-time-service               NodePort    100.69.120.197   <none>        9090:31888/TCP      75s

$ kubectl describe service go-current-time-service
Name:                     go-current-time-service
Namespace:                default
Labels:                   <none>
Annotations:              Selector:  app=go-current-time
Type:                     NodePort
IP:                       100.69.120.197
Port:                     http  9090/TCP
TargetPort:               8080/TCP
NodePort:                 http  31888/TCP
Endpoints:                100.107.59.219:8080,100.107.59.220:8080,100.107.59.221:8080
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

curl the URL for the service

1. <NodeIP>:spec.ports[*].nodePort 
2. .spec.clusterIP:spec.ports[*].port

```
$ kubectl get nodes -o wide
NAME                         STATUS   ROLES    AGE    VERSION            INTERNAL-IP      EXTERNAL-IP      OS-IMAGE                 KERNEL-VERSION   CONTAINER-RUNTIME
test-control-plane-nblbw     Ready    master   5d5h   v1.18.3+vmware.1   10.160.216.105   10.160.216.105   VMware Photon OS/Linux   4.19.126-1.ph3   containerd://1.3.4
test-md-0-5d4c9f854c-6j2ks   Ready    <none>   38h    v1.18.3+vmware.1   10.160.207.190   10.160.207.190   VMware Photon OS/Linux   4.19.126-1.ph3   containerd://1.3.4
```

I ran into a bug - cannot view EXTERNAL-IP of my nodes, which was fixed by running:
```
$ kubectl -n kube-system delete po -l k8s-app=vsphere-cloud-controller-manager
pod "vsphere-cloud-controller-manager-7969w" deleted
$ kubectl get nodes -o wide
NAME                         STATUS   ROLES    AGE     VERSION            INTERNAL-IP      EXTERNAL-IP      OS-IMAGE                 KERNEL-VERSION   CONTAINER-RUNTIME
test-control-plane-nblbw     Ready    master   6d3h    v1.18.3+vmware.1   10.160.216.105   10.160.216.105   VMware Photon OS/Linux   4.19.126-1.ph3   containerd://1.3.4
test-md-0-5d4c9f854c-6j2ks   Ready    <none>   2d12h   v1.18.3+vmware.1   10.160.207.190   10.160.207.190   VMware Photon OS/Linux   4.19.126-1.ph3   containerd://1.3.4
```

Now we can see the EXTERNAL-IP of workload node "test-md-0-5d4c9f854c-6j2ks"
we have 2 options:
- option1: <nodeIP>:<nodePort>
- option2: <ClusterIP>:<port>

Try option 1:
```
$ curl http://10.160.207.190:30007
curl: (7) Failed to connect to 10.160.207.190 port 30007: Connection refused
```

There is a connection error, then run:
```
$ kubectl get endpoints
NAME                                  ENDPOINTS                                                     AGE
go-current-time-service               100.107.59.226:8080,100.107.59.227:8080,100.107.59.228:8080   3h18m
```
In the go app, I wrote "http.ListenAndServe(":8081", nil)"

but then I did port-forwarding 8081:8080 so I can curl localhost:8080 and get correct response

now change the port in the YAML file "- containerPort: 8081"

"targetPort: 8081" and "port: CAN BE ANY"

```
# <nodeIP>:<nodePort>
$ curl 10.160.207.190:30007
2020-08-11 21:37:01.312486378 +0000 UTC m=+21.799144038
```

### Yay!!! It's working now!
