---
layout:      post
title:      "Cloud DevOps"
date:        2020-02-17
author:     "dichen16"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - DevOps
  - Cloud
---

## OpenStack

#### Server Migration

> [Guide](https://docs.fuga.cloud/migrate-an-instance-from-one-openstack-to-another)

```bash
# get image info
$ openstack image list | grep dichen

# save the image
$ openstack image save --file <file_name> <image_id>

# upload the image to another provider
## download another RC file and source it
source <another .RC>

## upload the image from local file
$ openstack image create --container-format bare --disk-format qcow2 --file <image_file_path> <name_your_snapshot>
 
# launch the instance
$ openstack server create --flavor <flavor> --network <network_name> --security-group <name_of_sc> --image <snapshot_name>  <name_your_instance>
```

## Docker

- use docker without sudo

```bash
# add 
[root@dichen-tools dichen]# sudo groupadd docker
[root@dichen-tools dichen]# sudo usermod -aG docker $USER
[root@dichen-tools dichen]# newgrp docker 
# undo
[root@dichen-tools dichen]# gpasswd -d $USER docker
```

- Get repo info from a remote registry
    - `$ curl -X GET https://<IP_ADDR>:5000/v2/_catalog -k`
- [Protect the Docker daemon socket](https://docs.docker.com/engine/security/https/)
- [Protest the DOcker Registry](https://docs.docker.com/registry/insecure/#use-self-signed-certificates)

## KVM

- [Network Bridge](https://computingforgeeks.com/how-to-create-a-linux-network-bridge-on-rhel-centos-8/)
- [Network Interface Bonding](https://linuxconfig.org/how-to-configure-network-interface-bonding-on-red-hat-enterprise-linux-8)
- [KVM with non-root-user](https://computingforgeeks.com/use-virt-manager-as-non-root-user/)
- `touch /etc/cloud/cloud-init.disabled` if the image is exported from OpenStack

```bash
$ sudo virt-install --name=dichen --memory=2048 --vcpus=2 --file=/home/dichen/Downloads/iso/rpulp.qcow2 --network bridge=virbr0,model=virtio --graphics type=vnc,port=6900,listen=0.0.0.0 --import

$ virsh list --all
```

## Linux

#### RPM Package and Tom

- [`$ rpm -qf {file_path}`](https://unix.stackexchange.com/questions/4705/which-fedora-package-does-a-specific-file-belong-to) is used to query a file
- `repoquery -ql <package>`
- `package-cleanup --cleandupes`

## AWS

#### How to attach the Lambda to CloudFront Distro

> [ add triggers for CloudFront events to a Lambda function](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-edge-add-triggers-cf-console.html)
> [a Lambda function is triggered by a CloudFront viewer request event](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-cloudfront-trigger-events.html)

- pay attention to `Lambda Execution role` related with IAM

## Kubernetes

> [minikube on kvm](https://computingforgeeks.com/how-to-run-minikube-on-kvm/)
