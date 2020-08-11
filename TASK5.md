# Task 5
### Expose the HTTP Server using Ingress
Deploy the Contour ingress controller on your Kubernetes cluster and use the Ingress API

just in case there is some mix up in the configuration:
```
$ kubectl delete namespace projectcontour
```


Deploy Countour by applying the YAML file
```
$ kubectl apply -f https://projectcontour.io/quickstart/contour.yaml
```
This command creates:
- A new namespace projectcontour
- Two instances of Contour in the namespace
- A Kubernetes Daemonset running Envoy on each node in the cluster listening on host ports 80/443
- A Service of type: LoadBalancer that points to the Contour’s Envoy instances
- Depending on your deployment environment, new cloud resources – for example, a cloud load balancer

<br>
<br>
Now retrieve the external address of Countour's Envoy load balancer:
```
$ kubectl get -n projectcontour service envoy -o wide
NAME    TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE     SELECTOR
envoy   LoadBalancer   100.70.20.188   <pending>     80:30973/TCP,443:31369/TCP   3d20h   app=envoy
```

verify that the Contour and Envoy pods are running
```
$ kubectl get pods -n projectcontour
NAME                           READY   STATUS      RESTARTS   AGE
contour-5cf47f5cf9-5jdck       1/1     Running     0          25m
contour-5cf47f5cf9-mbhsq       1/1     Running     0          25m
contour-certgen-v1.5.0-5kl4m   0/1     Completed   0          25m
envoy-glvgb                    2/2     Running     0          25m
```

OPTIONAL: apply our written network policy
```
$ kubectl apply -f web-deny-all.yaml
networkpolicy "web-deny-all" created
```
To delete:
```
$ kubectl delete networkpolicy web-deny-all
```


Ping the load balancer using the external IP address
```
$ ping 10.160.207.190
$ curl 10.160.207.190:80
```
we got an empty message


Create an Ingress by creating the following YAML file named ingress-test.yaml
```
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: test-ingress
spec:
  rules:
  - host: go.current.time
    http:
      paths:
      - path: /
        backend:
          serviceName: go-current-time-service
          servicePort: 8081
```

apply this Ingress file:
```
$ kubectl apply -f ingress-test.yaml
```

verify 2 services are created
```
$ kubectl get ingress
NAME                                   CLASS    HOSTS                        ADDRESS   PORTS     AGE
harbor-service-harbor-ingress-notary   <none>   notary.harbor.system.tanzu             80, 443   4d
test-ingress                           <none>   go.current.time                        80        8s

$ vim /etc/hosts
```
add 
```
10.160.207.190 go.current.time
``` 
at the bottom of this file


Now verify you can access the services using the ingress routes
```
$ curl go.current.time
2020-08-11 22:53:07.263547211 +0000 UTC m=+4587.750225439
```
It's working! 

Wecan also visit http://go.current.time/kitchen
- see "10:53PM"
