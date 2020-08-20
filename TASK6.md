## Task #6 Timezone based HTTP Server
Update the HTTP server to return the current timestamp in a particular timezone. 

The timezone can be set based on an environment variable passed into the container. The HTTP server should be updated to read the value in the environment variable and convert the timestamp accordingly. 

Run 3 different deployments of the HTTP server, with each deployment returning the current timestamp in EST, PST, and UTC.

The ingress controller should request to the 3 deployments based on the timezone specified in the hostname. For example:
```
$ curl est.current-time.io/
<current time in eastern time>   
$ curl pst.current-time.io/
<current time in pacific time>
$ curl utc.current-time.io/
<current time in UTC>
```

<br><br>

First set environment variables, in each deployment yaml file, add:
```
env:
        - name: timezone
          value: PST / UTC / EST
```

now we have 3 deployments with different env values

see: https://www.magalix.com/blog/kubernetes-patterns-environment-variables-configuration-pattern

All programming languages have a way to retrieve environment variables from the OS.

edit main.go to refelect the change

All the environment variables can be found in e.g. 
```
os.Getenv(“{Key}”)
```

We build the new image for our go app and tag it with name "my-go-app"

```
$ docker build -t my-go-app .
$ docker images
REPOSITORY                                     TAG                                 IMAGE ID            CREATED              SIZE
my-go-app                                      latest                              021eb0817105        About a minute ago   367MB

$ docker tag 021eb0817105 <YOUR DOCKER HUB ID>/test-go:secondtry
$ docker push <YOUR DOCKER HUB ID>/test-go
```


<br>
Now we have a new image, remember to edit the image name of the container in deployment YAML file and apply the change

```
$ kubectl get deployments
NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
est-current-time                      0/3     3            0           7s
pst-current-time                      3/3     3            3           7s
utc-current-time                      0/3     3            0           7s
```

There are 3 YAML files called "pst-deployment.yaml", etc
```
$ kubectl apply -f pst-deployment.yaml
deployment.apps/pst-current-time-deployment created
service/pst-current-time-service created

$ kubectl apply -f est-deployment.yaml
deployment.apps/est-current-time-deployment created
service/est-current-time-service created

$ kubectl apply -f utc-deployment.yaml
deployment.apps/utc-current-time-deployment created
service/utc-current-time-service created

$ kubectl get deployments
NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
est-current-time-deployment           3/3     3            3           72s
pst-current-time-deployment           3/3     3            3           119s
utc-current-time-deployment           3/3     3            3           25s

$ kubectl get pods
NAME                                                   READY   STATUS    RESTARTS   AGE
est-current-time-deployment-788df977bc-6xdqb           1/1     Running   0          111s
est-current-time-deployment-788df977bc-stnf9           1/1     Running   0          111s
est-current-time-deployment-788df977bc-tfhd2           1/1     Running   0          111s
pst-current-time-deployment-6f6bc79cb4-6c2mz           1/1     Running   0          2m38s
pst-current-time-deployment-6f6bc79cb4-82sg8           1/1     Running   0          2m38s
pst-current-time-deployment-6f6bc79cb4-jcp2h           1/1     Running   0          2m38s
utc-current-time-deployment-854f4468bf-7ffxs           1/1     Running   0          64s
utc-current-time-deployment-854f4468bf-jzjnv           1/1     Running   0          64s
utc-current-time-deployment-854f4468bf-m79vh           1/1     Running   0          64s
```

<br>
Edit ingress-test.yaml and apply it
```
$ kubectl apply -f ingress-test.yaml
ingress.networking.k8s.io/test-ingress configured
```

Add 3 hosts and add them in /etc/hosts, then see if everything works:
```
$ curl est.current-time.io/kitchen
8:03PM
$ curl pst.current-time.io/kitchen
5:03PM
$ curl utc.current-time.io/kitchen
2:06AM
```
