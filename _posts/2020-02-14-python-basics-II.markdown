---
layout:      post
title:      "Python Basics II"
date:        2020-02-14
author:     "dichen16"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - python
  - notes
---

## Generator

> [PEP 255 -- Simple Generators](https://www.python.org/dev/peps/pep-0255/)

Here is a good [use case](https://realpython.com/introduction-to-python-generators/). Generator functions are a special kind of function that return a lazy iterator. Unlike lists, lazy iterators **do not store** their contents **in memory**.

Like list comprehensions, generator expressions allow you to quickly create a generator object in just a few lines of code. 
` nums_squared_lc = [num**2 for num in range(5)]`

## Command Line Script

> [Writing automated tests for Python command-line scripts](https://youtu.be/ApTZib0L2X8) 

## Method Types in Python OOP: @classmethod, @staticmethod, and Instance Methods

- @classmethod usage

```
class date:

    def __new__(cls, year, month=None, day=None):
        """Constructor.

        Arguments:

        year, month, day (required, base 1)
        """

    # Additional constructors

    @classmethod
    def fromtimestamp(cls, t):
        "Construct a date from a POSIX timestamp (like time.time())."
        y, m, d, hh, mm, ss, weekday, jday, dst = _time.localtime(t)
        return cls(y, m, d)

    @classmethod
    def today(cls):
        "Construct a date from time.time()."
        t = _time.time()
        return cls.fromtimestamp(t)

```

## `__init__`, `__new__`, and "cls"

> [why-is-init-always-called-after-new](https://stackoverflow.com/questions/674304/why-is-init-always-called-after-new)

## krbV

## ssl module 

> [Gentle introduction to TLS, PKI, and Python's ssl module - Christian Heimes](https://youtu.be/_YjX7rtiAsk)
