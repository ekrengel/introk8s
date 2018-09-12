# Step By Step

A guide to walk through the workshop on your own.

## Required Installs

* Minikube
* Kubectl
* Text Editor
* Terminal

Installation instructions for `kubectl` and `minikube` can be found in the project's [Readme](./README.md).

## Getting Started

1. Start your minikube by running:

    ```bash
    minikube start
    ```

2. Make sure your context is set to minikube.

    ```bash
    kubectl config current-context
    ```

    If it not `minikube` you can set it to `minikube` as your context by running:

    ```bash
    kubectl config use-context minikube
    ```

## The Application

A hello world image has been pushed to [Docker Hub](hub.docker.com/r/ekrengel/introk8s). The application runs on port 8080. It has one handler at `/hello` that returns "Hello ACT-W!". The other handler `/kubernetes` returns the pod name and the node name, which are set through the environment variables `POD_NAME` and `NODE_NAME` respectively. In the case that these environment variables are not set, the application will default to the value `UNKNOWN`.

## First Pod

We will now create our first pod. Create a file `k8s/pod.yaml`.

Follow the solutions in `k8s/solutions/pod1.yaml`. This file specifies that we are creating a resource of kind `pod` with the name `my-first-pod`. We will specify one container with the `ekrengel/introk8s:latest` image, which is the image of the application disucssed above. Additionally, we will need to specify a container port of `8080`, which is the port on which the application runs.

We will now create our pod by applying the file.

```bash
kubectl apply -f k8s/pod.yaml
```

We can see our pod has been created by getting the pods.

```bash
kubectl get pods
```

Now if we want to be able to hit our service we can use the port-forwarding feature of kubernetes, which allows us to map a port on our local machine to the port on a pod.

```bash
kubectl port-forward my-first-pod 8080:8080
```

We can now use a browser or curl our service.

```bash
curl http://localhost:8080/hello
curl http://localhost:8080/kubernetes
```

You will notice the kubernetes endpoint returns the default value of `UNKNOWN`. We can now press Ctrl-C to stop the port forwarding. Lets go ahead and delete the pod so we can fix that.

```bash
kubectl delete -f k8s/pod.yaml
```

## Adding Environment Variables to our Pod

We now want to add the environment variables to our pod.

Add to your `k8s/pod.yaml` file by following the solutions in `k8s/solutions/pod2.yaml`. Under our container we will specify an env section, which takes an array of environment variables to expose on the container. We will set the environment variable `POD_NAME` to the value from the field `metadata.name`. We will also set the environment variable `NODE_NAME` to the value from the field `spec.nodeName`.

We will now re-create our pod by applying the file.

```bash
kubectl apply -f k8s/pod.yaml
```

Lets use port-forwarding again to hit our app.

```bash
kubectl port-forward my-first-pod 8080:8080
```

We can now use a browser or curl our service and see the pod name and node name are set.

```bash
curl http://localhost:8080/kubernetes
```

We can now press Ctrl-C to stop the port forwarding. Then, we will now delete the pod using its name instead of referencing the file.

```bash
kubectl delete my-first-pod
```

If we get our pods again, we will see the pod is gone. More to come on this when we go over deployments.

```bash
kubectl get pods
```

## First Deployment

We will now create our first deployment. Create a file `k8s/deployment.yaml`.

Follow the solutions in `k8s/solutions/deployment.yaml`. We will now create a resource of kind `deployment`. Under spec, we will specify we want to create 3 replicas of our container and that the selector for being considered one of those replicas is matching the label `app: my-first-deployment`. Also under spec, we will create the template section where we specify the the label to use and the containers to create. We will copy over most of what we created in our `k8s/pod.yaml` file under the template container.

We will now create our deployment by applying the file.

```bash
kubectl apply -f k8s/deployment.yaml
```

We can see 3 pods were created as part of our deployment.

```bash
kubectl get pods
```

Now lets copy one of those pod names and try deleting it.

```bash
kubectl delete pod <copied-pod-name>
```

If you get your pods again, you will notice that we still have three pods. The deployment ensure we always have the three pods.

```bash
kubectl get pods
```

Now if we do a port forwarding and curl our `/kubernetes` endpoint we see our pod name has been updated accordingly.

```bash
kubectl port-forward <copied-pod-name> 8080:8080
curl http://localhost:8080/kubernetes
```

We can now press Ctrl-C to stop the port forwarding.

# Service

We will now create our first service. Create a file `k8s/service.yaml`.

Follow the solutions in `k8s/solutions/service.yaml`. We will specify that we are creating a resource of kind `Service`. We can give our service a name under metadata. Under spec, we want to the service to be of type `NodePort` and the selector is the label we specified in the deployment above. We also need to specify the ports related to the service.

Lets apply that.

```bash
kubectl apply -f k8s/service.yaml
```

We can get our services to confirm its there.

```bash
kubectl get services
```

There is a minikube command that allows us to get the ip and node port minikube is exposed to your computer on. 

```bash
minikube service my-first-service --url
```

We can now use that url to curl our `/kubernetes` endpoint.
```bash
curl <url-from-command-above>/kubernetes
```

# Shutting Down

When you're done, you will want to delete these resources and shut down your `minikube`.

```bash
kubectl delete -f k8s/deployment.yaml
kubectl delete -f k8s/service.yaml
minikube stop
```
