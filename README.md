# ACT-W Introduction to K8s

This workshop was developed for ACT-W Seattle 2018.

## Pre-Workshop Installations

There is a number of open-source software you will need to install to follow along with the workshop.

### Minikube

Minikube is a tool that allows you to run Kubernetes locally on your computer.

You can find install instructions for minikube [here](https://github.com/kubernetes/minikube/releases). If you use homebrew the easiest thing to do is run this command:

```bash
brew cask install minikube
```

You can verify you have installed minikube by running:

```bash
minikube version
```

### Kubectl

Kubectl is the command-line tool for interfacing with Kubernetes. It allows you to deploy and manage your apps.

Installation instructions can be found [here](https://kubernetes.io/docs/tasks/tools/install-kubectl/). Once again our recommendation is to use homebrew for this:

```bash
brew install kubernetes-cli
```

You can verify you have installed kubectl by running:

```bash
kubectl version
```

## Getting Started

1. Start your minikube by running:

    ```bash
    minikube start
    ```

2. Make sure your context is set to minikube.

    ```bash
    kubectl config get-contexts
    ```

    `minikube` should have a star next to it when you run this command. If it doesn't, you can set minikube to your context by running:

    ```bash
    kubectl config use-context minikube
    ```