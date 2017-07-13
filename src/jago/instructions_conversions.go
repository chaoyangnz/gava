package jago

import "fmt"

/*
i2l

== Operation

Convert int to long

== Format

i2l

== Forms

i2l = 133 (0x85)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type int. It is popped from the operand stack and sign-extended to
a long result. That result is pushed onto the operand stack.

== Notes

The i2l instruction performs a widening primitive conversion (JLS §5.1.2). Because all values of type int are exactly
representable by type long, the conversion is exact.
 */
/*133 (0x85)*/
func I2L(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)

	f.push(Long(value))
}

/*
i2f

== Operation

Convert int to float

== Format

i2f

== Forms

i2f = 134 (0x86)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type int. It is popped from the operand stack and converted to
the float result using IEEE 754 round to nearest mode. The result is pushed onto the operand stack.

== Notes

The i2f instruction performs a widening primitive conversion (JLS §5.1.2), but may result in a loss of precision
because values of type float have only 24 significand bits.
 */
/*134 (0x86)*/
func I2F(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)

	f.push(Float(value))
}

/*
i2d

== Operation

Convert int to double

== Format

i2d

== Forms

i2d = 135 (0x87)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type int. It is popped from the operand stack and converted to a double result. The result is pushed onto the operand stack.

== Notes

The i2d instruction performs a widening primitive conversion (JLS §5.1.2). Because all values of type int are exactly representable by type double, the conversion is exact.
 */
/*135 (0x87)*/
func I2D(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)

	f.push(Double(value))
}

/*
l2i

== Operation

Convert long to int

== Format

l2i

== Forms

l2i = 136 (0x88)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type long. It is popped from the operand stack and converted to an
int result by taking the low-order 32 bits of the long value and discarding the high-order 32 bits.
The result is pushed onto the operand stack.

== Notes

The l2i instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value. The result may also not have the same sign as value.
 */
/*136 (0x88)*/
func L2I(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Long)

	f.push(Int(value)) // truncate to 32 bits
}

/*
l2f

== Operation

Convert long to float

== Format

l2f

== Forms

l2f = 137 (0x89)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type long. It is popped from the operand stack and converted to
a float result using IEEE 754 round to nearest mode. The result is pushed onto the operand stack.

Notes

The l2f instruction performs a widening primitive conversion (JLS §5.1.2) that may lose precision because values of
type float have only 24 significand bits.
 */
/*137 (0x89)*/
func L2F(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Long)

	f.push(Float(value))
}

/*
l2d

== Operation

Convert long to double

== Format

l2d

== Forms

l2d = 138 (0x8a)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type long. It is popped from the operand stack and converted to a
double result using IEEE 754 round to nearest mode. The result is pushed onto the operand stack.

== Notes

The l2d instruction performs a widening primitive conversion (JLS §5.1.2) that may lose precision because values of
type double have only 53 significand bits.
 */
/*138 (0x8A)*/
func L2D(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Long)

	f.push(Double(value))
}

/*
f2i

== Operation

Convert float to int

== Format

f2i

== Forms

f2i = 139 (0x8b)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type float. It is popped from the operand stack and undergoes
value set conversion (§2.8.3), resulting in value'. Then value' is converted to an int result.
This result is pushed onto the operand stack:

- If the value' is NaN, the result of the conversion is an int 0.
- Otherwise, if the value' is not an infinity, it is rounded to an integer value V, rounding towards zero using IEEE 754
  round towards zero mode. If this integer value V can be represented as an int, then the result is the int value V.
- Otherwise, either the value' must be too small (a negative value of large magnitude or negative infinity), and the
 result is the smallest representable value of type int, or the value' must be too large (a positive value of large
 magnitude or positive infinity), and the result is the largest representable value of type int.

== Notes

The f2i instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value' and may also lose precision.
 */
/*139 (0x8B)*/
func F2I(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Float)

	f.push(Int(int32(float32(value))))
}

/*
f2l

== Operation

Convert float to long

== Format

f2l

== Forms

f2l = 140 (0x8c)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type float. It is popped from the operand stack and undergoes
value set conversion (§2.8.3), resulting in value'. Then value' is converted to a long result.
This result is pushed onto the operand stack:

- If the value' is NaN, the result of the conversion is a long 0.
- Otherwise, if the value' is not an infinity, it is rounded to an integer value V, rounding towards zero using IEEE 754
 round towards zero mode. If this integer value V can be represented as a long, then the result is the long value V.
- Otherwise, either the value' must be too small (a negative value of large magnitude or negative infinity), and the
 result is the smallest representable value of type long, or the value' must be too large (a positive value of large
 magnitude or positive infinity), and the result is the largest representable value of type long.

== Notes

The f2l instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value' and may also lose precision.
 */
