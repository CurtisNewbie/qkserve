#!/bin/bash

version="v0.0.4"
GOOS=linux GOARCH=amd64 go build -o qkserve_linux_$version main.go