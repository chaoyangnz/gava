#!/bin/bash

export GOPATH=`pwd`
go get github.com/urfave/cli
go build -o bin/jago
chmod +x bin/jago

export GOBIN=/usr/local/bin
go install

#go build -buildmode=c-shared -o build/libgvm.dylib launcher.go
#go tool cgo -srcdir src -objdir build/gni -exportheader build/gni.h -- -Isrc/gvm  gvm/gni.go