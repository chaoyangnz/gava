package jago

/*00 (0x00)*/
func NOP(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
}

/*01 (0x01)*/
/*
aconst_null

== Operation

Push NULL

== Format

aconst_null

== Forms

aconst_null = 1 (0x1)

== Operand Stack

... →

..., NULL

== Description

Push the NULL object reference onto the operand stack.

== Notes

The Java Virtual Machine does not mandate a concrete value for NULL.
 */
func ACONST_NULL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(NULL)
}

/*
iconst_<i>

== Operation

Push int constant

== Format

iconst_<i>

== Forms

iconst_m1 = 2 (0x2)

iconst_0 = 3 (0x3)

iconst_1 = 4 (0x4)

iconst_2 = 5 (0x5)

iconst_3 = 6 (0x6)

iconst_4 = 7 (0x7)

iconst_5 = 8 (0x8)

== Operand Stack

... →

..., <i>

== Description

Push the int constant <i> (-1, 0, 1, 2, 3, 4 or 5) onto the operand stack.

== Notes

Each of this family of instructions is equivalent to bipush <i> for the respective value of <i>, except that the operand <i> is implicit.
 */

/*02 (0x02)*/
func ICONST_M1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(-1))
}

/*03 (0x03)*/
func ICONST_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(0))
}

/*04 (0x04)*/
func ICONST_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(1))
}

/*05 (0x05)*/
func ICONST_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(2))
}

/*06 (0x06)*/
func ICONST_3(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(3))
}

/*07 (0x07)*/
func ICONST_4(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(4))
}

/*08 (0x08)*/
func ICONST_5(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Int(5))
}

/*
lconst_<l>

== Operation

Push long constant

== Format

lconst_<l>

== Forms

lconst_0 = 9 (0x9)

lconst_1 = 10 (0xa)

== Operand Stack

... →

..., <l>

== Description

Push the long constant <l> (0 or 1) onto the operand stack.
 */

/*09 (0x09)*/
func LCONST_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Long(0))
}

/*10 (0x0A)*/
func LCONST_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Long(1))
}

/*
fconst_<f>

== Operation

Push float

== Format

fconst_<f>

== Forms

fconst_0 = 11 (0xb)

fconst_1 = 12 (0xc)

fconst_2 = 13 (0xd)

== Operand Stack

... →

..., <f>

== Description

Push the float constant <f> (0.0, 1.0, or 2.0) onto the operand stack.
 */

/*11 (0x0B)*/
func FCONST_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Float(0.0))
}

/*12 (0x0C)*/
func FCONST_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Float(1.0))
}

/*13 (0x0D)*/
func FCONST_2(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Float(2.0))
}

/*
dconst_<d>

== Operation

Push double

== Format

dconst_<d>

== Forms

dconst_0 = 14 (0xe)

dconst_1 = 15 (0xf)

== Operand Stack

... →

..., <d>

== Description

Push the double constant <d> (0.0 or 1.0) onto the operand stack.
 */

/*14 (0x0E)*/
func DCONST_0(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Double(0.0))
}

/*15 (0x0F)*/
func DCONST_1(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	f.push(Double(1.0))
}

/*
bipush

== Operation

Push byte

== Format

bipush
byte

== Forms

bipush = 16 (0x10)

== Operand Stack

... →

..., value

== Description

The immediate byte is sign-extended to an int value. That value is pushed onto the operand stack.
 */
/*16 (0x10)*/
func BIPUSH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	b := f.const8(1)
	f.push(Int(b))
}

/*
sipush

== Operation

Push short

== Format

sipush
byte1
byte2

== Forms

sipush = 17 (0x11)

== Operand Stack

... →

..., value

== Description

The immediate unsigned byte1 and byte2 values are assembled into an intermediate short, where the value of the short is
(byte1 << 8) | byte2. The intermediate value is then sign-extended to an int value. That value is pushed onto the
operand stack.
 */
/*17 (0x11)*/
func SIPUSH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	s := Short(f.offset16())
	f.push(Int(s))
}

/*
ldc

== Operation

Push item from run-time constant pool

== Format

ldc
index

== Forms

ldc = 18 (0x12)

== Operand Stack

... →

..., value

== Description

The index is an unsigned byte that must be a valid index into the run-time constant pool of the current class (§2.6).
The run-time constant pool entry at index either must be a run-time constant of type int or float, or a reference to a
string literal, or a symbolic reference to a class, method type, or method handle (§5.1).

If the run-time constant pool entry is a run-time constant of type int or float, the numeric value of that run-time
constant is pushed onto the operand stack as an int or float, respectively.

Otherwise, if the run-time constant pool entry is a reference to an instance of class String representing a string
literal (§5.1), then a reference to that instance, value, is pushed onto the operand stack.

Otherwise, if the run-time constant pool entry is a symbolic reference to a class (§5.1), then the named class is
resolved (§5.4.3.1) and a reference to the Class object representing that class, value, is pushed onto the operand stack.

Otherwise, the run-time constant pool entry must be a symbolic reference to a method type or a method handle (§5.1).
The method type or method handle is resolved (§5.4.3.5) and a reference to the resulting instance of
java.lang.invoke.MethodType or java.lang.invoke.MethodHandle, value, is pushed onto the operand stack.

== Linking Exceptions

During resolution of a symbolic reference to a class, any of the exceptions pertaining to class resolution (§5.4.3.1)
can be thrown.

During resolution of a symbolic reference to a method type or method handle, any of the exception pertaining to method
type or method handle resolution (§5.4.3.5) can be thrown.

== Notes

The ldc instruction can only be used to push a value of type float taken from the float value set (§2.3.2) because a
constant of type float in the constant pool (§4.4.4) must be taken from the float value set.
 */
