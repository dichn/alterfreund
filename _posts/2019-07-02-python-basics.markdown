---
layout:      post
title:      "Python Basics"
subtitle:   "Beijing is so hot"
date:        2019-07-02
author:     "dichen16"
header-img: "img/home-bg.jpg"
catalog:     true
header-mask: 0.4
tags:
  - python
  - notes
---

## string

```python
print(help(str.lower))
```

## numeric data

```python
num = 3.14
print(type(num))

# Arithmetic operators:
# addition: 3 + 2
# subtraction: 3 -2
# multiplication: 3 * 2
# division: 3 / 2
# floor division: 3 // 2
# exponent: 3 ** 2
# mudulus: 3 % 2

print(abs(-3))
print(round(3.45)) # 四舍五入
```

#### casting

```python
num_1 = '100'
num_2 = '200'

num_1 = int(num_1)
num_2 = int(num_2)

print(num_1 + num_2)
```

## Lists

```python
# list
courses = ['history', 'math', 'phy', 'computer']
courses.append('art') # append
courses.insert(0, 'art') # insert in specific pos 
print(len(courses))
print(courses[-1])
print(courses[0:2]) # range operations

courses_2 = ['shit', 'chendi']
courses.extend(courses_2)

courses.remove('math')

```

1. add

2. remove

3. reverse           `lists.sort(reverse=True)`

4. sort                  `sorted_courses = sorted(courses)`

   1. in alphabetic
   2. in order

5. min and max

6. get index from element

7. check if elements exists in list     `print('art' in courses)`

8. traverse in `for` loop

   ```python
   for index, course in enumerate(courses, start = 1):
       print(index, course)
   ```

9. turn `list` into `string`

   ```python
   # list -> string
   courses_str = ', '.join(courses)
   # string -> list
   new_list = courses_str.split(', ')
   ```

## tuple

   > cannot modify tuple, in program, it is called mutable and immutable. Lists are mutable, tuples are not. empty_tuple = tuple()

   ```
   tuple_1 = ('history', 'math', 'phy')
   tuple_2 = tuple_1
   ```

## set 

   > do not care **order**, and throw away **duplicates**. empty_set = set()

   ```python
   set_1 = {'history', 'math', 'phy'}
   ```

   1. member check

      ```python
      print('math' in set_1)
      ```

   2. intersection

      ```python
      print(set_1.intersection(set_2))
      print(set_1.difference(set_2))
      print(set_1.union(set_2))
      ```

## dictionary

   > dictionary is allows work with key-value pairs, like hash map. key is a **unique identifier**.

   1. construct

      ```python
      student = {'name':'chendi', 'age':23, 'courses':['math', 'art']}
      print(student['name'])
      print(student.get('name'))
      print(student.get('phone','default_not_found'))
      ```

   2. key check

   3. add key-value pair

      ```python
      student['phone'] = '555-5555'
      student.update({'name':'Di', 'age':34, 'phone':'33-3333'})
      ```

   4. delete

      ```python
      del student['age']
      popped_age = student.pop('age')
      ```

   5. the length of a dictionary is how many keys this dictionary has.

   6. get all keys, values, items

      ```python
      student.keys()
      student.values()
      student.items()
      ```

   7. traverse

      ```python
      student = {'name':'chendi', 'age':23, 'courses':['math', 'art']}
      for key, value in student.items()
      	print(key, value)
      ```

## conditions

   ```python
   language = 'python'
   if language == 'python':
       print('language is python')
   elif language == 'java':
       print('language is java')
   else:
       print('no match')
   ```

#### logical Ops

   ```python
   # and 
   # or
   # not
   
   if user == "admin" and logged_in == True:
       print("x")
   else:
       print('y')
   
   if not logged_in:
       print('please log in')
   else:
       print('welcome')
   ```

#### id

   ```python
   a = [1, 2, 3]
   b = [1, 2, 3]
   # b = a
   print(id(a))
   print(id(b))
   print(a is b )
   ```

## Loop and Iteration

```python
for i in range(1, 11):
    print(i)
```

## Function

```python
def hello_func():
    return 'Hello Function.'

print(hello_func().upper())
#######################################
# parameter
def hello_func(greeting, name = 'You'):
    return '{}, {}'.format(greeting, name)

print(hello_func('Hi'))
#######################################
def student_info(*args, **kwargs):
    print(args)
    print(kwargs)
courses = ['math', 'art']
info= {'name':'Di', 'age':23}
student_info(*courses, **info)
```

## Module

> https://doc.python.org/3/py-modindex.html             **Modules** has two types: built-in modules and external modules

```python
# my_module.py
print('here is my module')

test = 'Test string'

def find_index(to_search, target):
    for i, value in enumerate(to_search):
        if value == target:
            return i

    return -1
```

```python
# intro.py
# import my_module 不能显示方法出自哪个包，不合理
from my_module import find_index, test
courses = ['history', 'art', 'math']
index =  find_index(courses, 'art')
print(index)
print(test)
```



#### module is not from the same directory

1. (deprecated)append the path of module to the sys.path
2. (deprecated)add the path of module to `~/.bash_profile`

#### pip

`pip install python-docx`

`pip uninstall python-docx`

## try-except block

```python
try:
    value = 10 / 0
    number = int(input("Enter a number: "))
    print(number)
# specify specific exception
except ZeroDivisionError as err:
    print(err)
# except ZeroDivisionError as err:
# print(err)
except ValueError:
    print("invalid error")
```

## reading from a external file

```python
my_file = open("file.txt", "r")	#r, w, a, r+ mode
# check if the file is readable
print(my_file.readable())	# the "readable" method returns a boolean

# spit all the info
print(my_file.read())

#read individial line
print(my_file.readline())	# return <type 'str'> and simultaneously move the cursor

for ln in my_file.readlines():
    print(ln)
    
my_file.close()
```

## writing to a file

```python
my_file = open("file.txt", "a")
#w mode: if file does not exist, there will be a new file created
my_file.write("\nTom - HR")	#attention

```

## Class ，Objects，and Class Function

```python
class Student:
    # initialize function
    def __init__(self, name, major, gpa, is_on_probation):
        self.name = name
        self.major = major
        self.gpa = gpa
        self.is_on_probation = is_on_probation
    def on_honor_roll(self):
        if self.gpa >= 2.4:
            return True
        else:
            return False
    
```

## Inheritance 

````python
from Chef import Chef

​```
假如Chef类有一个make_special_dish的方法，可以在子类中重写overwrite
回顾CPP：
overwrite and overload
overload：chendi（std::string name）,chendi(std::string name, int age), 也就是同名方法不同parameters
overwrite：继承下来的方法，父类和子类的同名方法不同实现
​```
class ChineseChef(Chef):
    def 
````

