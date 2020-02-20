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

## Max retries exceeded with url

```
(pipenv_G) [dichen@dilaptop Downloads]$ openstack image create --container-format bare --disk-format raw --file ./dichen-ss.raw dichen-tools
Failed to discover available identity versions when contacting https://<somewebsite.com>:<someport>/v3. Attempting to parse version from URL.
Unable to establish connection to https://<somewebsite.com>:<someport>/v3/auth/tokens: HTTPSConnectionPool(host='<somewebsite.com>', port=<someport>): Max retries exceeded with url: /v3/auth/tokens (Caused by ProxyError('Cannot connect to proxy.', OSError('Tunnel connection failed: 403 Forbidden')))
```

> [Turn off proxy solved my question, but don't know why](https://stackoverflow.com/questions/23013220/max-retries-exceeded-with-url-in-requests), I guess it's a [MTU](https://en.wikipedia.org/wiki/Maximum_transmission_unit) issue.
