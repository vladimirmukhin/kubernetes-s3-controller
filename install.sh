#!/bin/bash

make install
make docker-build docker-push IMG=public.ecr.aws/v6p5s4u6/s3controller:latest
make deploy IMG=public.ecr.aws/v6p5s4u6/s3controller:latest
kubectl -n s3controller-system scale deploy s3controller-controller-manager --replicas 0
kubectl -n s3controller-system scale deploy s3controller-controller-manager --replicas 1
kubectl -n s3controller-system delete -f config/samples/awscloud_v1_s3bucket.yaml
kubectl -n s3controller-system apply -f config/samples/awscloud_v1_s3bucket.yaml
until kubectl -n s3controller-system logs $(kubectl -n s3controller-system get pods | grep s3controller | awk '{print $1}') -c manager --follow
do
    sleep 3
    echo "Trying to fetch logs..."
done
