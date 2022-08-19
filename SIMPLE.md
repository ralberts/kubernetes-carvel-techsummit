# Simple App

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
