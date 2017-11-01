#!/bin/bash
export BASE_GOPATH=`dirname $(pwd)`
export GOPATH=`pwd`:$(dirname ${BASE_GOPATH})
echo $GOPATH
export GOROOT=`which go`/../../