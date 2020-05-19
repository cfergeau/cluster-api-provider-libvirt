#!/bin/sh

RUNTIME=${RUNTIME:-podman}

if [ "$IS_CONTAINER" != "" ]; then
  go vet "${@}"
else
  "$RUNTIME" run --rm \
    --env IS_CONTAINER=TRUE \
    --volume "${PWD}:/go/src/github.com/openshift/cluster-api-provider-libvirt:z" \
    --workdir /go/src/github.com/openshift/cluster-api-provider-libvirt \
    openshift/origin-release:golang-1.13 \
    ./hack/go-vet.sh "${@}"
fi;
