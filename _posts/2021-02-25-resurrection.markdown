---
layout:      post
title:      "Mahler: Symphony No.2 Resurrection"
subtitle:   "Tom is working on piano"
date:        2021-02-25
author:     "Di Chen"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - cuda
  - hpc
  - installation
  - devops
---

## Installaion for Nvidia driver and CUDA toolkit

- [NVIDIA Drivers Install Guide](https://www.if-not-true-then-false.com/2015/fedora-nvidia-guide/)
- [NVIDIA Official Doc for CUDA](https://docs.nvidia.com/cuda/cuda-installation-guide-linux/index.html)
- [NVIDIA CUDA Toolkit](https://www.if-not-true-then-false.com/2018/install-nvidia-cuda-toolkit-on-fedora/)

```
// helloWorld.cu
#include <stdio.h>

__global__ void hello() {
    printf("Hello world from device");
}

int main() {
    hello<<<1, 1>>>();
    printf("Hello world from host");
    cudaDeviceSynchronize();
    return 0;
}

// nvcc -o helloWorld helloWorld.cu

```

## Others

- [VS Code](https://code.visualstudio.com/docs/setup/linux)
- [Golang](https://golang.org/doc/install?download=go1.16.linux-amd64.tar.gz)
- [vim tab setting](https://stackoverflow.com/questions/234564/tab-key-4-spaces-and-auto-indent-after-curly-braces-in-vim)

```
$ cat ~/.bashrc
...
export PATH="/usr/local/cuda-11.2/bin:$PATH"
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin

$ $ cat ~/.vimrc 
filetype plugin indent on
" show existing tab with 4 spaces width
set tabstop=4
" when indenting with '>', use 4 spaces width
set shiftwidth=4
" On pressing tab, insert 4 spaces
set expandtab
```
