#!/bin/sh
# build image
suffix="$1"
suffix=${suffix:=v1}

go build

image="ali-alert:$suffix"
echo -e "building image: $image\n"
tag="harbor.haodai.net/ops/$image"
docker build -t $tag .
docker push $tag
