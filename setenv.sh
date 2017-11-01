#!/bin/bash
BASE_GOPATH=$(dirname "$(dirname `pwd`)")
export GOPATH=`pwd`:${BASE_GOPATH}
export GOROOT=`which go`/../../