package jago

type Constant interface {}

type SymbolRef interface {
	Constant
	resolve()
}

type ClassRef struct {
	hostClass *Class
	className string
	class     ClassType // a class ref may be non-array class, array class or an interface
}

func (this *ClassRef) ResolvedClass() ClassType {
	if this.class == nil {
		this.resolve()
	}
	return this.class
}

func (this *ClassRef) resolve()  {
	this.class = BOOTSTRAP_CLASSLOADER.CreateClass(this.className)
}

type MemberRef struct {
	hostClass   *Class

	className    string
	class        *Class
	name        string
	descriptor  string
}

func (this *MemberRef) ResolvedClass() *Class {
	if this.class == nil {
		this.resolve()
	}
	return this.class
}

func (this *MemberRef) resolve() {
	this.class = BOOTSTRAP_CLASSLOADER.CreateClass(this.className).(*Class)
	//this.class.Link()
}

type FieldRef struct {
	MemberRef
	field *Field
}

func (this *FieldRef) ResolvedField() *Field {
	if this.field == nil {
		this.resolve()
	}
	return this.field
}

func (this *FieldRef) resolve() {
	class := this.ResolvedClass()
	this.field = class.FindField(this.name, this.descriptor)
}

type MethodRef struct {
	MemberRef
	method *Method
}

func (this *MethodRef) ResolvedMethod() *Method {
	if this.method == nil {
		this.resolve()
	}
	return this.method
}

func (this *MethodRef) resolve() {
	class := this.ResolvedClass()
	this.method = class.FindMethod(this.name, this.descriptor)
}

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func (this *InterfaceMethodRef) ResolvedMethod() *Method {
	if this.method == nil {
		this.resolve()
	}
	return this.method
}

func (this *InterfaceMethodRef) resolve() {
	//TODO
}

type StringConstant struct {
	hostClass *Class
	value     string
	jstring   ObjectRef
}

func (this *StringConstant) ResolvedString() ObjectRef {
	if this.jstring.IsNull() {
		this.resolve()
	}
	return this.jstring
}

func (this *StringConstant) resolve() {
	this.jstring = NewJavaLangString(this.value)
}


type UTF8Constant struct {
	hostClass   *Class
	value       string
}

type IntegerConstant struct {
	hostClass *Class
	value     Int
}

type LongConstant struct {
	hostClass *Class
	value     Long
}

type FloatConstant struct {
	hostClass *Class
	value     Float
}

type DoubleConstant struct {
	hostClass *Class
	value     Double
}

type NameAndTypeConstant struct {
	hostClass   *Class
	name        string
	descriptor  string
}

type MethodTypeConstant struct {
	hostClass   *Class
	descriptor   string
}

type MethodHandleConstant struct {
	hostClass      *Class
	referenceKind  u1
	referenceIndex u2
}

type InvokeDynamicConstant struct {
	hostClass   *Class
	bootstrapMethod          string
	name                     string
	descriptor               string
}