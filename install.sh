#!/bin/bash

make install
make docker-build docker-push IMG=public.ecr.aws/v6p5s4u6/s3controller:latest
make deploy IMG=public.ecr.aws/v6p5s4u6/s3controller:latest
kubectl -n tutorial-system scale deploy s3controller-controller-manager --replicas 0
kubectl -n tutorial-system scale deploy s3controller-controller-manager --replicas 1
