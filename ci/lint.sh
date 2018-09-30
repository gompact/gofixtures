#!/bin/sh
p=$GOPATH/src/github.com/ishehata/gofixtures
mkdir -p $p
cp -R code/* $p/
cd $p
echo "calling go vet..."
go vet .
