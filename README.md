# Jago

A simplified Java interpreter written in Go language. This project is to learn JVM specification in depth and try to understand the
under-the-water mechanism when a Java program runs.

Basically, I only refer to "Java Virtual Machine Specification" and then compose how we should design one. 
I have to say many implementation details are naive and I make it as simplified as possible so as to
demonstrate the idea. For the educational purpose, it is enough.

# Roadmap

- Java class parser
- Interpreter
- Bridge JNI native method
- NO GC
- NO JIT
- NO multi-thread support

# Plan

- week01: class file reader refine
- week02: runtime structure (vm stack, stack frame)
- week03: type system and class loader
- week04: method area and heap object creation
- week05: instruction (1) constants, comparision, conversion
- week06: instruction (2) loads, stores, stack
- week07: instruction (3) reference, control, math
- week08: native methods (1) System, Object, Class, ClassLoader, Array, String
- week09: native methods (2) Throwable, StackTraceElement and exception handling
- week10: native methods (3) file io - FileInputStream, FileOutputStream
- week11: create test case

# How to run

## build and set gopath

```bash
❯ cd ~/jago
❯ export GOPATH=`pwd`
❯ export PATH=$PATH:`pwd`/bin
❯ ./build.sh
```

## jago command help
```bash
❯ jago -h
    _                   
   (_) __ _  __ _  ___  
   | |/ _` |/ _` |/ _ \ 
   | | (_| | (_| | (_) |   version 1.0.0
  _/ |\__,_|\__, |\___/    
 |__/       |___/     


Command: jago -h 

NAME:
   Jago - A simplified Java Virtual Machine for the educational purpose

USAGE:
   jago [-options] class [args...]

VERSION:
   1.0.0

DESCRIPTION:
   A Java Virtual Machine demonstrating the basic features of execution engine, class loading, type/value system, exception handling, native methods etc.

AUTHOR:
   Chao Yang <richd.yang@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --classpath value, --cp value  application classpath separated by colon
   --log:execution value          log level of instruction execution, options: info, debug, trace
   --log:classload value          log level of class loading, options: info, debug, trace
   --help, -h                     show help
   --version, -v                  print the version
```

## Run a calendar program

```bash
❯ jago --log:execution info -cp . example/Calendar 8 2017
    _                   
   (_) __ _  __ _  ___  
   | |/ _` |/ _` |/ _ \ 
   | | (_| | (_| | (_) |   version 1.0.0
  _/ |\__,_|\__, |\___/    
 |__/       |___/     


Command: jago --log:execution info -cp . example/Calendar 8 2017 

Add new classpath: .
------------------------------------------------------------

   August 2017
 S  M Tu  W Th  F  S
       1  2  3  4  5 
 6  7  8  9 10 11 12 
13 14 15 16 17 18 19 
20 21 22 23 24 25 26 
27 28 29 30 31 
```

## Run a pyramid program

```bash
❯ jago --log:execution info -cp . example/Pyramid        
    _                   
   (_) __ _  __ _  ___  
   | |/ _` |/ _` |/ _ \ 
   | | (_| | (_| | (_) |   version 1.0.0
  _/ |\__,_|\__, |\___/    
 |__/       |___/     


Command: jago --log:execution info -cp . example/Pyramid 

Add new classpath: .
------------------------------------------------------------

Pyramid pattern of star in Java : 
     * 
    * * 
   * * * 
  * * * * 
 * * * * * 
Pyramid of numbers in Java : 
     0 
    0 1 
   0 1 2 
  0 1 2 3 
 0 1 2 3 4 
```

## run a program traversing a tree in the level order

```bash
❯ jago --log:execution info -cp . example/TreeLevelOrderTraverse
    _                   
   (_) __ _  __ _  ___  
   | |/ _` |/ _` |/ _ \ 
   | | (_| | (_| | (_) |   version 1.0.0
  _/ |\__,_|\__, |\___/    
 |__/       |___/     


Command: jago --log:execution info -cp . example/TreeLevelOrderTraverse 

Add a new classpath: .
------------------------------------------------------------


    ___3__
   /      \
   9      20
         /  \
        15   7
[[3], [9, 20], [15, 7]]

    ___1__
   /      \
  _2       3
 /          \
 4           5
[[1], [2, 3], [4, 5]]

                _______________1______________
               /                              \
               2______                  _______3______
                      \                /              \
                       4            ___5__             6  
                                   /      \                
                                  _7       8                
                                 /                           
                                 9                            
[[1], [2, 3], [4, 5, 6], [7, 8], [9]]

```

## Run a program with exception

```bash
❯ jago --log:execution info -cp . test/test_athrow
    _                   
   (_) __ _  __ _  ___  
   | |/ _` |/ _` |/ _ \ 
   | | (_| | (_| | (_) |   version 1.0.0
  _/ |\__,_|\__, |\___/    
 |__/       |___/     


Command: jago --log:execution info -cp . test/test_athrow 

Add new classpath: .
------------------------------------------------------------

prepare to go

Exception in thread "main" test/Exception2: this is exception 2
         at test.test_athrow.rest(test_athrow.java:33)
         at test.test_athrow.onRoad(test_athrow.java:25)
         at test.test_athrow.go(test_athrow.java:20)
         at test.test_athrow.main(test_athrow.java:8)

```
# trace the execution

### log/execution.log 
this log file records the each instruction execution and method call hierarchy. Set the log level to info to view call hierarchy more clearly.
### log/classload.log 
this log file records the class loading process, including what triggers a class loading and when its class initialization method \<clinit\> is invoked.
### log/misc.log
other trivial logs 
# Runtime libraries solution (TBD)

![](Jago-jdk.png)