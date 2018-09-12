# ACT-W Introduction to K8s

This workshop was developed for ACT-W Seattle 2018.

## Installations

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

## Step by Step Guide

If you would like to walk through the workshop on your own, follow the [step-by-step guide](./step_by_step.md).
