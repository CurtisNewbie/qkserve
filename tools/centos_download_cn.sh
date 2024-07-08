#!/bin/bash

download_qkserve="https://mirror.ghproxy.com/https://github.com/CurtisNewbie/qkserve/releases/download/v0.0.1/qkserve_linux_v0.0.1"
qkserve="qkserve"
if [ ! -f "$qkserve" ]; then
    yum install -y curl procps \
        && curl "$download_qkserve" -L -o "$qkserve" \
        && chmod +x $qkserve
fi

