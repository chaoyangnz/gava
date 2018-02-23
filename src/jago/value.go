package jago

/*
Value system:

<Value>
	|- Byte     -> int8
	|- Short    -> int16
	|- Char     -> uint16
	|- Int      -> int32
	|- Long     -> int64
	|- Float    -> float32
	|- Double   -> float64
	|- Boolean  -> int8
	|- Reference ( -> *Object)
		|= : ObjectRef
		|= : ArrayRef
	    |= : JavaLangString
	    |= : JavaLangThread
	    |= : JavaLangClass
	    |= : JavaLangClassLoader
	    |= : ...

ObjectRef and ArrayRef are only reference value holding a pointer to real heap object <Object> or <Array>.
The reference itself will be never nil, but its containing pointer can be nil, which means the reference is `NULL` in Java.
 */
type Value interface {
	Type() Type
}
