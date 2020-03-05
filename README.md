# Skupper enables inter-cluster TCP communication



TCP tunneling with [Skupper](https://skupper.io/)

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Step 1: Set up the demo](#step-1-set-up-the-demo)
* [Step 2: Deploy the Virtual Application Network](#step-2-deploy-the-virtual-application-network)



## Overview

This is a simple demonstration of TCP communication tunneled through a Skupper network from a private to a public cluster and back again. During development of this demonstration, the private cluster was running locally, while the public cluster was on AWS.
<br/>
We will set up a Skupper network between the two clusters, start a TCP echo-server on the public cluster, then communicate to it from the private cluster and receive its replies. At no time is any port opened on the machine running the private cluster.
<br/>


<img src="images/entities.svg" width="800"/>

## Prerequisites

* The `kubectl` command-line tool, version 1.15 or later ([installation guide](https://kubernetes.io/docs/tasks/tools/install-kubectl/))
* The `skupper` command-line tool, the latest version ([installation guide](https://skupper.io/start/index.html#step-1-install-the-skupper-command-line-tool-in-your-environment))
* Two Kubernetes namespaces, from any providers you choose, on any clusters you choose. ( In this example, the namespaces are called 'public' and 'private'. )
* A private cluster running on your local machine.
* A public cluster is running on a public cloud provider.


## Step 1: Set up the demo

1. On your local machine, make a directory for this tutorial and clone the example repo:

   ```bash
   mkdir ${HOME}/tcp-echo
   cd ${HOME}/tcp-echo
   git clone https://github.com/skupperproject/skupper-example-tcp-echo

   ```

2. Prepare the target clusters.

   1. On your local machine, log in to both clusters in a separate terminal session.
   2. In each cluster, set the kubectl config context to use the demo namespace [(see kubectl cheat sheet)](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)





## Step 2: Set Up the Virtual Application Network


1. In the terminal for the public cluster, create the public namespace and deploy the tcp echo server in it :

   ```bash
   kubectl create namespace public
   kubectl config set-context --current --namespace=public
   kubectl apply -f tcp-echo/public-deployment.yaml
   ```

2. Still in the public cluster, start Skupper, expose the tcp-echo deployment, and generate a connection token :

   ```bash
   skupper init
   skupper expose --port 9090 deployment tcp-go-echo
   skupper connection-token ${HOME}/tcp-echo/public_secret.yaml
   ```

3. In the private cluster, create the private namespace : 

   ```bash
   kubectl create namespace private
   kubectl config set-context --current --namespace=private
   ```

4. Start Skupper in the private cluster, and connect to public :

   ```bash
   skupper init
   skupper connect ${HOME}/tcp-echo//public_secret.yaml
   ```

5. See that Skupper is now exposing the public-cluster tcp-echo service on this cluster. (This may take a few seconds. If it's not there immediately, wait a few seconds and try again.) :

   ```bash
   kubectl get svc
   ```
