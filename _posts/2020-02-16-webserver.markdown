---
layout:      post
title:      "Web Server Basics"
date:        2020-02-16
author:     "dichen16"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - web-server
---

## Apache Httpd

> Thanks to [Digitalocean Guide](https://www.digitalocean.com/community/tutorials/how-to-install-the-apache-web-server-on-centos-7) and [Secure Httpd](https://docs.fedoraproject.org/en-US/quick-docs/getting-started-with-apache-http-server/)

#### SELinux 

- get all SELinux Booleans: `getsebool -a` 
	-  `sudo semanage boolean -l` can **State  Default Description**
- print any security context of each file: `sudo ls -dZ /var/www/example.com/log/`
	- [-Z, --context]              print any security context of each file
	
> Apache has a type context of `httpd_t`. There is a policy rule that permits Apache access to files
and directories with the `httpd_sys_content_t` type context.

- semanage fcontext -a -t 
	- [-a, --add]             Add a record of the fcontext object type
	- [-t TYPE, --type TYPE]  SELinux Type for the object
- restorecon -R -v 

#### AllowOverride

> [Official Doc](https://httpd.apache.org/docs/2.4/mod/core.html#allowoverride)

I got this `The requested URL was not found on this server.` when `AllowOverride None`

#### Reminder

After adding new .conf file, remember to update the [local DNS resolver](https://en.wikipedia.org/wiki/Hosts_(file)).

```
# need the maps hostnames to IP addresse
# {tom.com} -> {ip_addr}
<VirtualHost *:80>
    ServerName www.tom.com
    ServerAlias tom.com
    DocumentRoot /var/www/example.com/html
    ErrorLog /var/www/example.com/log/error.log
    CustomLog /var/www/example.com/log/requests.log combined
</VirtualHost>

```

