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

#### `TLS alert, unknown CA (560)`, `TLS alert, illegal parameter`, `SSLCertificate Chain`

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

#### LVM Expand logical volume and it's related file system

- sysptom: `resize2fs: Bad magic number in super-block while trying to open`
- diagnosis: [xfs_growfs rather than resize2fs](https://stackoverflow.com/questions/26305376/resize2fs-bad-magic-number-in-super-block-while-trying-to-open)

## Docker

- for a user want to pull or push image to a private registry, need to get a SSLCertificateFile(also is REGISTRY_HTTP_TLS_CERTIFICATE) from the remote registry server, and add it to `/etc/docker/certs.d/<remote_registry_ip_addr:port>` 

#### [secure registry](https://stackoverflow.com/questions/34151612/docker-open-certs-domain-crt-permission-denied)

- symptom: ` open /certs/domain.crt: permission denied`
- diagnosis: `SELinux setup issue`

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

## kubevirt

```
[dichen@dihost-amd kubevirt]$  make cluster-sync
./hack/cluster-build.sh
Building ...
go version go1.12.8 linux/amd64

a683dac26ea2216b2a6174c5c52bddc534a572124e631db08f71f671e6296750
go version go1.12.8 linux/amd64
Another command holds the client lock: 
pid=15
owner=client
cwd=/root/go/src/kubevirt.io/kubevirt

Waiting for it to complete...
Another command (pid=15) is running.  Waiting for it to complete on the server...
WARNING: Download from https://copr-be.cloud.fedoraproject.org/results/%40virtmaint-sig/for-kubevirt/fedora-30-x86_64/01034621-libvirt/libvirt-devel-5.0.0-2.fc30.x86_64.rpm failed: class com.google.devtools.build.lib.bazel.repository.downloader.UnrecoverableHttpException GET returned 404 Not Found
WARNING: Download from https://copr-be.cloud.fedoraproject.org/results/%40virtmaint-sig/for-kubevirt/fedora-30-x86_64/01034621-libvirt/libvirt-libs-5.0.0-2.fc30.x86_64.rpm failed: class com.google.devtools.build.lib.bazel.repository.downloader.UnrecoverableHttpException GET returned 404 Not Found
INFO: Call stack for the definition of repository 'org_golang_x_net' which is a go_repository (rule definition at /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_gazelle/internal/go_repository.bzl:187:17):
 - /root/go/src/kubevirt.io/kubevirt/WORKSPACE:589:1
INFO: Call stack for the definition of repository 'org_golang_google_grpc' which is a go_repository (rule definition at /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_gazelle/internal/go_repository.bzl:187:17):
 - /root/go/src/kubevirt.io/kubevirt/WORKSPACE:582:1
ERROR: An error occurred during the fetch of repository 'org_golang_x_net':
   failed to fetch org_golang_x_net: fetch_repo: unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
INFO: Call stack for the definition of repository 'remote_coverage_tools' which is a http_archive (rule definition at /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:292:16):
 - /DEFAULT.WORKSPACE.SUFFIX:9:1
INFO: Call stack for the definition of repository 'remote_java_tools_linux' which is a http_archive (rule definition at /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:292:16):
 - /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/utils.bzl:205:9
 - /DEFAULT.WORKSPACE.SUFFIX:260:1
INFO: Call stack for the definition of repository 'remotejdk11_linux' which is a http_archive (rule definition at /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/http.bzl:292:16):
 - /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/bazel_tools/tools/build_defs/repo/utils.bzl:205:9
 - /DEFAULT.WORKSPACE.SUFFIX:216:1
ERROR: /root/.cache/bazel/_bazel_root/6f347497f91c9a385dcd9294645b76e0/external/io_bazel_rules_go/proto/BUILD.bazel:21:1: @io_bazel_rules_go//proto:go_grpc depends on @org_golang_x_net//context:go_default_library in repository @org_golang_x_net which failed to fetch. no such package '@org_golang_x_net//context': failed to fetch org_golang_x_net: fetch_repo: unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
ERROR: Analysis of target '//cmd/virt-launcher:go_default_library' failed; build aborted: no such package '@org_golang_x_net//context': failed to fetch org_golang_x_net: fetch_repo: unrecognized import path "golang.org/x/net" (https fetch: Get https://golang.org/x/net?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)
INFO: Elapsed time: 37.532s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (707 packages loaded, 10660 targets configured)
    Fetching @libvirt; fetching 35s
    Fetching @fedora; fetching 35s
    Fetching @com_github_google_go_containerregistry; fetching 35s
make: *** [Makefile:102: cluster-build] Error 1
```
