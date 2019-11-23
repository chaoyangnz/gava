package javo

/*
Type system:

<Type>
  |- PrimitiveType
        |- *ByteType
        |- *ShortType
        |- *CharType
        |- *IntType
        |- *LongType
        |- *FloatType
        |- *DoubleType
        |- *BooleanType
  |- *ReturnAddressType
  |- *Class
 */

type Type interface {
	Name() string
	Descriptor() string
	ClassObject() JavaLangClass
}