/*18 (0x12)*/
func LDC(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index8()

	cpInfo := c.constantPool[index]
	switch cpInfo.(type) {
	case *IntegerConstant:
		f.push(cpInfo.(*IntegerConstant).value)
	case *FloatConstant:
		f.push(cpInfo.(*FloatConstant).value)
	case *StringConstant:
		f.push(cpInfo.(*StringConstant).ResolvedString())
	case *ClassRef:
		f.push(cpInfo.(*ClassRef).ResolvedClass().ClassObject())
	case *MethodTypeConstant, *MethodHandleConstant:
		// TODO
	default:
		panic("Must be a run-time constant of type int or float, or a reference to a string literal, or a symbolic reference to a class, method type, or method handle")
	}
}

/*
ldc_w

== Operation

Push item from run-time constant pool (wide index)

== Format

ldc_w
indexbyte1
indexbyte2

== Forms

ldc_w = 19 (0x13)

== Operand Stack

... →

..., value

== Description

The unsigned indexbyte1 and indexbyte2 are assembled into an unsigned 16-bit index into the run-time constant pool of
the current class (§2.6), where the value of the index is calculated as (indexbyte1 << 8) | indexbyte2. The index must
be a valid index into the run-time constant pool of the current class. The run-time constant pool entry at the index
either must be a run-time constant of type int or float, or a reference to a string literal, or a symbolic reference to
a class, method type, or method handle (§5.1).

If the run-time constant pool entry is a run-time constant of type int or float, the numeric value of that run-time
constant is pushed onto the operand stack as an int or float, respectively.

Otherwise, if the run-time constant pool entry is a reference to an instance of class String representing a string
literal (§5.1), then a reference to that instance, value, is pushed onto the operand stack.

Otherwise, if the run-time constant pool entry is a symbolic reference to a class (§4.4.1). The named class is resolved
 (§5.4.3.1) and a reference to the Class object representing that class, value, is pushed onto the operand stack.

Otherwise, the run-time constant pool entry must be a symbolic reference to a method type or a method handle (§5.1).
The method type or method handle is resolved (§5.4.3.5) and a reference to the resulting instance of
java.lang.invoke.MethodType or java.lang.invoke.MethodHandle, value, is pushed onto the operand stack.

== Linking Exceptions

During resolution of the symbolic reference to a class, any of the exceptions pertaining to class resolution (§5.4.3.1)
can be thrown.

During resolution of a symbolic reference to a method type or method handle, any of the exception pertaining to method
type or method handle resolution (§5.4.3.5) can be thrown.

== Notes

The ldc_w instruction is identical to the ldc instruction (§ldc) except for its wider run-time constant pool index.

The ldc_w instruction can only be used to push a value of type float taken from the float value set (§2.3.2) because a
constant of type float in the constant pool (§4.4.4) must be taken from the float value set.
 */
/*19 (0x13)*/
func LDC_W(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index16()

	cpInfo := c.constantPool[index]
	switch cpInfo.(type) {
	case *IntegerConstant:
		f.push(cpInfo.(*IntegerConstant).value)
	case *FloatConstant:
		f.push(cpInfo.(*FloatConstant).value)
	case *StringConstant:
		f.push(cpInfo.(*StringConstant).ResolvedString())
	case *ClassRef:
		f.push(cpInfo.(*ClassRef).ResolvedClass().ClassObject())
	case *MethodTypeConstant, *MethodHandleConstant:
		// TODO
	default:
		panic("Must be a run-time constant of type int or float, or a reference to a string literal, or a symbolic reference to a class, method type, or method handle")
	}
}

/*
ldc2_w

== Operation

Push long or double from run-time constant pool (wide index)

== Format

ldc2_w
indexbyte1
indexbyte2

== Forms

ldc2_w = 20 (0x14)

== Operand Stack

... →

..., value

== Description

The unsigned indexbyte1 and indexbyte2 are assembled into an unsigned 16-bit index into the run-time constant pool of
the current class (§2.6), where the value of the index is calculated as (indexbyte1 << 8) | indexbyte2. The index must
be a valid index into the run-time constant pool of the current class. The run-time constant pool entry at the index
must be a run-time constant of type long or double (§5.1). The numeric value of that run-time constant is pushed onto
the operand stack as a long or double, respectively.

== Notes

Only a wide-index version of the ldc2_w instruction exists; there is no ldc2 instruction that pushes a long or double
with a single-byte index.

The ldc2_w instruction can only be used to push a value of type double taken from the double value set (§2.3.2) because
a constant of type double in the constant pool (§4.4.5) must be taken from the double value set.
 */
/*20 (0x14)*/
func LDC2_W(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	index := f.index16()

	cpInfo := c.constantPool[index]
	switch cpInfo.(type) {
	case *LongConstant:
		f.push(cpInfo.(*LongConstant).value)
	case *DoubleConstant:
		f.push(cpInfo.(*DoubleConstant).value)
	default:
		panic("Not long or double in constant pool")
	}
}