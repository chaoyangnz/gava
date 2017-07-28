package jago

import (
	"math"
)

/*148 (0x94)*/
/*
== Operation

Compare long

== Format

lcmp

== Forms

lcmp = 148 (0x94)

== Operand Stack

..., value1, value2 →

..., result

== Description

Both value1 and value2 must be of type long. They are both popped from the operand stack, and a signed integer
comparison is performed. If value1 is greater than value2, the int value 1 is pushed onto the operand stack.
If value1 is equal to value2, the int value 0 is pushed onto the operand stack. If value1 is less than value2,
the int value -1 is pushed onto the operand stack.
 */
func LCMP(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Long)
	value1 := f.pop().(Long)

	if value1 > value2 {
		f.push(Int(1))
	} else if value1 == value2 {
		f.push(Int(0))
	} else if value1 < value2 {
		f.push(Int(-1))
	}

	f.nextPc()
}

/*
fcmp<op>

== Operation

Compare float

== Format

fcmp<op>

== Forms

fcmpg = 150 (0x96)

fcmpl = 149 (0x95)

== Operand Stack

..., value1, value2 →

..., result

== Description

Both value1 and value2 must be of type float. The values are popped from the operand stack and undergo value set
conversion (§2.8.3), resulting in value1' and value2'. A floating-point comparison is performed:

- If value1 is greater than value2, the int value 1 is pushed onto the operand stack.
- Otherwise, if value1 is equal to value2, the int value 0 is pushed onto the operand stack.
- Otherwise, if value1 is less than value2, the int value -1 is pushed onto the operand stack.
- Otherwise, at least one of value1 or value2 is NaN. The fcmpg instruction pushes the int value 1 onto the operand
  stack and the fcmpl instruction pushes the int value -1 onto the operand stack.

Floating-point comparison is performed in accordance with IEEE 754. All values other than NaN are ordered, with negative
 infinity less than all finite values and positive infinity greater than all finite values. Positive zero and negative
  zero are considered equal.

== Notes

The fcmpg and fcmpl instructions differ only in their treatment of a comparison involving NaN. NaN is unordered, so any
 float comparison fails if either or both of its operands are NaN. With both fcmpg and fcmpl available, any float
 comparison may be compiled to push the same result onto the operand stack whether the comparison fails on non-NaN
 values or fails because it encountered a NaN. For more information, see §3.5.
 */

/*149 (0x95)*/
func FCMPL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)

	if math.IsNaN(float64(value1)) || math.IsNaN(float64(value2)) {
		f.push(Int(-1))
	} else if value1 > value2 {
		f.push(Int(1))
	} else if value1 == value2 {
		f.push(Int(0))
	} else if value1 < value2 {
		f.push(Int(-1))
	}

	f.nextPc()
}

/*150 (0x96)*/
func FCMPG(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Float)
	value1 := f.pop().(Float)

	if math.IsNaN(float64(value1)) || math.IsNaN(float64(value2)) {
		f.push(Int(1))
	} else if value1 > value2 {
		f.push(Int(1))
	} else if value1 == value2 {
		f.push(Int(0))
	} else if value1 < value2 {
		f.push(Int(-1))
	}

	f.nextPc()
}

/*
dcmp<op>

== Operation

Compare double

== Format

dcmp<op>

== Forms

dcmpg = 152 (0x98)

dcmpl = 151 (0x97)

== Operand Stack

..., value1, value2 →

..., result

== Description

Both value1 and value2 must be of type double. The values are popped from the operand stack and undergo value set
conversion (§2.8.3), resulting in value1 and value2. A floating-point comparison is performed:

- If value1 is greater than value2, the int value 1 is pushed onto the operand stack.
- Otherwise, if value1 is equal to value2, the int value 0 is pushed onto the operand stack.
- Otherwise, if value1 is less than value2, the int value -1 is pushed onto the operand stack.
- Otherwise, at least one of value1 or value2 is NaN. The dcmpg instruction pushes the int value 1 onto the operand
 stack and the dcmpl instruction pushes the int value -1 onto the operand stack.

Floating-point comparison is performed in accordance with IEEE 754. All values other than NaN are ordered, with negative
infinity less than all finite values and positive infinity greater than all finite values. Positive zero and negative
zero are considered equal.

Notes

The dcmpg and dcmpl instructions differ only in their treatment of a comparison involving NaN. NaN is unordered, so any
double comparison fails if either or both of its operands are NaN. With both dcmpg and dcmpl available, any double
comparison may be compiled to push the same result onto the operand stack whether the comparison fails on non-NaN values
 or fails because it encountered a NaN. For more information, see §3.5.
 */

/*151 (0x97)*/
func DCMPL(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)

	if math.IsNaN(float64(value1)) || math.IsNaN(float64(value2)) {
		f.push(Int(-1))
	} else if value1 > value2 {
		f.push(Int(1))
	} else if value1 == value2 {
		f.push(Int(0))
	} else if value1 < value2 {
		f.push(Int(-1))
	}

	f.nextPc()
}

/*152 (0x98)*/
func DCMPG(t *Thread, f *Frame, c *Class, m *Method) {
	value2 := f.pop().(Double)
	value1 := f.pop().(Double)

	if math.IsNaN(float64(value1)) || math.IsNaN(float64(value2)) {
		f.push(Int(1))
	} else if value1 > value2 {
		f.push(Int(1))
	} else if value1 == value2 {
		f.push(Int(0))
	} else if value1 < value2 {
		f.push(Int(-1))
	}

	f.nextPc()
}

