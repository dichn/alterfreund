---
layout:      post
title:      "Mahler: Symphony No.2 Resurrection"
subtitle:   "Tom is working on piano"
date:        2021-01-06
author:     "dichen16 needs IMAGINATION"
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
