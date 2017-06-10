#!/usr/bin/bash

go tool cgo -srcdir src -- -Isrc/gvm src/gvm gvm/gni.go