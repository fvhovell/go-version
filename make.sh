#!/bin/bash

export VERSION_PKG="github.com/fvhovell/go-version"
export GO_VERSION=$(go version | sed -e 's/go version //; s/ /-/g;')
export LINK_OPTIONS="\
	   -X ${VERSION_PKG}/version.AppName=version-test \
	   -X ${VERSION_PKG}/version.AppVersion=local-test-build \
	   -X ${VERSION_PKG}/version.BuildType=integration-build \
	   -X ${VERSION_PKG}/version.BuildTs=$(date +%FT%T%z) \
	   -X ${VERSION_PKG}/version.BuildHostname=$HOSTNAME \
	   -X ${VERSION_PKG}/version.BuildUsername=$USER \
	   -X ${VERSION_PKG}/version.BuildGoVersion=${GO_VERSION} \
	   -X ${VERSION_PKG}/version.GitCommit=$(cd $GOPATH/src/${VERSION_PKG} && git rev-parse HEAD) \
	   -X ${VERSION_PKG}/version.GitBranch=$(cd $GOPATH/src/${VERSION_PKG} && git symbolic-ref --short HEAD) \
"

go install -ldflags "${LINK_OPTIONS}" ${VERSION_PKG}/...
