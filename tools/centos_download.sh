#!/bin/bash

version="v0.0.1"
download_qkserve="https://github.com/CurtisNewbie/qkserve/releases/download/$version/qkserve_linux_$version"
tmp="/tmp/qkserve"
qkserve="/usr/local/bin/qkserve"
if [ ! -f "$qkserve" ]; then
    yum install -y curl procps \
        && curl "$download_qkserve" -L -o "$tmp" \
        && chmod +x $tmp \
        && mv $tmp $qkserve \
        && echo "qkserve installed at location: $qkserve"
fi
