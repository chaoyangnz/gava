package jago

/*
istore

== Operation

Store int into local variable

== Format

istore
index

== Forms

istore = 54 (0x36)

== Operand Stack

..., value →

...

== Description

The index is an unsigned byte that must be an index into the local variable array of the current this (§2.6).
The value on the top of the operand stack must be of type int.
It is popped from the operand stack, and the value of the local variable at index is set to value.

== Notes

The istore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.
 */
/*54 (0x36)*/
func ISTORE(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.storeVar(uint(index), f.pop().(Int))
	f.nextPc()
}

/*
lstore

== Operation

Store long into local variable

== Format

lstore
index

== Forms

lstore = 55 (0x37)

== Operand Stack

..., value →

...

== Description

The index is an unsigned byte. Both index and index+1 must be indices into the local variable array of the current this (§2.6). The value on the top of the operand stack must be of type long. It is popped from the operand stack, and the local variables at index and index+1 are set to value.

== Notes

The lstore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.
 */
/*55 (0x37)*/
func LSTORE(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.storeVar(uint(index), f.pop().(Long))
	f.nextPc()
}

/*
fstore

== Operation

Store float into local variable

== Format

fstore
index

== Forms

fstore = 56 (0x38)

== Operand Stack

..., value →

...

== Description

The index is an unsigned byte that must be an index into the local variable array of the current this (§2.6).
The value on the top of the operand stack must be of type float.
It is popped from the operand stack and undergoes value set conversion (§2.8.3), resulting in value'.
The value of the local variable at index is set to value'.

== Notes

The fstore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a
two-byte unsigned index.
 */
/*56 (0x38)*/
func FSTORE(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.storeVar(uint(index), f.pop().(Float))
	f.nextPc()
}

/*
dstore

== Operation

Store double into local variable

== Format

dstore
index

== Forms

dstore = 57 (0x39)

== Operand Stack

..., value →

...

== Description

The index is an unsigned byte. Both index and index+1 must be indices into the local variable array of the current this (§2.6).
The value on the top of the operand stack must be of type double.
It is popped from the operand stack and undergoes value set conversion (§2.8.3), resulting in value'.
The local variables at index and index+1 are set to value'.

== Notes

The dstore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a
two-byte unsigned index.
 */
/*57 (0x39)*/
func DSTORE(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.storeVar(uint(index), f.pop().(Double))
	f.nextPc()
}

/*
astore

== Operation

Store reference into local variable

== Format

astore
index

== Forms

astore = 58 (0x3a)

== Operand Stack

..., objectref →

...

== Description

The index is an unsigned byte that must be an index into the local variable array of the current this (§2.6).
The objectref on the top of the operand stack must be of type returnAddress or of type reference.
It is popped from the operand stack, and the value of the local variable at index is set to objectref.

Notes

The astore instruction is used with an objectref of type returnAddress when implementing the finally clause of the Java programming language (§3.13).

The aload instruction (§aload) cannot be used to load a value of type returnAddress from a local variable onto the operand stack.
This asymmetry with the astore instruction is intentional.

The astore opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a two-byte unsigned index.
 */
/*58 (0x3A)*/
func ASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()

	value := f.pop()
	switch value.(type) {
	case Reference: f.storeVar(uint(index), value.(Reference))
	case ReturnAddress: f.storeVar(uint(index), value.(ReturnAddress))
	}
	f.nextPc()
}

/*
istore_<n>

== Operation

Store int into local variable

== Format

istore_<n>

== Forms

istore_0 = 59 (0x3b)

istore_1 = 60 (0x3c)

istore_2 = 61 (0x3d)

istore_3 = 62 (0x3e)

== Operand Stack

..., value →

...

== Description

The <n> must be an index into the local variable array of the current this (§2.6).
The value on the top of the operand stack must be of type int. It is popped from the operand stack,
and the value of the local variable at <n> is set to value.

== Notes

Each of the istore_<n> instructions is the same as istore with an index of <n>, except that the operand <n> is implicit.
 */

/*59 (0x3B)*/
func ISTORE_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(0, f.pop().(Int))
	f.nextPc()
}

/*60 (0x3C)*/
func ISTORE_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(1, f.pop().(Int))
	f.nextPc()
}



/*61 (0x3D)*/
func ISTORE_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(2, assertIntCompatible(f.pop()))
	f.nextPc()
}

/*62 (0x3E)*/
func ISTORE_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(3, assertIntCompatible(f.pop()))
	f.nextPc()
}

/*
lstore_<n>

== Operation

Store long into local variable

== Format

lstore_<n>

== Forms

lstore_0 = 63 (0x3f)

lstore_1 = 64 (0x40)

lstore_2 = 65 (0x41)

lstore_3 = 66 (0x42)

== Operand Stack

..., value →

...

== Description

Both <n> and <n>+1 must be indices into the local variable array of the current this (§2.6).
The value on the top of the operand stack must be of type long. It is popped from the operand stack,
and the local variables at <n> and <n>+1 are set to value.

== Notes

Each of the lstore_<n> instructions is the same as lstore with an index of <n>, except that the operand <n> is implicit.
 */
/*63 (0x3F)*/
func LSTORE_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(0, f.pop().(Long))
	f.nextPc()
}

/*64 (0x40)*/
func LSTORE_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(1, f.pop().(Long))
	f.nextPc()
}

/*65 (0x41)*/
func LSTORE_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(2, f.pop().(Long))
	f.nextPc()
}

/*66 (0x42)*/
func LSTORE_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(3, f.pop().(Long))
	f.nextPc()
}

