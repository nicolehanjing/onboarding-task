# Onboarding Task: 
Andrew offers me this great onboarding task to get started with k8s

It takes me some time to finish all the tasks and understand key concepts behind it

I will share the steps I take and record some notes that might be helpful for beginners

## Prerequisite
- A running Kubernetes cluster

- Install [docker](https://docs.docker.com/get-docker/) and get familiar with the basic concepts and commands

- Install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) and get familiar with [basic commands](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands)

```bash
brew install kubectl 
```

## Steps

### Task #1: Building a HTTP Server in Go
Build a http server in Go that returns the current timestamp. It should support 3 URL paths that return the Thcurrent timestamp in different formats.

return the current time using the default format in Go
```
$ curl localhost/
2009-11-10 23:00:00 +0000 UTC m=+0.000000001
```

return the current time using the Unix format
```
$ curl localhost/unix
Tue Nov 10 23:00:00 UTC 2009
```

return the current time using the Unix format
```
$ curl localhost/kitchen
11:00PM
```

<br/><br/>
### Task #2: Containerize HTTP Server
Containerize the HTTP Server built in task #1 with Docker and push the image to a public registry on Docker Hub. 

<br/><br/>
### Task #3: Deploy the HTTP Server to Kubernetes
Deploy the HTTP Server that was containerized in task #2 to a Kubernetes cluster. 

Any other pod in the cluster should be able to send requests to the http server via a pod IP directly or via a Service’s cluster IP or cluster DNS. 

<br/><br/>
### Task #4  Expose the HTTP Server using NodePorts
Expose the HTTP Server via NodePorts so that clients outside the cluster can connect to the http server. 

Any client that can connect to any VMs in the a cluster should be able to send requests to the HTTP server using a node port.

```
$ curl <node-ip>:<node-port>/
2009-11-10 23:00:00 +0000 UTC m=+0.000000001
```
<br/><br/>
### Task #5  Expose the HTTP Server using Ingress
Deploy the Contour ingress controller on your Kubernetes cluster and use the Ingress API
```
$ curl my-service.com/
2009-11-10 23:00:00 +0000 UTC m=+0.000000001
```

### Task #6 Timezone based HTTP Server
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
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