/*140 (0x8C)*/
func F2L(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*
f2d

== Operation

Convert float to double

== Format

f2d

== Forms

f2d = 141 (0x8d)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type float. It is popped from the operand stack and undergoes value
 set conversion (§2.8.3), resulting in value'. Then value' is converted to a double result. This result is pushed onto
 the operand stack.

== Notes

Where an f2d instruction is FP-strict (§2.8.2) it performs a widening primitive conversion (JLS §5.1.2). Because all
values of the float value set (§2.3.2) are exactly representable by values of the double value set (§2.3.2), such a
conversion is exact.

Where an f2d instruction is not FP-strict, the result of the conversion may be taken from the double-extended-exponent
value set; it is not necessarily rounded to the nearest representable value in the double value set. However, if the
operand value is taken from the float-extended-exponent value set and the target result is constrained to the double
value set, rounding of value may be required.
 */
/*141 (0x8D)*/
func F2D(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*
d2i

== Operation

Convert double to int

== Format

d2i

== Forms

d2i = 142 (0x8e)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type double. It is popped from the operand stack and undergoes
value set conversion (§2.8.3) resulting in value'. Then value' is converted to an int. The result is pushed onto the
operand stack:

If the value' is NaN, the result of the conversion is an int 0.

Otherwise, if the value' is not an infinity, it is rounded to an integer value V, rounding towards zero using IEEE 754
round towards zero mode. If this integer value V can be represented as an int, then the result is the int value V.

Otherwise, either the value' must be too small (a negative value of large magnitude or negative infinity), and the
result is the smallest representable value of type int, or the value' must be too large (a positive value of large
magnitude or positive infinity), and the result is the largest representable value of type int.

== Notes

The d2i instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value' and may also lose precision.
 */
/*142 (0x8E)*/
func D2I(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*
d2l

== Operation

Convert double to long

== Format

d2l

== Forms

d2l = 143 (0x8f)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type double. It is popped from the operand stack and undergoes
value set conversion (§2.8.3) resulting in value'. Then value' is converted to a long. The result is pushed onto the
operand stack:
If the value' is NaN, the result of the conversion is a long 0.

Otherwise, if the value' is not an infinity, it is rounded to an integer value V, rounding towards zero using IEEE 754
round towards zero mode. If this integer value V can be represented as a long, then the result is the long value V.

Otherwise, either the value' must be too small (a negative value of large magnitude or negative infinity), and the
result is the smallest representable value of type long, or the value' must be too large (a positive value of large
magnitude or positive infinity), and the result is the largest representable value of type long.

== Notes

The d2l instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value' and may also lose precision.
 */
/*143 (0x8F)*/
func D2L(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*
d2f

== Operation

Convert double to float

== Format

d2f

== Forms

d2f = 144 (0x90)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type double. It is popped from the operand stack and undergoes
value set conversion (§2.8.3) resulting in value'. Then value' is converted to a float result using IEEE 754 round to
nearest mode. The result is pushed onto the operand stack.

Where an d2f instruction is FP-strict (§2.8.2), the result of the conversion is always rounded to the nearest
representable value in the float value set (§2.3.2).

Where an d2f instruction is not FP-strict, the result of the conversion may be taken from the float-extended-exponent
value set (§2.3.2); it is not necessarily rounded to the nearest representable value in the float value set.

A finite value' too small to be represented as a float is converted to a zero of the same sign; a finite value' too
large to be represented as a float is converted to an infinity of the same sign. A double NaN is converted to a
float NaN.

== Notes

The d2f instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value' and may also lose precision.
 */
/*144 (0x90)*/
func D2F(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	panic(fmt.Sprintf("Not implemented for opcode %d\n", opcode))
}

/*
i2b

== Operation

Convert int to byte

== Format

i2b

== Forms

i2b = 145 (0x91)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type int. It is popped from the operand stack, truncated to a byte,
 then sign-extended to an int result. That result is pushed onto the operand stack.

== Notes

The i2b instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value. The result may also not have the same sign as value.
 */
/*145 (0x91)*/
func I2B(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)

	b := Byte(value) // truncate to 8 bits
	f.push(Int(b)) // // sign-extended
}

/*
i2c

== Operation

Convert int to char

== Format

i2c

== Forms

i2c = 146 (0x92)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type int. It is popped from the operand stack, truncated to char,
then zero-extended to an int result. That result is pushed onto the operand stack.

== Notes

The i2c instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value. The result (which is always positive) may also not have the same sign as value.
 */
/*146 (0x92)*/
func I2C(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)

	ch := Char(int16(value)) // truncate to 16 bits
	f.push(Int(ch)) // 0-extended
}

/*
i2s

== Operation

Convert int to short

== Format

i2s

== Forms

i2s = 147 (0x93)

== Operand Stack

..., value →

..., result

== Description

The value on the top of the operand stack must be of type int. It is popped from the operand stack, truncated to a
short, then sign-extended to an int result. That result is pushed onto the operand stack.

== Notes

The i2s instruction performs a narrowing primitive conversion (JLS §5.1.3). It may lose information about the overall
magnitude of value. The result may also not have the same sign as value.
 */
/*147 (0x93)*/
func I2S(opcode uint8, t *Thread, f *Frame, c *Class, m *Method, jumped *bool) {
	value := f.pop().(Int)

	s := Short(value) // truncate to 16 bits
	f.push(Int(s)) // sign-extended
}