# How to build
## Build Operator Binary

```shell
$ make build-operator
```

## Build Image

Before building images, you should set the following environment variables:

* For Qiniu index-dev registry:
  * QiniuRegistry=index-dev.qiniu.io
  * QiniuRegistryPrefix=index-dev.qiniu.io/kelibrary
  * QiniuRegistryAK=XXX
  * QiniuRegistrySK=XXX

```shell
$ make image-operator
```

# How to use
## Deploy Operator

```shell
$ helm inspect ./helm/operator
$ helm install --set imageOperator.tag="XXX" --set resyncSeconds=300 ./helm operator
```

## Deploy a Qiniu CRD instance

```shell
$ helm inspect ./helm/service
$ helm install --set specData.name=flyer ./helm/service
```
