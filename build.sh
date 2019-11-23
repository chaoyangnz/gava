#!/bin/bash

export GOPATH=`pwd`
go build -o bin/javo
#export GOBIN=$JAVO_HOME/bin

if [ "$JAVO_HOME" == '' ]; then
    echo 'Cannot find JAVO_HOME, install to default location: /usr/local/javo'
    sudo mkdir -p /usr/local/javo
    sudo chown -R $USER /usr/local/javo
    mkdir -p /usr/local/javo/bin
    mkdir -p /usr/local/javo/log
    export JAVO_HOME=/usr/local/javo
    echo "export JAVO_HOME=$JAVO_HOME" >> ~/.bash_profile
    source ~/.bash_profile
fi

sudo cp bin/* $JAVO_HOME/bin/
sudo ln -sf $JAVO_HOME/bin/javo /usr/local/bin/javo

sudo cp -rf jdk $JAVO_HOME/

#go build -buildmode=c-shared -o build/libgvm.dylib launcher.go
#go tool cgo -srcdir src -objdir build/gni -exportheader build/gni.h -- -Isrc/gvm  gvm/gni.go