/*
fstore_<n>

== Operation

Store float into local variable

== Format

fstore_<n>

== Forms

fstore_0 = 67 (0x43)

fstore_1 = 68 (0x44)

fstore_2 = 69 (0x45)

fstore_3 = 70 (0x46)

== Operand Stack

..., value →

...

== Description

The <n> must be an index into the local variable array of the current this (§2.6).
The value on the top of the operand stack must be of type float.
It is popped from the operand stack and undergoes value set conversion (§2.8.3), resulting in value'.
The value of the local variable at <n> is set to value'.

Notes

Each of the fstore_<n> instructions is the same as fstore with an index of <n>, except that the operand <n> is implicit.
 */
/*67 (0x43)*/
func FSTORE_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(0, f.pop().(Float))
	f.nextPc()
}

/*68 (0x44)*/
func FSTORE_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(1, f.pop().(Float))
	f.nextPc()
}

/*69 (0x45)*/
func FSTORE_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(2, f.pop().(Float))
	f.nextPc()
}

/*70 (0x46)*/
func FSTORE_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(3, f.pop().(Float))
	f.nextPc()
}

/*
dstore_<n>

== Operation

Store double into local variable

== Format

dstore_<n>

== Forms

dstore_0 = 71 (0x47)

dstore_1 = 72 (0x48)

dstore_2 = 73 (0x49)

dstore_3 = 74 (0x4a)

Operand Stack

..., value →

...

== Description

Both <n> and <n>+1 must be indices into the local variable array of the current this (§2.6).
The value on the top of the operand stack must be of type double.
It is popped from the operand stack and undergoes value set conversion (§2.8.3), resulting in value'.
The local variables at <n> and <n>+1 are set to value'.

== Notes

Each of the dstore_<n> instructions is the same as dstore with an index of <n>, except that the operand <n> is implicit.
 */
/*71 (0x47)*/
func DSTORE_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(0, f.pop().(Double))
	f.nextPc()
}

/*72 (0x48)*/
func DSTORE_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(1, f.pop().(Double))
	f.nextPc()
}

/*73 (0x49)*/
func DSTORE_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(2, f.pop().(Double))
	f.nextPc()
}

/*74 (0x4A)*/
func DSTORE_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.storeVar(3, f.pop().(Double))
	f.nextPc()
}

/*
astore_<n>

== Operation

Store reference into local variable

== Format

astore_<n>

== Forms

astore_0 = 75 (0x4b)

astore_1 = 76 (0x4c)

astore_2 = 77 (0x4d)

astore_3 = 78 (0x4e)

== Operand Stack

..., objectref →

...

== Description

The <n> must be an index into the local variable array of the current this (§2.6).
The objectref on the top of the operand stack must be of type returnAddress or of type reference.
It is popped from the operand stack, and the value of the local variable at <n> is set to objectref.

== Notes

An astore_<n> instruction is used with an objectref of type returnAddress when implementing the finally clauses of the Java programming language (§3.13).

An aload_<n> instruction (§aload_<n>) cannot be used to load a value of type returnAddress from a local variable onto the operand stack.
This asymmetry with the corresponding astore_<n> instruction is intentional.

Each of the astore_<n> instructions is the same as astore with an index of <n>, except that the operand <n> is implicit.
 */

/*75 (0x4B)*/
func ASTORE_0(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop()
	switch value.(type) {
	case Reference: f.storeVar(0, value.(Reference))
	case ReturnAddress: f.storeVar(0, value.(ReturnAddress))
	}
	f.nextPc()
}

/*76 (0x4C)*/
func ASTORE_1(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop()
	switch value.(type) {
	case Reference: f.storeVar(1, value.(Reference))
	case ReturnAddress: f.storeVar(1, value.(ReturnAddress))
	}
	f.nextPc()
}

/*77 (0x4D)*/
func ASTORE_2(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop()
	switch value.(type) {
	case Reference: f.storeVar(2, value.(Reference))
	case ReturnAddress: f.storeVar(2, value.(ReturnAddress))
	}
	f.nextPc()
}

/*78 (0x4E)*/
func ASTORE_3(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop()
	switch value.(type) {
	case Reference: f.storeVar(3, value.(Reference))
	case ReturnAddress: f.storeVar(3, value.(ReturnAddress))
	}
	f.nextPc()
}

/*79 (0x4F)*/
func IASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.SetElement(index, value)
	//TODO check component type and boundary
	f.nextPc()
}

/*80 (0x50)*/
func LASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Long)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.SetElement(index, value)
	//TODO check component type and boundary
	f.nextPc()
}

/*81 (0x51)*/
func FASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Float)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.SetElement(index, value)
	//TODO check component type and boundary
	f.nextPc()
}

/*82 (0x52)*/
func DASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Double)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.SetElement(index, value)
	//TODO check component type and boundary
	f.nextPc()
}

/*83 (0x53)*/
func AASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(ObjectRef)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.SetElement(index, value)
	//TODO check component type, boundary and subtypes
	f.nextPc()
}

/*84 (0x54)*/
func BASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.SetElement(index, Byte(value))
	//TODO check component type and boundary
	f.nextPc()
}

/*85 (0x55)*/
func CASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Int)
	index := f.pop().(Int)

	arrayref := f.pop().(ArrayRef)
	arrayref.SetElement(index, Char(value))
	//TODO check component type and boundary
	f.nextPc()
}

/*86 (0x56)*/
func SASTORE(t *Thread, f *Frame, c *Class, m *Method) {
	value := f.pop().(Int)
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	arrayref.SetElement(index, Short(value))
	//TODO check component type and boundary
	f.nextPc()
}
