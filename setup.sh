#!/bin/bash
export GOPATH=`pwd`:`pwd`/../..
export GOROOT=`which go`/../../
export PATH=$PATH:$GOPATH/bin
if [ -z `which glide` ]
then
    curl https://glide.sh/get | sh
fi
glide up
