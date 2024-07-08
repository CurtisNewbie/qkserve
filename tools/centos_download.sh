#!/bin/bash

version="v0.0.3"
download_qkserve="https://github.com/CurtisNewbie/qkserve/releases/download/$version/qkserve_linux_$version"
tmp="/tmp/qkserve"
qkserve="/usr/local/bin/qkserve"

echo "Downloading qkserve@$version"
curl "$download_qkserve" -L -o "$tmp" \
    && chmod +x $tmp \
    && mv $tmp $qkserve \
    && echo "qkserve downloaded to location: $qkserve"

