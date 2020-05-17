---
layout:      post
title:      "Iptables Basics"
date:        2020-05-16
author:     "dichen16"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - iptables
  - firewall
  - selinux
---

## iptables

> [Iptables appends rule default after “-A INPUT -j REJECT --reject-with icmp-host-prohibited”](https://serverfault.com/questions/525801/iptables-appends-rule-default-after-a-input-j-reject-reject-with-icmp-host)

> [how-to-list-and-delete-iptables-firewall-rules](https://www.digitalocean.com/community/tutorials/how-to-list-and-delete-iptables-firewall-rules)

- In RHEL 6, the iptables file locates in `/etc/sysconfig/iptables`
- In RHEL 7, the iptables file locates in `/etc/sysconfig/iptables`, but the Iptables feature is not included in RHEL 7 by default. Iptables is replaced with firewall-cmd.

#### CMDs

- `iptables -L --line-numbers`
- `iptables -vnL`
- add rules to iptables file, like `-A INPUT -p tcp -m tcp --dport 80 -j ACCEPT`
- `service iptables save` and `service iptables restart`

## SELinux

#### Situation: pub-hub can not show web-page because of SELinux

1. check SElinux log `/var/log/audit/audit.log`

```
type=AVC msg=audit(1589644580.157:123): avc:  denied  { name_connect } for  pid=1665 comm="httpd" dest=5432 scontext=system_u:system_r:httpd_t:s0 tcontext=system_u:object_r:postgresql_port_t
:s0 tclass=tcp_socket                                                                          
type=SYSCALL msg=audit(1589644580.157:123): arch=c000003e syscall=42 success=no exit=-13 a0=b a1=55d3bbbc1750 a2=10 a3=7fff832f6864 items=0 ppid=1657 pid=1665 auid=4294967295 uid=48 gid=48 e
uid=48 suid=48 fsuid=48 egid=48 sgid=48 fsgid=48 tty=(none) ses=4294967295 comm="httpd" exe="/usr/sbin/httpd" subj=system_u:system_r:httpd_t:s0 key=(null)
type=AVC msg=audit(1589644580.236:124): avc:  denied  { name_connect } for  pid=1665 comm="httpd" dest=5432 scontext=system_u:system_r:httpd_t:s0 tcontext=system_u:object_r:postgresql_port_t
:s0 tclass=tcp_socket
type=SYSCALL msg=audit(1589644580.236:124): arch=c000003e syscall=42 success=no exit=-13 a0=b a1=55d3bcd8cf60 a2=10 a3=7fff832f7524 items=0 ppid=1657 pid=1665 auid=4294967295 uid=48 gid=48 e
uid=48 suid=48 fsuid=48 egid=48 sgid=48 fsgid=48 tty=(none) ses=4294967295 comm="httpd" exe="/usr/sbin/httpd" subj=system_u:system_r:httpd_t:s0 key=(null)
type=AVC msg=audit(1589644580.243:125): avc:  denied  { name_connect } for  pid=1665 comm="httpd" dest=5432 scontext=system_u:system_r:httpd_t:s0 tcontext=system_u:object_r:postgresql_port_t
:s0 tclass=tcp_socket
type=SYSCALL msg=audit(1589644580.243:125): arch=c000003e syscall=42 success=no exit=-13 a0=b a1=55d3bcd8d6e0 a2=10 a3=7fff832f73b4 items=0 ppid=1657 pid=1665 auid=4294967295 uid=48 gid=48 e
uid=48 suid=48 fsuid=48 egid=48 sgid=48 fsgid=48 tty=(none) ses=4294967295 comm="httpd" exe="/usr/sbin/httpd" subj=system_u:system_r:httpd_t:s0 key=(null)
type=AVC msg=audit(1589644580.254:126): avc:  denied  { name_connect } for  pid=1665 comm="httpd" dest=5432 scontext=system_u:system_r:httpd_t:s0 tcontext=system_u:object_r:postgresql_port_t
:s0 tclass=tcp_socket
```

2. check related SELinux boolean by `# getsebool -a | grep httpd_can_network`

```
[root@mini-pub ~]# getsebool -a | grep httpd_can_network
httpd_can_network_connect --> off
httpd_can_network_connect_cobbler --> off
httpd_can_network_connect_db --> on
httpd_can_network_memcache --> off
httpd_can_network_relay --> off
```

3. adjust SELinux policy with booleans

```
[root@mini-pub ~]# setsebool -P httpd_can_network_connect_db 1
```
