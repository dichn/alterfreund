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

> [Real Python's MockGuide](https://realpython.com/python-mock-library/) kind obsure,
> [mock](https://www.youtube.com/watch?v=ClAdw7ZJf5E)
> [mock usage](https://www.youtube.com/watch?v=WFRljVPHrkE)
> [autospec=True](https://docs.python.org/3/library/unittest.mock.html#autospeccing)

> When specifying the path to what you want to patch, you need to take care to point to **where it's being used, not where it's defined**

#### Multiple `patch()`

> [The order of your patches should be reversed](https://stackoverflow.com/questions/47042196/mock-patches-appearing-in-the-wrong-order)

## ssl module 

> [Gentle introduction to TLS, PKI, and Python's ssl module - Christian Heimes](https://youtu.be/_YjX7rtiAsk)

**ssl, TLS, cert stuff always confused me**

## Module, Packge, and `__init__.py`

A [module](https://docs.python.org/3/tutorial/modules.html) is a file containing Python definitions and statements.The file name is the module name with the suffix .py appended. 

Technically, a [package]() is a Python module with an __path__ attribute. 

## `pip` 

> WARNING: Running pip install with root privileges is generally not a good idea. Try `pip install --user` instead.

## Python Package Format, `wheel` v.s. `egg`

> [Wheel vs Egg](https://packaging.python.org/discussions/wheel-vs-egg/)

```bash
# install a .whl package
(pipenv_G) [dichen@dilaptop pipenv_G]$ pip install opencv_python-4.2.0.32-cp37-cp37m-manylinux1_x86_64.whl 
Processing ./opencv_python-4.2.0.32-cp37-cp37m-manylinux1_x86_64.whl
Collecting numpy>=1.14.5
  Downloading numpy-1.18.1-cp37-cp37m-manylinux1_x86_64.whl (20.1 MB)
     |████████████████████████████████| 20.1 MB 758 kB/s 
Installing collected packages: numpy, opencv-python
Successfully installed numpy-1.18.1 opencv-python-4.2.0.32

(pipenv_G) [dichen@dilaptop site-packages]$ pwd
/home/dichen/.local/share/virtualenvs/pipenv_G-utFzEIRm/lib/python3.7/site-packages
(pipenv_G) [dichen@dilaptop site-packages]$ ls opencv_python-4.2.0.32.dist-info/
INSTALLER  METADATA  RECORD  top_level.txt  WHEEL

(pipenv_G) [dichen@dilaptop site-packages]$ cat opencv_python-4.2.0.32.dist-info/RECORD 
cv2/.libs/libQtCore-ada04e4a.so.4.8.7,sha256=BM7GQ4FJK9zgNoa7C3-vKy2wGH5icqwH7YNE8w879Mw,3496040
cv2/.libs/libQtGui-903938cd.so.4.8.7,sha256=9zx6W-Rmc-oepRj2byDzdaYAZ8wk7aly570sJn0-SvI,13562584
cv2/.libs/libQtTest-1183da5d.so.4.8.7,sha256=Y20ba0MPzU8dz5g80-v0v2WfujXXnOIjpiAAFoBcHVU,195136
cv2/.libs/libavcodec-4cf96bc1.so.58.65.103,sha256=aN4_swPXs65vjyBwFBRhx_TbNEObw2K6oZ-2NfFSUoc,13277808
# omitted some RECORD
```
- [Python Egg usage](http://peak.telecommunity.com/DevCenter/PythonEggs#building-eggs) 
