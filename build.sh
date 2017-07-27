#!/bin/bash

export GOPATH=`pwd`
go build -buildmode=c-shared -o build/libgvm.dylib launcher.go
go tool cgo -srcdir src -objdir build/gni -exportheader build/gni.h -- -Isrc/gvm  gvm/gni.go