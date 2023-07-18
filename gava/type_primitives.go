package gava

type (
	ByteType    struct{ classObject JavaLangClass }
	ShortType   struct{ classObject JavaLangClass }
	CharType    struct{ classObject JavaLangClass }
	IntType     struct{ classObject JavaLangClass }
	LongType    struct{ classObject JavaLangClass }
	FloatType   struct{ classObject JavaLangClass }
	DoubleType  struct{ classObject JavaLangClass }
	BooleanType struct{ classObject JavaLangClass }

	ReturnAddressType struct{}
)

// their singletons
var (
	BYTE_TYPE    = &ByteType{}
	CHAR_TYPE    = &CharType{}
	SHORT_TYPE   = &ShortType{}
	INT_TYPE     = &IntType{}
	LONG_TYPE    = &LongType{}
	FLOAT_TYPE   = &FloatType{}
	DOUBLE_TYPE  = &DoubleType{}
	BOOLEAN_TYPE = &BooleanType{}

	RETURN_ADDRESS_TYPE = &ReturnAddressType{}
)

func (this *ByteType) Name() string                { return JVM_SIGNATURE_BYTE }
func (this *ShortType) Name() string               { return JVM_SIGNATURE_SHORT }
func (this *CharType) Name() string                { return JVM_SIGNATURE_CHAR }
func (this *IntType) Name() string                 { return JVM_SIGNATURE_INT }
func (this *LongType) Name() string                { return JVM_SIGNATURE_LONG }
func (this *FloatType) Name() string               { return JVM_SIGNATURE_FLOAT }
func (this *DoubleType) Name() string              { return JVM_SIGNATURE_DOUBLE }
func (this *BooleanType) Name() string             { return JVM_SIGNATURE_BOOLEAN }
func (this *ReturnAddressType) Name() string       { return "&" }
func (this *ByteType) Descriptor() string          { return JVM_SIGNATURE_BYTE }
func (this *ShortType) Descriptor() string         { return JVM_SIGNATURE_SHORT }
func (this *CharType) Descriptor() string          { return JVM_SIGNATURE_CHAR }
func (this *IntType) Descriptor() string           { return JVM_SIGNATURE_INT }
func (this *LongType) Descriptor() string          { return JVM_SIGNATURE_LONG }
func (this *FloatType) Descriptor() string         { return JVM_SIGNATURE_FLOAT }
func (this *DoubleType) Descriptor() string        { return JVM_SIGNATURE_DOUBLE }
func (this *BooleanType) Descriptor() string       { return JVM_SIGNATURE_BOOLEAN }
func (this *ReturnAddressType) Descriptor() string { return "&" }

func (this *ByteType) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}
func (this *ShortType) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}
func (this *CharType) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}
func (this *IntType) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}
func (this *LongType) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}
func (this *FloatType) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}
func (this *DoubleType) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}
func (this *BooleanType) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}
func (this *ReturnAddressType) ClassObject() JavaLangClass {
	Bug("Why does Java code need to access ReturnAddress??")
	return NULL // never should be accessed from Java code
}
