#!/bin/bash

cluster1=$1
cluster2=$2

kubectl config use-context $cluster1

kubectl create ns kube-multicluster-public
kubectl apply --validate=false -f vendor/k8s.io/cluster-registry/cluster-registry-crd.yaml

kubectl create clusterrolebinding federation-admin --clusterrole=cluster-admin \
    --serviceaccount="federation-system:default"
 
kubectl apply --validate=false -f hack/install-latest.yaml

for f in ./config/federatedtypes/*.yaml; do
    kubectl -n federation-system apply -f "${f}"
done

docker run -v $HOME/.kube:/root/.kube erstaples/kubefed2:latest join $cluster1 \
  --cluster-context $cluster1 \
  --host-cluster-context $cluster1 \
  --add-to-registry \
  --v=2

docker run -v $HOME/.kube:/root/.kube erstaples/kubefed2:latest join $cluster2 \
  --cluster-context $cluster2 \
  --host-cluster-context $cluster1 \
  --add-to-registry \
  --v=2
