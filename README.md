# qkserve

For CentOS, e.g.,

```sh
yum install curl -y

# Download qkserve (linux) to /usr/local/bin
curl https://raw.githubusercontent.com/CurtisNewbie/qkserve/main/tools/linux_download.sh | bash

qkserve -file "my_file.txt" -port 80 -one-time
```

Alternatively, we can use a proxy for github:

```sh
yum install curl -y

curl https://ghfast.top/https://github.com/CurtisNewbie/qkserve/releases/download/v0.0.4/qkserve_linux_v0.0.4 -o qkserve && chmod +x qkserve && mv qkserve /usr/local/bin/qkserve
```
