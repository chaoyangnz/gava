package jago

/*
iload

== Operation

Load int from local variable

== Format

iload
index

== Forms

iload = 21 (0x15)

== Operand Stack

... →

..., value

== Description

The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
The local variable at index must contain an int.
The value of the local variable at index is pushed onto the operand stack.

== Notes

The iload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a
two-byte unsigned index.
 */
/*21 (0x15)*/
func ILOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.push(f.loadVar(uint(index)).(Int))
	f.nextPc()
}

/*
lload

== Operation

Load long from local variable

== Format

lload
index

== Forms

lload = 22 (0x16)

== Operand Stack

... →

..., value

== Description

The index is an unsigned byte. Both index and index+1 must be indices into the local variable array of the current
frame (§2.6). The local variable at index must contain a long. The value of the local variable at index is pushed onto
 the operand stack.

Notes

The lload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a
two-byte unsigned index.
 */
/*22 (0x16)*/
func LLOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.push(f.loadVar(uint(index)).(Long))
	f.nextPc()
}

/*
fload

== Operation

Load float from local variable

== Format

fload
index

== Forms

fload = 23 (0x17)

== Operand Stack

... →

..., value

== Description

The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
The local variable at index must contain a float.
The value of the local variable at index is pushed onto the operand stack.

== Notes

The fload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a
two-byte unsigned index.
 */
/*23 (0x17)*/
func FLOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.push(f.loadVar(uint(index)).(Float))
	f.nextPc()
}

/*
dload

== Operation

Load double from local variable

== Format

dload
index

== Forms

dload = 24 (0x18)

== Operand Stack

... →

..., value

== Description

The index is an unsigned byte. Both index and index+1 must be indices into the local variable array of the current frame (§2.6).
The local variable at index must contain a double. The value of the local variable at index is pushed onto the operand stack.

Notes

The dload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a
two-byte unsigned index.
 */
/*24 (0x18)*/
func DLOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.push(f.loadVar(uint(index)).(Double))
	f.nextPc()
}

/*
aload

== Operation

Load reference from local variable

== Format

aload
index

== Forms

aload = 25 (0x19)

== Operand Stack

... →

..., objectref

== Description

The index is an unsigned byte that must be an index into the local variable array of the current frame (§2.6).
The local variable at index must contain a reference.
The objectref in the local variable at index is pushed onto the operand stack.

== Notes

The aload instruction cannot be used to load a value of type returnAddress from a local variable onto the operand stack.
This asymmetry with the astore instruction (§astore) is intentional.

The aload opcode can be used in conjunction with the wide instruction (§wide) to access a local variable using a
two-byte unsigned index.
 */
/*25 (0x19)*/
func ALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex8()
	f.push(f.loadVar(uint(index)).(Reference))
	f.nextPc()
}

/*
iload_<n>

== Operation

Load int from local variable

== Format

iload_<n>

== Forms

iload_0 = 26 (0x1a)

iload_1 = 27 (0x1b)

iload_2 = 28 (0x1c)

iload_3 = 29 (0x1d)

== Operand Stack

... →

..., value

== Description

The <n> must be an index into the local variable array of the current frame (§2.6). The local variable at <n> must
contain an int. The value of the local variable at <n> is pushed onto the operand stack.

Notes

Each of the iload_<n> instructions is the same as iload with an index of <n>, except that the operand <n> is implicit.
 */

/*26 (0x1A)*/
func ILOAD_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(intCompatible(f.loadVar(0)))
	f.nextPc()
}

/*27 (0x1B)*/
func ILOAD_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(intCompatible(f.loadVar(1)))
	f.nextPc()
}

/*28 (0x1C)*/
func ILOAD_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(intCompatible(f.loadVar(2)))
	f.nextPc()
}

/*29 (0x1D)*/
func ILOAD_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(intCompatible(f.loadVar(3)))
	f.nextPc()
}

