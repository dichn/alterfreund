---
layout:      post
title:      "Dvořák: Symphony No. 9 From the New World"
subtitle:   "Tom is working on piano"
date:        2021-02-27
author:     "Di Chen"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - make
  - makefile
  - bash
  - devops
---

## Makefile

- [GNU make manual](https://www.gnu.org/software/make/manual/make.html)

```
default: learn-make

# The shell Function:
# The shell function performs the same function that 
# backquotes (‘`’) perform in most shells: 
# it does command expansion.
define doubleddollar
	for file_exe in `find . -name "helloWorld*"`; do \
	echo $${file_exe}; \
	done
endef

# The call Function:
# Evaluate the variable var replacing any references to 
# $(1), $(2) with the first, second, etc. param values.
define varfunc
    echo $(1) $(2)
endef

1 = one

learn-make:
	@echo '@ = $@'
  @echo '1 = $1'
	@$(call varfunc, HelloTom, Pokemon)
	@$(call doubleddollar)

# Phony Targets: https://stackoverflow.com/questions/2145590/what-is-the-purpose-of-phony-in-a-makefile/2145605
.PHONY: learn-make

# Notes are bellow
# Each line in the recipe must start with a tab
```
