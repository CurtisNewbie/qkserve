# qkserve

For CentOS, e.g.,

```bash
# yum install -y curl

# Download qkserve (linux) to /usr/local/bin
curl https://raw.githubusercontent.com/CurtisNewbie/qkserve/main/tools/centos_download.sh | bash
# curl https://raw.githubusercontent.com/CurtisNewbie/qkserve/main/tools/centos_download_cn.sh | bash

qkserve -file "my_file.txt" -port 80 -one-time
```
