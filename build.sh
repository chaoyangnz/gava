#!/bin/bash

echo "Building Gava ..."
go mod download
go build -o bin/gava

if [ "$GAVA_HOME" == '' ]; then
    echo 'Cannot find GAVA_HOME, install to default location: /usr/local/gava'
    sudo mkdir -p /usr/local/gava
    sudo chown -R $USER /usr/local/gava
    mkdir -p /usr/local/gava/bin
    mkdir -p /usr/local/gava/log
    export GAVA_HOME=/usr/local/gava
    echo "export GAVA_HOME=$GAVA_HOME" >> ~/.bash_profile
    source ~/.bash_profile
fi

echo "Installing package to $GAVA_HOME ..."

sudo cp bin/* $GAVA_HOME/bin/
sudo ln -sf $GAVA_HOME/bin/gava /usr/local/bin/gava

sudo cp -rf jdk $GAVA_HOME/

echo "Done! 🍻"

gava -cp example HelloWorld