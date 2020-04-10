#! /bin/bash

set -e

kill $(lsof -t -i :8080) &
POD_NAME=$(kubectl get pods -l app=service-a -o json | jq -r .items[0].metadata.name)
kubectl port-forward $POD_NAME 8080 &
