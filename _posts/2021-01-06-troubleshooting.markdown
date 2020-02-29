---
layout:      post
title:      "2020 Troubleshooting Record"
date:        2021-01-06
author:     "dichen16"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - troubleshooting
  - notes
---

## Others

#### Max retries exceeded with url

```
(pipenv_G) [dichen@dilaptop Downloads]$ openstack image create --container-format bare --disk-format raw --file ./dichen-ss.raw dichen-tools
Failed to discover available identity versions when contacting https://<somewebsite.com>:<someport>/v3. Attempting to parse version from URL.
Unable to establish connection to https://<somewebsite.com>:<someport>/v3/auth/tokens: HTTPSConnectionPool(host='<somewebsite.com>', port=<someport>): Max retries exceeded with url: /v3/auth/tokens (Caused by ProxyError('Cannot connect to proxy.', OSError('Tunnel connection failed: 403 Forbidden')))
```

> [Turn off proxy solved my question, but don't know why](https://stackoverflow.com/questions/23013220/max-retries-exceeded-with-url-in-requests), I guess it's a [MTU](https://en.wikipedia.org/wiki/Maximum_transmission_unit) issue.

#### `/tmp` dir

> [clear time](https://serverfault.com/questions/377348/when-does-tmp-get-cleared/377349)

#### TLS, `TLS alert, unknown CA (560)`, `TLS alert, illegal parameter`, `SSLCertificate Chain`

```bash
# (optional) Modify SSLCertificateChainFile
$ cat ca.crt >> chain.crt
$ cat server.crt >> chain.crt 

# modify /etc/httpd/conf.d/ssl.conf
    'SSLCertificateChainFile /etc/pki/example/chain.crt'

# [important] login to the client machine to require the certificate
$ openssl s_client -showcerts -servername rpulp-e2e -connect rpulp-e2e:443 > cacert.pem

# check the cert's detail
$ openssl x509 -noout -modulus -in ~/cacert.pem
```

- Pay attention to the certificate's `CN(common name)`, It's better to be the IP_ADDR
