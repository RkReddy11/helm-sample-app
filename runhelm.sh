#!/bin/bash
IMAGE_TAG=$1
cd ./charts
kubectl create ns rollback-poc
helmfile sync --set image.tag=$IMAGE_TAG --namespace rollback-poc
