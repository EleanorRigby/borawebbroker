# BoreWeb Broker

This is an implementation of a Service Broker that uses Helm to provision
instances of mariadb, mysql, drupal and wordpress. This is a
**proof-of-concept** for the borathon idea, and should not
be used in production.


## Prerequisites

1. Kubernetes cluster ([minikube](https://github.com/kubernetes/minikube))
2. [Helm 2.x](https://github.com/kubernetes/helm)
3. [Service Catalog API](https://github.com/kubernetes-incubator/service-catalog) - follow the [walkthrough](https://github.com/kubernetes-incubator/service-catalog/blob/master/docs/walkthrough.md)

## Installing the Broker

The boraweb Service Broker can be installed using the Helm chart in this
repository.

```
$ git clone https://github.com/EleanorRigby/borawebbroker.git
$ cd borawebbroker
$ helm install --name borawebbroker --namespace borawebbroker charts/borawebbroker
```

