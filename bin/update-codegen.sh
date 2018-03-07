#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

REPO=github.com/prateeknayak/pg-kube-crd-controller
HEADER_FILE="${GOPATH}/src/${REPO}/bin/boilerplate.txt"

# there is a bug in codegen where deepcopy-gen still points to this location
# and client-gen does not like --go-header-file
mkdir -p ${GOPATH}/src/k8s.io/kubernetes/hack/boilerplate/
cp ${HEADER_FILE} "${GOPATH}/src/k8s.io/kubernetes/hack/boilerplate/boilerplate.go.txt"

pushd "${GOPATH}/src/${REPO}/vendor/k8s.io/code-generator"

./generate-groups.sh all \
    ${REPO}/pkg/client \
    ${REPO}/pkg/apis "sharedconfig:v1alpha1"

popd