/*
if<cond>

== Operation

Branch if int comparison with zero succeeds

== Format

if<cond>
branchbyte1
branchbyte2

== Forms

ifeq = 153 (0x99)

ifne = 154 (0x9a)

iflt = 155 (0x9b)

ifge = 156 (0x9c)

ifgt = 157 (0x9d)

ifle = 158 (0x9e)

== Operand Stack

..., value →

...

== Description

The value must be of type int. It is popped from the operand stack and compared against zero. All comparisons are signed.
The results of the comparisons are as follows:

- ifeq succeeds if and only if value = 0
- ifne succeeds if and only if value ≠ 0
- iflt succeeds if and only if value < 0
- ifle succeeds if and only if value ≤ 0
- fgt succeeds if and only if value > 0
- ifge succeeds if and only if value ≥ 0

If the comparison succeeds, the unsigned branchbyte1 and branchbyte2 are used to construct a signed 16-bit offset, where
the offset is calculated to be (branchbyte1 << 8) | branchbyte2. Execution then proceeds at that offset from the address
of the opcode of this if<cond> instruction. The target address must be that of an opcode of an instruction within the
method that contains this if<cond> instruction.
Otherwise, execution proceeds at the address of the instruction following this if<cond> instruction.
 */

/*153 (0x99)*/
func IFEQ(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value := f.pop().(Int)

	if value == 0 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*154 (0x9A)*/
func IFNE(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value := f.pop().(Int)

	if value != 0 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*155 (0x9B)*/
func IFLT(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value := f.pop().(Int)

	if value < 0 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*156 (0x9C)*/
func IFGE(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value := f.pop().(Int)

	if value >= 0 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*157 (0x9D)*/
func IFGT(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value := f.pop().(Int)

	if value > 0 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*158 (0x9E)*/
func IFLE(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value := f.pop().(Int)

	if value <= 0 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*
if_icmp<cond>

== Operation

Branch if int comparison succeeds

== Format

if_icmp<cond>
branchbyte1
branchbyte2

== Forms

if_icmpeq = 159 (0x9f)

if_icmpne = 160 (0xa0)

if_icmplt = 161 (0xa1)

if_icmpge = 162 (0xa2)

if_icmpgt = 163 (0xa3)

if_icmple = 164 (0xa4)

== Operand Stack

..., value1, value2 →

...

== Description

Both value1 and value2 must be of type int. They are both popped from the operand stack and compared. All comparisons
are signed. The results of the comparison are as follows:

- if_icmpeq succeeds if and only if value1 = value2
- if_icmpne succeeds if and only if value1 ≠ value2
- if_icmplt succeeds if and only if value1 < value2
- if_icmple succeeds if and only if value1 ≤ value2
- if_icmpgt succeeds if and only if value1 > value2
- if_icmpge succeeds if and only if value1 ≥ value2

If the comparison succeeds, the unsigned branchbyte1 and branchbyte2 are used to construct a signed 16-bit offset, where
the offset is calculated to be (branchbyte1 << 8) | branchbyte2. Execution then proceeds at that offset from the address
of the opcode of this if_icmp<cond> instruction. The target address must be that of an opcode of an instruction within
the method that contains this if_icmp<cond> instruction.

Otherwise, execution proceeds at the address of the instruction following this if_icmp<cond> instruction.
 */

/*159 (0x9F)*/
func IF_ICMPEQ(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	if value1 == value2 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*160 (0xA0)*/
func IF_ICMPNE(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	if value1 != value2 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*161 (0xA1)*/
func IF_ICMPLT(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	if value1 < value2 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*162 (0xA2)*/
func IF_ICMPGE(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	if value1 >= value2 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*163 (0xA3)*/
func IF_ICMPGT(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	if value1 > value2 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*164 (0xA4)*/
func IF_ICMPLE(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value2 := f.pop().(Int)
	value1 := f.pop().(Int)

	if value1 <= value2 {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*
if_acmp<cond>

== Operation

Branch if reference comparison succeeds

== Format

if_acmp<cond>
branchbyte1
branchbyte2

== Forms

if_acmpeq = 165 (0xa5)

if_acmpne = 166 (0xa6)

== Operand Stack

..., value1, value2 →

...

== Description

Both value1 and value2 must be of type reference. They are both popped from the operand stack and compared.
The results of the comparison are as follows:

- if_acmpeq succeeds if and only if value1 = value2
- if_acmpne succeeds if and only if value1 ≠ value2

If the comparison succeeds, the unsigned branchbyte1 and branchbyte2 are used to construct a signed 16-bit offset,
where the offset is calculated to be (branchbyte1 << 8) | branchbyte2. Execution then proceeds at that offset from
the address of the opcode of this if_acmp<cond> instruction. The target address must be that of an opcode of an
instruction within the method that contains this if_acmp<cond> instruction.

Otherwise, if the comparison fails, execution proceeds at the address of the instruction following this if_acmp<cond>
instruction.
 */

/*165 (0xA5)*/
func IF_ACMPEQ(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value2 := f.pop().(Reference)
	value1 := f.pop().(Reference)

	if value1.IsEqual(value2) {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}

/*166 (0xA6)*/
func IF_ACMPNE(t *Thread, f *Frame, c *Class, m *Method) {
	offset := f.operandOffset16()
	value2 := f.pop().(Reference)
	value1 := f.pop().(Reference)

	if !value1.IsEqual(value2) {
		f.offsetPc(offset)
	} else {
		f.nextPc()
	}
}
