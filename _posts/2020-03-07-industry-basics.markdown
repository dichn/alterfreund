---
layout:      post
title:      "Tom's Industry Basics"
date:        2020-03-07
author:     "dichen16"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - DevOps
  - Basics
---

## OS

#### LVM

- disk -> partion -> physical volume -> virtual group -> logical volume -> fs
- disk -> partion -> fs

**Mount with LVM logical volume/Disk partion**

> [Standard partition vs LVM physical volume](https://unix.stackexchange.com/questions/292327/create-partition-standard-partition-vs-lvm-physical-volume-in-centos-installat)

- `mkfs.ext4 /dev/<virtual group>/<logical volume>`
    - `mkfs.ext4 /dev/testvg/lv_data1`
- `mkfs.xfs /dev/<disk partion>`
    - `mkfs.xfs /dev/sdb1`

> [Delete a virtual group](https://www.thegeekdiary.com/centos-rhel-how-to-delete-a-volume-group-in-lvm/)

**Before removing the vg, deactivate it first**

#### Expand logical volume and the related file system

```bash
# prepare a standard partion already with dichen
[root@centos ~]# pvcreate /dev/sdc1
  Physical volume "/dev/sdc1" successfully created.

# extend virutal group
[root@centos ~]# vgextend centos /dev/sdc1
  Volume group "centos" successfully extended

# expand logical volume
[root@centos ~]# lvextend -L +8G /dev/centos/root
  Size of logical volume centos/root changed from 47.39 GiB (12133 extents) to 55.39 GiB (14181 extents).
  Logical volume centos/root successfully resized.

# error with resize2fs
[root@centos ~]# resize2fs /dev/centos/root
resize2fs 1.42.9 (28-Dec-2013)
resize2fs: Bad magic number in super-block while trying to open /dev/centos/root
Couldn't find valid filesystem superblock.

# xfs_frowfs rather than risize2fs
[root@centos ~]# xfs_growfs /dev/centos/root
meta-data=/dev/mapper/centos-root isize=512    agcount=9, agsize=1467648 blks
         =                       sectsz=512   attr=2, projid32bit=1
         =                       crc=1        finobt=0 spinodes=0
data     =                       bsize=4096   blocks=12424192, imaxpct=25
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0 ftype=1
log      =internal               bsize=4096   blocks=2866, version=2
         =                       sectsz=512   sunit=0 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0
data blocks changed from 12424192 to 14521344
```
