# dinghy-worker

This repo is to be used in conjunction with https://github.com/izaakdale/dinghy-agent.

*It is imperitive this you have one agent running in your cluster before continuing.*

Dinghy is my attempt to create a distributed key value store using Serf (https://github.com/hashicorp/serf) and Raft (https://github.com/hashicorp/raft) taking inspiration from Kubernetes (k8s) etcd.
The aim when starting this project was to gain a deeper understanding of distributed systems in general but also as research into the inner workings of k8s.

## Get started

Point your terminal's docker-cli to the Docker Engine inside minikube

```eval $(minikube docker-env)```

Create the worker container

```make docker```

```make up```

This will deploy 3 replicas of the worker app. I then use Postman and server reflection to store and fetch k/v pairs from the distributed DB.
