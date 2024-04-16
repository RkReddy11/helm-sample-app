#!/bin/bash
IMAGE_TAG=$1
cd ./charts
helmfile sync --set image.tag=$IMAGE_TAG 
