JVM Spec tells us: there are 8 primitive types, a reference type and a returnAddress.

I design the value system as following.
```
<Value>
|- Byte
|- Short
|- Char
|- Int
|- Long
|- Float
|- Double
|- Boolean
|- <ObjectRef> \
|- <JavaLangString> \
|- <JavaLangThread> \
|- <JavaLangClass> Reference ( -> *Object)
|- <JavaLangClassLoader> /
|- ... /
|- <ArrayRef> /
```

Primitive values are simple and each of them is directly mapped to a Go native type, like Java Int is mapped to int32 in Go while Double is mapped to float64. A reference is a Go struct with a pointer to an instance object. All these value is never nil, even a Java null is repented as a struct object with nil pointer. All these values are passed as value rather than pointer, which can be consistent with many semantics constraints in JVM specification.

Internally, to represent these values, I designed a type system accordingly.
```
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
```
I don’t differentiate the Array class (which is created internally by VM), non-array class and interface. All the subtypes of PrimitiveType should be singleton.