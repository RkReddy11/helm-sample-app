#!/bin/bash
IMAGE_TAG=$1
cd ./charts
kubectl create ns rollback-poc
helmfile sync --namespace rollback-poc --set image.tag=$IMAGE_TAG 
