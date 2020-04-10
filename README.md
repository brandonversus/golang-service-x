# Service-X

This application is simple proxy that can easily be configured to chain calls to other services. I've found it very helpful for standing up applications quickly in GKE and testing routing configurations.

## Usage
You can download the latest image from the **brandonversus** Docker Hub page here: https://hub.docker.com/repository/docker/brandonversus/service-x

## Example
The example in this repository shows three services that all use the `service-x` image and are configured to call each other sequentially. 
We have `service-a` calling `service-b` calling `service-c`.

We have a virtual service for `service-b` configured so that half of the traffic will randomly be re-routed to `service-c`. After deploying and port-forwarding to `service-a`, we get the following responses when we call localhost on port 8080:

* `{"hops":["service-c","service-b","service-a"]}`
* `{"hops":["service-c","service-a"]}`
* `{"hops":["service-c","service-b","service-a"]}`

You can see in the second call that the call from `service-a` to `service-b` got rerouted to `service-c`.
