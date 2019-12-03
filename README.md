![](https://github.com/chaoyangnz/javo/workflows/build/badge.svg?branch=master)
![](https://img.shields.io/github/go-mod/go-version/chaoyangnz/javo?style=flat-square&branch=master)

# Javo ![](https://avatars0.githubusercontent.com/u/58212253?s=50&u=22ffd520fcf8842a796909f8ba760bb03438062b&v=4)

**NOTE**: (previously named `Jago`)

A simplified Java virtual machine written in Go language. One aim is to learn JVM specification in depth and try to understand the behind-the-scene behaviour when a Java program runs. This VM can be used for the educational purpose, for example, for a course about how to design a runtime for a language with the complete specification.

I only refer to "Java Virtual Machine Specification" and then look into how we should design one.
Some production-level features are intentionally ignored and it is supposed to make it as simplified as possible so as to
demonstrate the general idea. For the educational purpose, it is more than enough.

If you have no time to read OpenJDK source code or always guess the JVM behaviour when you need to tune your program, then your right here to be the lord of your universe.

A overview presentation can be found: https://drive.google.com/file/d/1XNlfhbCjPxynO1Io0qTVobc6W8wzYogt/view?usp=sharing

![](https://i.imgur.com/IWWtSYY.gif)

Any thought is welcome and I am happy to be wrong.

# Roadmap

- [x] Java class file reader
- [x] Interpreter engine
- [x] Class loader delegation
- [x] multi-threading support
- [x] monitor, `sleep`, `wait`, `notify` support
- [x] JDK native methods
- [ ] GC
- [ ] JIT

# How to run

## build and install

```bash
❯ cd ~/javo
❯ ./build.sh
```

By default, javo will be installed to `/usr/local/javo`
```text
/usr/local/javo
├── bin
│   └── javo
├── jdk
│   └── classes
│       ├── META-INF
│       ├── apple
│       ├── com
│       ├── java
│       ├── javax
│       ├── jdk
│       ├── org
│       └── sun
└── log
    ├── classloader.log
    ├── io.log
    ├── misc.log
    ├── thread-bootstrap.log
    ├── thread-main.log
    └── threads.log
```

## javo command
```bash
❯ javo -h
NAME:
   Javo - A simplified Java Virtual Machine for the educational purpose

USAGE:
   javo [-options] class [args...]

VERSION:
   1.0.0

DESCRIPTION:
   A Java Virtual Machine demonstrating the basic features of execution engine, class loading, type/value system, exception handling, native methods etc.

AUTHOR:
   Chao Yang <chaoyangnz@gmail.com>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --classpath value, --cp value  application classpath separated by colon
   --noLogo, --nl                 don't show logo
   --debug, -d                    debug mode
   --trace, -t                    trace mode
   --log:thread value             log level of instruction execution in a thread context, options: info, debug, trace
   --log:classloader value        log level of class loading, options: info, debug, trace
   --profile, -p                  profile javo
   --help, -h                     show help
   --version, -v                  print the version
```

### Run a calendar program

```bash
❯ javo -nl -cp . Calendar 8 2018
```
![](https://i.imgur.com/teTi3Ph.gif)

### Run a pyramid program

```bash
❯ javo -nl -cp . Pyramid
```
![](https://i.imgur.com/lJsHits.gif)

### run a program traversing a tree in the level order

```bash
❯ javo -nl -cp . TreeLevelOrderTraverse
```
![](https://i.imgur.com/1Oyq7Ep.gif)

### Run a program into exception

```bash
❯ javo -nl . test/test_athrow
```
![](https://i.imgur.com/lDxICzN.gif)

### More examples to demonstrate features

- Print system properties
```bash
javo -d test.test_system_properties
```

- Multi-threading
```bash
javo -d test.test_thread
```

```bash
javo -d Multithread
```

- Thread sleep or `wait` / `interrupt`
```bash
javo -d SleepInterrupt
```

```bash
javo -d WaitInterrupt
```

- Java monitor lock with `synchronized`
```bash
javo -d Counter
```

- Java wait() and `notify()`, `notifyAll()`
```bash
javo -d ProducerConsumer
```

## trace the execution

### log/thread-[name].log 
This log file records the each instruction execution and method call hierarchy within a thread context. Set the log level to info to view call hierarchy more clearly.

The blue diamond symbol 🔹 means pure Java methods while the yellow one 🔸 means native methods which is implemented within Javo internally.

The fire symbol 🔥 means exception throwing (thrown by `athrow` bytecode, rethrown if uncaught or thrown by VM internally), while the blue water symbol 💧 means an exception is caught by a method.

- TRACE: log all low level bytecode instructions, method call and exception thrown/caught
- DEBUG: only method call and exception thrown/caught
- INFO: only exception thrown/caught info

![](https://i.imgur.com/q0033Kl.gif)


### log/classloader.log 
This log file records the class loading process, including what triggers a class loading and when its class initialization method \<clinit\> is invoked.
The trigger reason can be viewed after a class name.

![](https://i.imgur.com/eN8foXd.gif)

### log/threads.log

This log file records thread creation/exit information and object monitor enter/exit, thread sleep(), wait(), notify() etc.

### log/io.log
This log file records I/O details around system interactions.

![](https://i.imgur.com/1eQe95o.gif)

### log/misc.log
Other trivial logs

![](https://i.imgur.com/2kgfhhb.gif)

# Profiling

https://flaviocopes.com/golang-profiling/

- run javo with `-p` option
- `go tool pprof --pdf /usr/local/bin/javo /var/../cpu.pprof > cpu.pdf`

# Documentation

Really sorry, I have no time to write documentation but it is in my plan.
The Ebook about how it works internally is in progress: https://www.gitbook.com/book/chaoyangnz/go-my-jvm

# Testing

I am really keen to get missing tests done if I have time. Another reason I didn't enjoy writing tests in Go is that I cannot find an awesome unit test library like Jest. For some libraries I looked into, it requires me to change my code just for testing. I don't like the idea about breaking the encapsulation just for testing. 