/*
lload_<n>

== Operation

Load long from local variable

== Format

lload_<n>

== Forms

lload_0 = 30 (0x1e)

lload_1 = 31 (0x1f)

lload_2 = 32 (0x20)

lload_3 = 33 (0x21)

== Operand Stack

... →

..., value

== Description

Both <n> and <n>+1 must be indices into the local variable array of the current frame (§2.6). The local variable at <n>
must contain a long. The value of the local variable at <n> is pushed onto the operand stack.

Notes

Each of the lload_<n> instructions is the same as lload with an index of <n>, except that the operand <n> is implicit.
 */

/*30 (0x1E)*/
func LLOAD_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(0).(Long))
	f.nextPc()
}

/*31 (0x1F)*/
func LLOAD_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(1).(Long))
	f.nextPc()
}

/*32 (0x20)*/
func LLOAD_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(2).(Long))
	f.nextPc()
}

/*33 (0x21)*/
func LLOAD_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(uint(3)).(Long))
	f.nextPc()
}

/*
fload_<n>

== Operation

Load float from local variable

== Format

fload_<n>

== Forms

fload_0 = 34 (0x22)

fload_1 = 35 (0x23)

fload_2 = 36 (0x24)

fload_3 = 37 (0x25)

== Operand Stack

... →

..., value

== Description

The <n> must be an index into the local variable array of the current frame (§2.6). The local variable at <n> must
contain a float. The value of the local variable at <n> is pushed onto the operand stack.

== Notes

Each of the fload_<n> instructions is the same as fload with an index of <n>, except that the operand <n> is implicit.
 */

/*34 (0x22)*/
func FLOAD_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(0).(Float))
	f.nextPc()
}

/*35 (0x23)*/
func FLOAD_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(1).(Float))
	f.nextPc()
}

/*36 (0x24)*/
func FLOAD_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(2).(Float))
	f.nextPc()
}

/*37 (0x25)*/
func FLOAD_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(3).(Float))
	f.nextPc()
}

/*
dload_<n>

== Operation

Load double from local variable

== Format

dload_<n>

== Forms

dload_0 = 38 (0x26)

dload_1 = 39 (0x27)

dload_2 = 40 (0x28)

dload_3 = 41 (0x29)

== Operand Stack

... →

..., value

== Description

Both <n> and <n>+1 must be indices into the local variable array of the current frame (§2.6). The local variable at <n>
must contain a double. The value of the local variable at <n> is pushed onto the operand stack.

== Notes

Each of the dload_<n> instructions is the same as dload with an index of <n>, except that the operand <n> is implicit.
 */

/*38 (0x26)*/
func DLOAD_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(0).(Double))
	f.nextPc()
}

/*39 (0x27)*/
func DLOAD_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(1).(Double))
	f.nextPc()
}

/*40 (0x28)*/
func DLOAD_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(2).(Double))
	f.nextPc()
}

/*41 (0x29)*/
func DLOAD_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(3).(Double))
	f.nextPc()
}

/*
aload_<n>

== Operation

Load reference from local variable

== Format

aload_<n>

== Forms

aload_0 = 42 (0x2a)

aload_1 = 43 (0x2b)

aload_2 = 44 (0x2c)

aload_3 = 45 (0x2d)

== Operand Stack

... →

..., objectref

== Description

The <n> must be an index into the local variable array of the current frame (§2.6). The local variable at <n> must
contain a reference. The objectref in the local variable at <n> is pushed onto the operand stack.

== Notes

An aload_<n> instruction cannot be used to load a value of type returnAddress from a local variable onto the operand
stack. This asymmetry with the corresponding astore_<n> instruction (§astore_<n>) is intentional.

Each of the aload_<n> instructions is the same as aload with an index of <n>, except that the operand <n> is implicit.
 */

/*42 (0x2A)*/
func ALOAD_0(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(0).(Reference))
	f.nextPc()
}

/*43 (0x2B)*/
func ALOAD_1(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(1).(Reference))
	f.nextPc()
}

/*44 (0x2C)*/
func ALOAD_2(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(2).(Reference))
	f.nextPc()
}

