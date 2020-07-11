---
layout:      post
title:      "CentOS7 部署MongoDB和相关C++扩展"
subtitle:   "当初被MongoCXX的调用规则折磨的日子"
date:        2019-06-03
author:     "dichen16"
header-img: "img/post-bg-os-metro.jpg"
catalog:     true
header-mask: 0.4
tags:
  - mongodb
  - devops
---

## 部署

> 官方文档 https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/

**编译mongo-c需要cmake版本至少3，所以先升级cmake**

```
./bootstrap 
gmake
make install
reboot #重启生效
```

**安装MongoDB**

`vim /etc/yum.repos.d/mongodb-org-3.6.repo`

```
[mongodb-org-3.6]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/3.6/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-3.6.asc
```

`yum update`

`yum install -y mongodb-org`

**开放相关端口**

```
firewall-cmd --zone=public --add-port=27017/tcp --permanent
firewall-cmd --reload
```

**启动服务**

`systemctl start mongod.service`

## 安装mongo-c driver

```bash
git clone https://github.com/mongodb/mongo-c-driver.git
cd mongo-c-driver/
mkdir cmake-build
cd cmake-build/
cmake -DENABLE_AUTOMATIC_INIT_AND_CLEANUP=OFF ..
make && make install
```

## 安装mongocxx driver

```bash
git clone https://github.com/mongodb/mongo-cxx-driver.git
cd mongo-cxx-driver/build
cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=/usr/local ..
make && make install
```

## 测试程序

```c++
#include <iostream>

#include <bsoncxx/builder/stream/document.hpp>
#include <bsoncxx/json.hpp>

#include <mongocxx/client.hpp>
#include <mongocxx/instance.hpp>

int main(int, char**) {
    mongocxx::instance inst{};
    mongocxx::client conn{mongocxx::uri{}};

    bsoncxx::builder::stream::document document{};

    auto collection = conn["testdb"]["testcollection"];
    document << "hello" << "world";

    collection.insert_one(document.view());
    auto cursor = collection.find({});

    for (auto&& doc : cursor) {
        std::cout << bsoncxx::to_json(doc) << std::endl;
    }
}
```



**编译**

`c++ --std=c++11 test.cpp -o test -I/usr/local/include/mongocxx/v_noabi -I/usr/local/include/bsoncxx/v_noabi  -L/usr/local/lib -lmongocxx -lbsoncxx`

