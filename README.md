# Skupper enables inter-cluster TCP communication.

NOTE : This is a DRAFT in progress. The instructions do not actually work yet.

This is a simple demonstration of TCP communication from a private to a public cluster, and back again. We will set up a Skupper network between the two clusters, start a TCP echo-server on the public cluster, and then telnet to it from the private cluster, and receive its replies.

* [Prerequisites](#prerequisites)
* [Step 1](#step-1-set-up-the-demo)
* [Step 2](#step-2-set-up-the-inter-cluster-skupper-network)
* [Step 3](#step-3-deploy-the-server-and-services)
* [Step 4](#step-4-telnet-from-private-to-public)

## Prerequisites

You will need two clusters: one which we will call 'private' and one which we will call 'public'.

## Step 1 : Set up the demo

1. On the private cluster make a directory for this tutorial, clone the tutorial repo, and download the skupper CLI tool:

   ```bash
   $ mkdir ~/tcp-echo-demo
   $ cd !$
   $ git clone https://github.com/skupperproject/tcp-echo-demo
   $ # Here is the skupper CLI tool :
   $ curl -fL https://github.com/skupperproject/skupper-cli/releases/download/dummy3/linux.tgz -o skupper.tgz
   $ mkdir -p $HOME/bin
   $ tar -xf skupper.tgz --directory $HOME/bin
   $ export PATH=$PATH:$HOME/bin
   ```

   To test your installation, run the 'skupper' command with no arguments. It will print a usage summary.

   ```bash
   $ skupper [command]
   [...]
   ```


## Step 2 : Set up the inter-cluster Skupper network

1. On the 'public' cluster, deploy the 'public' Skupper router and create the secret that will allow 'private' to connect to 'public'.

   NOTE: We assume here that you have two separate shells (windows), one logged in to the 'public' cluster and on logged into the 'private' cluster. But both are on the same machine, and thus both have access to the secrets file which we are about to create.

  ```bash
  $ skupper init --id public
  $ skupper secret ~/tcp-echo-demo/secret.yaml -i public
   ```

2. On the 'private' cluster, deploy the 'private' Skupper router, and connect to 'public' cluster.

  ```bash
  $ skupper init --edge --id private
  $ skupper connect ~/tcp-echo-demo/secret.yaml --name public
   ```


## Step 3 : Deploy the server and services

1. On the 'public' cluster, deploy the tcp-echo server, and the service that exposes it.

  ```bash
  $ oc apply -f ~/tcp-echo-demo/public-deployment-and-service.yaml
   ```

2. Annotate the service. This will cause Skupper to notice the service and prepare to connect it to other clusters.

  ```bash
  TODO -- PUBLIC CLUSTER ANNOTATE COMMAND GOES HERE
  ```

3. On the 'private' cluster, deploy the service only.

  ```bash
  $ oc apply -f i~/tcp-echo-demo/private-service-only.yaml
  ```

4. Annotate the private service. This will cause Skupper to notice it, and connect it to its counterpart on the other cluster.

  ```bash
  TODO -- PRIVATE CLUSTER ANNOTATE COMMAND GOES HERE
  ```

## Step 4 : telnet from private to public

1. On the private cluster, find the IP address and port of the service and telnet to it. Skupper routes your message from the private to the public cluster where it is capitalized by the tcp-echo service and returned to you. ( Also, the server prepends its pod ID to the capitalized message. )

  ```bash
  $ kubectl get svc
  TODO -- SHOW EXAMPLE OUTPUT HERE


  $ telnet 172.30.67.11 27031/TCP 
  Trying 127.0.0.1...
  Connected to 127.0.0.1.
  Escape character is '^]'.
  Does this really work?
  53c9235e175e : DOES THIS REALLY WORK?
  Yes! This really works!
  53c9235e175e : YES! THIS REALLY WORKS!
  ```


