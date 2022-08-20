Kuberenetes Orchestration using Carvel VMWare

# Prerequisite

- Install a K8n cluster
- k3d cluster create --config k3dcluster.yml
- Run a simple docker registry
- docker run -d -p 5000:5000 --restart=always --name registry registry:2

# Simple App Walkthrough

## Ytt

ytt -f simple-demo -f ../../values.yml

## Kbld

kbld -f simple-demo

## Kapp

kapp deploy -a simple-demo -f- -c -y

## Chained

ytt -f simple-demo -f ../../values.yml | kbld -f - | kapp deploy -a simple-demo -f- -c -y

## Image Package

- ytt -f simple-demo -f ../../values.yml | kbld -f - --imgpkg-lock-output simple-demo/.imgpkg/images.yml

- imgpkg push -b ${REPO_HOST}/simple-demo:1.0.0 -f simple-demo/

- imgpkg pull -b ${REPO_HOST}/simple-demo:1.0.0 -o ../tmp/simple-demo

- Push to kubernetes

# YTT Library

cd ytt-example
ytt -f .