/*45 (0x2D)*/
func ALOAD_3(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.loadVar(3).(Reference))
	f.nextPc()
}

/*
iaload

== Operation

Load int from array

== Format

iaload

== Forms

iaload = 46 (0x2e)

== Operand Stack

..., arrayref, index →

..., value

== Description

The arrayref must be of type reference and must refer to an array whose components are of type int.
The index must be of type int. Both arrayref and index are popped from the operand stack. The int value in the component
 of the array at index is retrieved and pushed onto the operand stack.

== Run-time Exceptions

If arrayref is NULL, iaload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the iaload instruction throws an
ArrayIndexOutOfBoundsException.
 */
/*46 (0x2E)*/
func IALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)

	if arrayref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	if arrayref.Class().componentType != INT_TYPE {
		Bug("Not an int array")
	}
	if index < 0 || index >= arrayref.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}
	f.push(arrayref.GetElement(index).(Int))
	f.nextPc()
}

/*
laload

== Operation

Load long from array

== Format

laload

== Forms

laload = 47 (0x2f)

== Operand Stack

..., arrayref, index →

..., value

== Description

The arrayref must be of type reference and must refer to an array whose components are of type long.
The index must be of type int. Both arrayref and index are popped from the operand stack.
The long value in the component of the array at index is retrieved and pushed onto the operand stack.

== Run-time Exceptions

If arrayref is NULL, laload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the laload instruction
throws an ArrayIndexOutOfBoundsException.
 */
/*47 (0x2F)*/
func LALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	if arrayref.Class().componentType != LONG_TYPE {
		Bug("Not a long array")
	}
	if index < 0 || index >= arrayref.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}
	f.push(arrayref.GetElement(index).(Long))
	f.nextPc()
}

/*
faload

== Operation

Load float from array

== Format

faload

== Forms

faload = 48 (0x30)

== Operand Stack

..., arrayref, index →

..., value

== Description

The arrayref must be of type reference and must refer to an array whose components are of type float.
The index must be of type int. Both arrayref and index are popped from the operand stack.
The float value in the component of the array at index is retrieved and pushed onto the operand stack.

Run-time Exceptions

If arrayref is NULL, faload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the faload instruction
throws an ArrayIndexOutOfBoundsException.
 */
/*48 (0x30)*/
func FALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	if arrayref.Class().componentType != FLOAT_TYPE {
		Bug("Not an float array")
	}
	if index < 0 || index >= arrayref.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}
	f.push(arrayref.GetElement(index).(Float))
	f.nextPc()
}

/*
daload

== Operation

Load double from array

== Format

daload

== Forms

daload = 49 (0x31)

== Operand Stack

..., arrayref, index →

..., value

== Description

The arrayref must be of type reference and must refer to an array whose components are of type double.
The index must be of type int. Both arrayref and index are popped from the operand stack.
The double value in the component of the array at index is retrieved and pushed onto the operand stack.

== Run-time Exceptions

If arrayref is NULL, daload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the daload instruction
throws an ArrayIndexOutOfBoundsException.


 */
/*49 (0x31)*/
func DALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	if arrayref.Class().componentType != DOUBLE_TYPE {
		Bug("Not an double array")
	}
	if index < 0 || index >= arrayref.Length() {
		VM_throw("java/lang/java/lang/ArrayIndexOutOfBoundsException", "")
	}
	f.push(arrayref.GetElement(index).(Double))
	f.nextPc()
}

/*
aaload

== Operation

Load reference from array

== Format

aaload

== Forms

aaload = 50 (0x32)

== Operand Stack

..., arrayref, index →

..., value

== Description

The arrayref must be of type reference and must refer to an array whose components are of type reference.
The index must be of type int. Both arrayref and index are popped from the operand stack.
The reference value in the component of the array at index is retrieved and pushed onto the operand stack.

== Run-time Exceptions

If arrayref is NULL, aaload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the aaload instruction throws an
ArrayIndexOutOfBoundsException.
 */
