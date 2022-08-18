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

## Create docker registry

- docker run -d -p 5000:5000 --restart=always --name registry registry:2
- export REPO_HOST="`ifconfig | grep -A1 docker | grep inet | cut -f10 -d' '`:5000"
- OR export REPO_HOST=<IP_ADDRESS>:5000
- curl ${REPO_HOST}/v2/\_catalog
- export CFG_repo_host=localhost:60630 (the same as REPO_HOST)

# Create app

## Add app

Use kbld to record which container images are used:

- kbld -f app-package/app/ --imgpkg-lock-output app-package/.imgpkg/images.yml

Now we can publish our bundle to our registry:

- imgpkg push -b ${REPO_HOST}/packages/simple-app:1.0.0 -f app-package/

You can verify that we pushed something called packages/simple-app:

- curl ${REPO_HOST}/v2/\_catalog

## Creating the Custom Resources

Create package metadata which contains high level information and descriptions about our package

- ytt -f app-package/app/values.yml --data-values-schema-inspect -o openapi-v3 > app-package/schema-openapi.yml
- See package.yml for more info.

## Creating a Package Repository

A package repository is a collection of packages and their metadata. Similar to a maven repository or a rpm repository, adding a package repository to a cluster gives users of that cluster the ability to install any of the packages from that repository.

### Setup

- mkdir -p package-repo/.imgpkg package-repo/packages/simple-app.corp.com
- ytt --data-values-env CFG -f app-package/package.yml --data-value-file openapi=app-package/schema-openapi.yml -v version="1.0.0" > package-repo/packages/simple-app.corp.com/1.0.0.yml
- cp app-package/metadata.yml package-repo/packages/simple-app.corp.com
- kbld -f package-repo/packages/ --imgpkg-lock-output package-repo/.imgpkg/images.yml

With the bundle metadata files present, we can push our bundle to whatever OCI registry we plan to distribute it from

- imgpkg push -b ${REPO_HOST}/packages/package-repo:1.0.0 -f package-repo
- curl ${REPO_HOST}/v2/\_catalog

## Adding a PackageRepository

- kapp deploy -a repo -f repo.yml -y
- watch kubectl get packagerepository
