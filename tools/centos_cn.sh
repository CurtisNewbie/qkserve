#!/bin/bash

download_qkserve="https://mirror.ghproxy.com/https://github.com/CurtisNewbie/qkserve/releases/download/v0.0.1/qkserve_linux_v0.0.1"
file="$1"
if [ ! -f "/tmp/qkserve" ]; then
    yum install -y curl procps \
        && curl "$download_qkserve" -L -o /tmp/qkserve \
        && chmod +x /tmp/qkserve
fi
/tmp/qkserve -file "$file" -port 443