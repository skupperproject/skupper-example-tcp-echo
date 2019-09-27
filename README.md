# Skupper enables inter-cluster TCP communication.

This is a simple demonstration of TCP communication tunneled through a Skupper network from a private to a public namespace and back again. We will set up a Skupper network between the two namespaces, start a TCP echo-server on the public namespace, then communicate to it from the private namespace, and receive its replies. We will assume that Kubernetes is running on your local machine, and we will create and access both namespaces from within a single shell.

* [Prerequisites](#prereq)
* [Step 1: Set up the demo.](#step_1)
* [Step 2: Start your cluster and define two namespaces.](#step_2)
* [Step 3: Start Skupper in the public namespace.](#step_3)
* [Step 4: Make a connection token, and start the service.](#step_4)
* [Step 5: Start Skupper in the private namespace..](#step_5)
* [Step 6: Make the connection.](#step_6)
* [Step 7: Communicate across namespaces.](#step_7)
* [Step 8: Cleanup.](#step_8)

## Prerequisites  <a name="prereq"></a>

You will need the skupper command line tool installed, and on your executable path.



## Step 1: Set up the demo. <a name="step_1"></a>

On your machine make a directory for this tutorial, clone the tutorial repo, and download the skupper CLI tool:

   ```bash
    mkdir ${HOME}/tcp-echo-demo
    cd !$
    git clone https://github.com/skupperproject/skupper-example-tcp-echo
   ```



## Step 2: Start your cluster and define two namespaces.  <a name="step_2"></a>

   ```bash
   alias kc='kubectl'
   oc cluster up
   oc new-project public
   oc new-project private
   ```

## Step 3: Start Skupper in the public namespace.  <a name="step_3"></a>

   ```bash
   kc config set-context --current --namespace=public
   skupper status
   skupper init --cluster-local --id public
   skupper status
   ```

## Step 4: Make a connection token, and start the service. <a name="step_4"></a>

   ```bash
   skupper connection-token ${HOME}/secret.yaml
   oc apply -f ${HOME}/tcp-echo-demo/skupper-example-tcp-echo/public-deployment-and-service.yaml
   ```

## Step 5: Start Skupper in the private namespace.  <a name="step_5"></a>

   ```bash
   kc config set-context --current --namespace=private
   skupper status
   skupper init --cluster-local --id private
   skupper status
   ```


## Step 6: Make the connection.  <a name="step_6"></a>

After issuing the connect command, a new service will show up in this namespace called tcp-go-echo. (It may take as long as two minutes for the service to appear.)

   ```bash
   $ skupper connect ${HOME}/secret.yaml
   $ kc get svc
   NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)               AGE
   skupper-internal    ClusterIP   172.30.46.68     <none>        55671/TCP,45671/TCP   2m
   skupper-messaging   ClusterIP   172.30.180.253   <none>        5671/TCP              2m
   tcp-go-echo         ClusterIP   172.30.17.63     <none>        9090/TCP              38s
   ```


## Step 7: Communicate across namespaces.  <a name="step_7"></a>

Using the IP address and port number from the 'kc get svc' result, send a message to the local service. Skupper will route the message to the service that is running on the other namespace, and will route the reply back here.

   ```bash
   ncat Mr. Watson, come here. I want to see you.
   tcp-go-echo-67c875768f-kt6dc : MR. WATSON, COME HERE. I WANT TO SEE YOU.
   ```

The tcp-go-echo program returns a capitalized version of the message, prepended by its name and pod ID.


## What Just Happened ?

Your <i>ncat</i> TCP message was received by the Skupper-created tcp-go-echo proxy in namespace 'private', wrapped in an AMQP message, and sent over the Skupper network to the Skupper-created proxy in the 'public' namespace. That proxy sent the TCP packets to the tcp-go-echo server (which knows nothing about AMQP), received its response, and reversed the process. After another trip over the Skupper network, the TCP response packets arrived back at our ncat process.
<br/>
We demonstrated this using two namespaces in a single local cluster for ease of demonstration, but the establishment and use of Skupper connectivity works just as easily between any two (or more) clusters, public or private, anywhere.
<br/>


## Step 8: Cleanup. <a name="step_8"></a>

Let's tidy up so no one trips over any of this stuff later. In the private namespace, delete the Skupper artifacts. In public, delete both Kubernetes and Skupper atrifacts.

   ```bash
   kc config set-context --current --namespace=private
   skupper delete
   kc config set-context --current --namespace=public
   kc delete -f ${HOME}/tcp-echo-demo/skupper-example-tcp-echo/public-deployment-and-service.yaml
   skupper delete
   ```
<br/>
<br/>
