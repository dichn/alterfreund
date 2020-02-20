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

> [why-is-init-always-called-after-new](https://stackoverflow.com/questions/674304/why-is-init-always-called-after-new) tells when to use `__new__()`, In general, you **shouldn't** need to override __new__` unless you're subclassing an immutable type like str, int, unicode or tuple.

- Called to create a new instance of class cls. __new__() is a **static method** (special-cased so you need not declare it as such) that takes the class of which an instance was requested as its first argument. 
- If __new__() is invoked during object construction and it returns an instance or subclass of cls, then the new instance’s __init__() method will be invoked like __init__(self[, ...]), where **self is the new instance** and the remaining **arguments are the same** as were passed to the object constructor.
- If __new__() does not return an instance of cls, then the new instance’s __init__() method **will not** be invoked.

```
def __new__(cls, year, month=None, day=None):
    """Constructor.
    Arguments:
    year, month, day (required, base 1)
    """
    if (month is None and
        isinstance(year, (bytes, str)) and len(year) == 4 and
        1 <= ord(year[2:3]) <= 12):
        # Pickle support
        if isinstance(year, str):
            try:
                year = year.encode('latin1')
            except UnicodeEncodeError:
                # More informative error message.
                raise ValueError(
                    "Failed to encode latin1 string when unpickling "
                    "a date object. "
                    "pickle.load(data, encoding='latin1') is assumed.")
        self = object.__new__(cls)
        self.__setstate(year)
        self._hashcode = -1
        return self
    year, month, day = _check_date_fields(year, month, day)
    self = object.__new__(cls)
    self._year = year
    self._month = month
    self._day = day
    self._hashcode = -1
    return self
```

## kerberos and requests-kerberos

> [requests-kerberos](https://github.com/requests/requests-kerberos) is an authentication handler for using Kerberos with Python Requests.
> [This is Apple's kerberos library](https://github.com/apple/ccs-pykerberos)

## The Python Mock Library

> [Real Python's MockGuide](https://realpython.com/python-mock-library/)





## ssl module 

> [Gentle introduction to TLS, PKI, and Python's ssl module - Christian Heimes](https://youtu.be/_YjX7rtiAsk)

**ssl, TLS, cert stuff always confused me**

## Module, Packge, and `__init__.py`

A [module](https://docs.python.org/3/tutorial/modules.html) is a file containing Python definitions and statements.The file name is the module name with the suffix .py appended. 

Technically, a [package]() is a Python module with an __path__ attribute. 
