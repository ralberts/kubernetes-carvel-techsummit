Kuberenetes Orchestration using Carvel VMWare

# Prerequisite

- Install a K8n cluster
- k3d cluster create --config k3dcluster.yml

# Initial Setup

- k3d cluster create carvel
- kapp deploy -a kc -f https://github.com/vmware-tanzu/carvel-kapp-controller/releases/download/v0.38.0/release.yml -y
- kubectl get all -n kapp-controller (check it out)
- kubectl api-resources --api-group packaging.carvel.dev (custom CR type)
- kubectl api-resources --api-group data.packaging.carvel.dev
- kubectl api-resources --api-group kappctrl.k14s.io

# Create app

## Add app

- Add config.yml and values.yml from app/
- mkdir -p package-contents/config/
- Move config files into package-contents/config/
- mkdir -p package-contents/.imgpkg
- kbld -f package-contents/config/ --imgpkg-lock-output package-contents/.imgpkg/images.yml

## Create docker registry

- docker run -d -p 5000:5000 --restart=always --name registry registry:2
- export REPO_HOST="`ifconfig | grep -A1 docker | grep inet | cut -f10 -d' '`:5000"
- OR export REPO_HOST=localhost:5000
- curl ${REPO_HOST}/v2/\_catalog

## Creating the Custom Resources

- Add metadata.yml
- ytt -f package-contents/config/values.yml --data-values-schema-inspect -o openapi-v3 > schema-openapi.yml
- Add package-template.yml

## Creating a Package Repository

A package repository is a collection of packages and their metadata. Similar to a maven repository or a rpm repository, adding a package repository to a cluster gives users of that cluster the ability to install any of the packages from that repository.

- mkdir -p my-pkg-repo/.imgpkg my-pkg-repo/packages/simple-app.corp.com
- ytt -f package-template.yml --data-value-file openapi=schema-openapi.yml -v version="1.0.0" > my-pkg-repo/packages/simple-app.corp.com/1.0.0.yml
- cp metadata.yml my-pkg-repo/packages/simple-app.corp.com
- kbld -f my-pkg-repo/packages/ --imgpkg-lock-output my-pkg-repo/.imgpkg/images.yml
- imgpkg push -b ${REPO_HOST}/packages/my-pkg-repo:1.0.0 -f my-pkg-repo
- curl ${REPO_HOST}/v2/\_catalog

## Adding a PackageRepository

- Add the repo.yml
- kapp deploy -a repo -f repo.yml -y
- watch kubectl get packagerepository
