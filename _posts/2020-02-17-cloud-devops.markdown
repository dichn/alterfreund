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
