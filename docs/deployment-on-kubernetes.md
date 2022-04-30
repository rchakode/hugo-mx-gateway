# Deploying hugo-mx-gateway on Kubernetes

  - [Overview](#overview)
  - [Setup Procedure](#setup-procedure)

## Overview
From the source tree, the folder `./helm/` contains Helm manifest to ease the deployment of hugo-mx-gateway on Kubernetes clusters. 

> **Important:** The chart is validated with Helm 3 and the pod is run in an unprivileged mode within a **Security Context**.

## Setup Procedure
Proceed with the deployment as follows:

* First edit the [values.yaml](./helm/values.yaml) file to set [configuration values](#configuration-variables) appropriately.
* Choose a deployment namespace. In the sample commands provided hereafter, it's assumed that the target namespace is `hugo-mx-gateway`. If you opt for another namespace, do consider to adapt the commands accordingly.
* Apply the deployment with Helm
  ```
  helm upgrade \
    --namespace hugo-mx-gateway \
    --install hugo-mx-gateway \
    helm/
  ```

* Check that the application is up and running.
  ```
  kubectl -n hugo-mx-gateway port-forward service/hugo-mx-gateway 8080:80
  curl http://127.0.0.1:8080/
  ```
  
The output in case of success shall be `{"status": "ok"}`.
