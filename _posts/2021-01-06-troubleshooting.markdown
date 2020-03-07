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

## Docker

#### [secure registry](https://stackoverflow.com/questions/34151612/docker-open-certs-domain-crt-permission-denied)

- symptom: ` open /certs/domain.crt: permission denied`
- diagnosis: `SELinux setup issue`

#### []

#### [can not push to private registry](https://docs.docker.com/engine/security/certificates/)

- symtom: `x509: certificate signed by unknown authority`

```
[root@rpulp-e2e 192.168.122.84:443]# docker push 192.168.122.84/busybox
The push refers to a repository [192.168.122.84/busybox]
Get https://192.168.122.84/v1/_ping: x509: certificate signed by unknown authority
```

- diagnosis: copy the cert for `REGISTRY_HTTP_TLS_CERTIFICATE` to `/etc/docker/certs.d/192.168.122.84`

#### `doesn't contain any IP SANs`

- symptom:  `Get https://192.168.122.84/v1/_ping: x509: cannot validate certificate for 192.168.122.84 because it doesn't contain any IP SANs`
- diagnosis: [subjectAltName](https://serverfault.com/questions/611120/failed-tls-handshake-does-not-contain-any-ip-sans) for `REGISTRY_HTTP_TLS_CERTIFICATE`

```
$ openssl genrsa -aes256 -out ca-key.pem 4096
$ openssl req -new -x509 -days 365 -key ca-key.pem -sha256 -out ca.pem
$ openssl genrsa -out server-key.pem 4096
$ openssl req -subj "/CN=rpulp-e2e" -sha256 -new -key server-key.pem -out server.csr
$ echo subjectAltName = DNS:rpulp-e2e,IP:192.168.122.84,IP:127.0.0.1 >> extfile.cnf
$ openssl x509 -req -days 365 -sha256 -in server.csr -CA ca.pem -CAkey ca-key.pem   -CAcreateserial -out server-cert.pem -extfile extfile.cnf
$ cp server-key.pem domain.key
$ cp server-cert.pem domain.crt 
```
