# Skupper enables inter-cluster TCP communication



TCP tunneling with [Skupper](https://skupper.io/)

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Step 1: Set up the demo](#step-1-set-up-the-demo)
* [Step 2: Deploy the Virtual Application Network](#step-2-set-up-the-virtual-application-network)
* [Step 3: Access the public service remotely](#step-3-access-the-public-service-remotely)
* [What just happened?](#what-just-happened)
* [Cleaning up](#cleaning-up)
* [Next steps](#next-steps)




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
   2. In each cluster, set the kubectl config context to use the demo namespace [(see Skupper Getting Started Guide)](https://skupper.io/start/index.html)





## Step 2: Set Up the Virtual Application Network


1. In the terminal for the public cluster, create the public namespace and deploy the tcp echo server in it :

   ```bash
   kubectl create namespace public
   kubectl config set-context --current --namespace=public
   kubectl apply -f ${HOME}/tcp-echo/public-deployment.yaml
   ```

2. Still in the public cluster, start Skupper, expose the tcp-echo deployment, and generate a connection token :

   ```bash
   skupper init
   skupper expose --port 9090 deployment tcp-go-echo
   skupper connection-token ${HOME}/tcp-echo/public_secret.yaml
   ```

   Please Note: The connection token contains a secret and should only be shared with trusted sites.

3. In the private cluster, create the private namespace : 

   ```bash
   kubectl create namespace private
   kubectl config set-context --current --namespace=private
   ```

4. Start Skupper in the private cluster, and connect to public :

   ```bash
   skupper init
   skupper connect ${HOME}/tcp-echo/public_secret.yaml
   ```

5. See that Skupper is now exposing the public-cluster tcp-echo service on this private cluster. (This may take a few seconds. If it's not there immediately, wait a few seconds and try again.) :

   ```bash
   kubectl get svc

   # Example output :
   # NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)               AGE
   # skupper-internal    ClusterIP   172.30.202.39    <none>        55671/TCP,45671/TCP   22s
   # skupper-messaging   ClusterIP   172.30.207.178   <none>        5671/TCP              22s
   # tcp-go-echo         ClusterIP   172.30.106.241   <none>        9090/TCP              8s

   ```


## Step 3: Access the public service remotely

One the private cluster, run telnet on the cluster-IP and port that Skupper has exposed for the tcp-echo service.

   ```bash
   telnet 172.30.106.241 9090
   Trying 172.30.106.241...
   Connected to 172.30.106.241.
   Escape character is '^]'.
   Do what thou wilt shall be the whole of the law.
   tcp-go-echo-f55984966-v5px2 : DO WHAT THOU WILT SHALL BE THE WHOLE OF THE LAW.
   ^]
   telnet> quit
   Connection closed.
   ```
 

## What just happened?

The TCP echo server was deployed and running on a publicly accessible cluster. The use of Skupper on that cluster allowed us to generate a connection token, which we then used to securely connect to the public cluster from our privat one. Since the connection was initiated by the Skupper instance on the private cluster, no ports on the private cluster were opened.
<br/>
<br/>
Because we told Skupper to expose the TCP Echo service, when the two Skupper instances connected with each other the private instance learned about that service. The Skupper instance on the private cluster then made a forwarder to that service available on its cluster. 
<br/>
<br/>
We were then able to use telnet to send TCP messages to an IP address and port on the private cluster which Skupper then forwarded to the actual service running in public. Skupper then brought the response back to us, making it feel as if the service were runing locally.
<br/>
<br/>
All Skupper traffic between the two clusters was secured with mutual TLS.



## Cleaning Up

Delete the pod and the virtual application network that were created in the demonstration.

1. In the terminal for the **public** cluster:

   ```bash
   # Get POD ID with 'kubectl get pods'
   $ kubectl delete pod tcp-go-echo-<POD-ID>
   $ skupper delete
   ```

2. In the terminal for the **private** cluster:

   ```bash
   $ skupper delete
   ```


## Next steps

 - [Try the example for Sharing a PostgreSQL database across clusters](https://github.com/skupperproject/skupper-example-postgresql)
 - [Find more examples](https://skupper.io/examples/)

