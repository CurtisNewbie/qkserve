#!/bin/bash

download_qkserve="https://mirror.ghproxy.com/https://github.com/CurtisNewbie/qkserve/releases/download/v0.0.1/qkserve_linux_v0.0.1"
tmp="/tmp/qkserve"
qkserve="/usr/local/bin/qkserve"
if [ ! -f "$qkserve" ]; then
    yum install -y curl procps \
        && curl "$download_qkserve" -L -o "$tmp" \
        && chmod +x $tmp \
        && mv $tmp $qkserve \
        && echo "qkserve installed at location: $qkserve"
fi