# Skupper enables inter-cluster TCP communication



TCP tunneling with [Skupper](https://skupper.io/)

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Step 1: Set up your namespaces](#step-1-set-up-your-namespaces)
* [Step 2: Deploy the backend and frontend services](#step-2-deploy-the-backend-and-frontend-services)
* [Step 3: Connect your namespaces](#step-3-connect-your-namespaces)
* [Step 4: Expose the backend service on the Skupper network](#step-4-expose-the-backend-service-on-the-skupper-network)
* [Step 5: Test the application](#step-5-test-the-application)
* [What just happened?](#what-just-happened)
* [Cleaning up](#cleaning-up)
* [Next steps](#next-steps)


## Overview

This is a simple demonstration of TCP communication tunneled through a Skupper network from a private to a public cluster and back again. During development of this demonstration, the private cluster was running locally, while the public cluster was on AWS.
<br/>
We will set up a Skupper network between the two clusters, start a TCP echo-server on the public cluster, then communicate to it from the private cluster and receive its replies. At no time is any port opened on the machine running the private cluster.
<br/>


<img src="images/entities.svg" width="800"/>


