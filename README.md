# Skupper enables inter-cluster TCP communication



TCP tunneling with [Skupper](https://skupper.io/)

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Step 1: Set up the demo](#step-1-set-up-the-demo)
* [Step 2: xxx](#xxx)


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
   mkdir tcp-echo
   cd tcp-echo
   git clone https://github.com/skupperproject/skupper-example-tcp-echo

   ```

2. Prepare the target clusters.

   1. On your local machine, log in to both clusters in a separate terminal session.
   2. In the local cluster create the namespace 'private'. In the remote cluster, create the namespace 'public'.
   3. In each cluster, set the kubectl config context to use the demo namespace [(see kubectl cheat sheet)](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)





