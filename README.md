# Jago

A simplified Java virtual machine written in Go language. One aim is to learn JVM specification in depth and try to understand the
behind-the-scene behaviour when a Java program runs.

I only refer to "Java Virtual Machine Specification" and then look into how we should design one.
Some production-level features are intentionally ignored and it is supposed to make it as simplified as possible so as to
demonstrate the general idea. For the educational purpose, it is more than enough.

If you have no time to read OpenJDK source code or always guess the JVM behaviour when you need to tune your program, then your right here to be the lord of your universe.

<img src="https://i.imgur.com/7L7XyqL.gif" width="500" />

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
❯ cd ~/jago
❯ ./build.sh
```

By default, jago will be installed to `/usr/local/jago`
```text
/usr/local/jago
├── bin
│   └── jago
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

## jago command
```bash
❯ jago -h
NAME:
   Jago - A simplified Java Virtual Machine for the educational purpose

USAGE:
   jago [-options] class [args...]

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
   --profile, -p                  profile jago
   --help, -h                     show help
   --version, -v                  print the version
```

### Run a calendar program

```bash
❯ jago --log:thread info -cp . Calendar 8 2018
```
<img src="https://i.imgur.com/7l58Qwe.gif" width="500" />

### Run a pyramid program

```bash
❯ jago --log:thread info -cp . Pyramid
```
<img src="https://i.imgur.com/AxFFw8K.gif" width="500" />

### run a program traversing a tree in the level order

```bash
❯ jago --log:thread info -cp . TreeLevelOrderTraverse
```
<img src="https://i.imgur.com/RDDvqLA.gif" width="500" />

### Run a program into exception

```bash
❯ jago --log:thread info -cp . test/test_athrow
```
<img src="https://i.imgur.com/Dz0vPnB.gif" width="500" />

### More examples to demonstrate features

- Print system properties
```bash
jago -d test.test_system_properties
```

- Multi-threading
```bash
jago -d test.test_thread
```

```bash
jago -d Multithread
```

- Thread sleep or `wait` / `interrupt`
```bash
jago -d SleepInterrupt
```

```bash
jago -d WaitInterrupt
```

- Java monitor lock with `synchronized`
```bash
jago -d Counter
```

- Java wait() and `notify()`, `notifyAll()`
```bash
jago -d ProducerConsumer
```

## trace the execution

### log/thread-[name].log 
This log file records the each instruction execution and method call hierarchy within a thread context. Set the log level to info to view call hierarchy more clearly.

The blue diamond symbol 🔹 means pure Java methods while the yellow one 🔸 means native methods which is implemented within Jago internally.

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

- run jago with `-p` option
- `go tool pprof --pdf /usr/local/bin/jago /var/../cpu.pprof > cpu.pdf`

# Documentation

Really sorry, I have no time to write documentation but it is in my plan.
The Ebook about how it works internally is in progress: https://www.gitbook.com/book/richdyang/go-my-jvm
