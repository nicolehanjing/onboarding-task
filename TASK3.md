# Task 3
### Deploy the HTTP Server to Kubernetes
Deploy the HTTP Server that was containerized in task #2 to a Kubernetes cluster. 

Any other pod in the cluster should be able to send requests to the http server via a pod IP directly or via a Service’s cluster IP or cluster DNS. 

<br><br>
To deploy our app on Kubernetes, we need to first containerize it.

Deployments are a declarative way to instruct Kubernetes how to create and update instances of your application. 

A deployment consists of a set of identical, indistinguishable Pods.

first we created a manifest file with a bunch of configs containing the desired state of our application

```
$ kubectl config view
```

see there are 2 clusters running

```
$ kubectl cluster-info
Kubernetes master is running at https://10.160.210.215:6443
KubeDNS is running at https://10.160.210.215:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

### Write the YML file to create deployment
```
apiVersion: apps/v1
kind: Deployment                 
metadata:
  name: go-current-time         
spec:
  replicas: 3                    # Number of pods to run at any given time
  selector:
    matchLabels:
      app: go-current-time        # This deployment applies to any Pods matching the specified label
  template:                      # This deployment will create a set of pods using the configurations in this template
    metadata:
      labels:                    # The labels that will be applied to all of the pods in this deployment
        app: go-current-time 
    spec:                        # Spec for the container which will run in the Pod
      containers:
      - name: go-current-time
        image: nicolehan1996/test-go:firsttry
        imagePullPolicy: IfNotPresent
        ports:
          - containerPort: 8083  # Should match the port number that the Go application listens on
```

Carefully read the file and understand it

```
$ kubectl apply -f my-go-k8s-deployment.yml
deployment.apps/pst-current-time-deployment created

$ kubectl get deployments
NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
go-current-time                       3/3     3            3           22s

$ kubectl get pods
NAME                                                   READY   STATUS    RESTARTS   AGE
go-current-time-bcf9968d9-bwvtn                        1/1     Running   0          45s
go-current-time-bcf9968d9-h72xj                        1/1     Running   0          45s
go-current-time-bcf9968d9-m6clh                        1/1     Running   0          45s
```

Pods are allocated a **private IP address** by default and cannot be reached outside of the cluster

You can use the kubectl port-forward command to map a local port to a port inside the pod like this:
```
$ kubectl port-forward go-current-time-bcf9968d9-bwvtn 8080:8080
Forwarding from 127.0.0.1:8080 -> 8080
Forwarding from [::1]:8080 -> 8080
```

open another terminal
```
$ curl localhost:8080/
2020-08-10 22:25:48.3527677 +0000 UTC m=+7869.514009301
$ curl localhost:8080/kitchen
10:25PM
```

to view more logs:
```
$ kubectl logs -f go-current-time-bcf9968d9-bwvtn
```

from a pod, send a request to port 8080, if you meet errors, run:
```
$ docker inspect my-go-image
```
[view the output and find what's in "Cmd"\(https://stackoverflow.com/questions/29535015/error-cannot-start-container-stat-bin-sh-no-such-file-or-directory)

```
$ kubectl exec -it go-current-time-5f5c457d7f-6hfcs sh
```
