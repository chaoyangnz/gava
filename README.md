# GVM

A simple Java interpreter written with Go language. This project is to learn JVM specification in depth and try to understand the
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

# Details

## trace classload and bytecode execution

```
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ⤈ Main
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ⤈ java/lang/Object
      ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
      ⤈ java/lang/Class
      ⤉ java/lang/Class
      ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ⤉ java/lang/Object
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ⤉ Main
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
⤮ Main.main([Ljava/lang/String;)V
   0000 ➢ iconst_2
   0001 ➢ istore_1
   0002 ➢ iconst_2
   0003 ➢ istore_2
   0004 ➢ iload_1
   0005 ➢ iload_2
   0006 ➢ iadd
   0007 ➢ istore_3
   0008 ➢ iload_3
   0009 ➢ invokestatic
⤮ Main.increase(I)I
   0000 ➢ iload_0
   0001 ➢ iconst_1
   0002 ➢ iadd
   0003 ➢ ireturn
⤮ Main.main([Ljava/lang/String;)V
   0012 ➢ istore
   0014 ➢ new
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ⤈ B
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
    ⤈ A
    ⤉ A
    ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ⤉ B
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
   0017 ➢ dup
   0018 ➢ invokespecial
⤮ B.<init>()V
   0000 ➢ aload_0
   0001 ➢ invokespecial
⤮ A.<init>()V
   0000 ➢ aload_0
   0001 ➢ invokespecial
⤮ java/lang/Object.<init>()V
   0000 ➢ return
⤮ A.<init>()V
   0004 ➢ return
⤮ B.<init>()V
   0004 ➢ return
⤮ Main.main([Ljava/lang/String;)V
   0021 ➢ astore
   0023 ➢ iconst_1
   0024 ➢ putstatic
   0027 ➢ aload
   0029 ➢ iconst_2
   0030 ➢ putfield
   0033 ➢ aload
   0035 ➢ iconst_3
   0036 ➢ putfield
   0039 ➢ aload
   0041 ➢ invokevirtual
⤮ B.foo()V
   0000 ➢ aload_0
   0001 ➢ getfield
   0004 ➢ istore_1
   0005 ➢ getstatic
   0008 ➢ istore_2
   0009 ➢ return
⤮ Main.main([Ljava/lang/String;)V
   0044 ➢ iload_1
   0045 ➢ i2d
   0046 ➢ ldc2_w
   0049 ➢ ddiv
   0050 ➢ dstore
   0052 ➢ ldc
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ⤈ java/lang/String
  ⤉ java/lang/String
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
   0054 ➢ invokestatic
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
  ⤈ GVM
  ⤉ GVM
  ‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧‧
😎 invoke native method GVM#print(Ljava/lang/String;)V 

⤮ Main.main([Ljava/lang/String;)V
   0057 ➢ return
```