/*50 (0x32)*/
func AALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	_, ok := arrayref.Class().componentType.(*Class)
	if !ok {
		Bug("Not an reference array")
	}
	if index < 0 || index >= arrayref.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}
	f.push(arrayref.GetElement(index).(Reference))
	f.nextPc()
}

/*
baload

== Operation

Load byte or boolean from array

== Format

baload

== Forms

baload = 51 (0x33)

== Operand Stack

..., arrayref, index →

..., value

== Description

The arrayref must be of type reference and must refer to an array whose components are of type byte or of type boolean.
The index must be of type int. Both arrayref and index are popped from the operand stack.
The byte value in the component of the array at index is retrieved, sign-extended to an int value,
and pushed onto the top of the operand stack.

== Run-time Exceptions

If arrayref is NULL, baload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the baload instruction throws
an ArrayIndexOutOfBoundsException.

== Notes

The baload instruction is used to load values from both byte and boolean arrays.
In Oracle's Java Virtual Machine implementation, boolean arrays - that is, arrays of type T_BOOLEAN (§2.2, §newarray) - are implemented as arrays of 8-bit values.
Other implementations may implement packed boolean arrays; the baload instruction of such implementations must be used to access those arrays.
 */
/*51 (0x33)*/
func BALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	if arrayref.Class().componentType != BYTE_TYPE && arrayref.Class().componentType != BOOLEAN_TYPE {
		Bug("Not a byte or boolean array")
	}
	if index < 0 || index >= arrayref.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}

	if arrayref.Class().componentType == BYTE_TYPE {
		b := arrayref.GetElement(index).(Byte)
		f.push(Int(b)) // sign-extended
	}
	if arrayref.Class().componentType == BOOLEAN_TYPE {
		b := arrayref.GetElement(index).(Boolean)
		f.push(Int(b)) // sign-extended
	}
	f.nextPc()
}

/*
caload

== Operation

Load char from array

== Format

caload

== Forms

caload = 52 (0x34)

== Operand Stack

..., arrayref, index →

..., value

== Description

The arrayref must be of type reference and must refer to an array whose components are of type char.
The index must be of type int. Both arrayref and index are popped from the operand stack.
The component of the array at index is retrieved and zero-extended to an int value.
That value is pushed onto the operand stack.

== Run-time Exceptions

If arrayref is NULL, caload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the caload instruction throws
an ArrayIndexOutOfBoundsException.
 */
/*52 (0x34)*/
func CALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	if arrayref.Class().componentType != CHAR_TYPE {
		Bug("Not a char array")
	}
	if index < 0 || index >= arrayref.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "%d is not in range [0, %d)", index, arrayref.Length())
	}
	ch := arrayref.GetElement(index).(Char)
	f.push(Int(ch)) //zero-extended to an int value
	f.nextPc()
}

/*
saload

== Operation

Load short from array

== Format

saload

== Forms

saload = 53 (0x35)

== Operand Stack

..., arrayref, index →

..., value

== Description

The arrayref must be of type reference and must refer to an array whose components are of type short.
The index must be of type int. Both arrayref and index are popped from the operand stack.
The component of the array at index is retrieved and sign-extended to an int value.
That value is pushed onto the operand stack.

== Run-time Exceptions

If arrayref is NULL, saload throws a NullPointerException.

Otherwise, if index is not within the bounds of the array referenced by arrayref, the saload instruction throws an ArrayIndexOutOfBoundsException.
 */
/*53 (0x35)*/
func SALOAD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.pop().(Int)
	arrayref := f.pop().(ArrayRef)
	if arrayref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}
	if arrayref.Class().componentType != SHORT_TYPE {
		Bug("Not a short array")
	}
	if index < 0 || index >= arrayref.Length() {
		VM_throw("java/lang/ArrayIndexOutOfBoundsException", "")
	}
	s := arrayref.GetElement(index).(Short)
	f.push(Int(s)) // sign-extended to an int value
	f.nextPc()
}