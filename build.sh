#!/bin/bash

version="v0.0.2"
GOOS=linux GOARCH=amd64 go build -o qkserve_linux_$version main.go