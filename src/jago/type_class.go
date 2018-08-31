package jago

import (
	"sync"
	"strings"
)

const (
	FAILED        = -1
	UNINITIALIZED = 0
	INITIALIZING  = 1
	INITIALIZED   = 2
)

type Class struct {
	// shared
	name           string
	accessFlags    uint16
	superClassName string
	interfaceNames []string
	superClass     *Class
	interfaces     []*Class

	classObject JavaLangClass
	//classLoader         JavaLangClassLoader

	// support link and initialization
	defined bool
	linked  bool

	// these 3 flag is closely related class initialization
	initialized int
	T           *Thread
	LC          *sync.Cond

	// ---- these fields are only for non-array class ----
	constantPool []Constant
	fields       []*Field
	methods      []*Method

	maxInstanceVars   int
	instanceVarFields []*Field
	maxStaticVars     int
	staticVars        []Value
	staticVarFields   []*Field

	sourceFile string

	// ---- these fields are only for array class -------
	componentType Type // any type
	elementType   Type // must be not array type
	dimensions    int
}

func (this *Class) Name() string {
	return this.name
}

func (this *Class) ClassObject() JavaLangClass {
	if this.classObject.IsNull() {
		this.classObject = VM.NewJavaLangClass(this)
	}
	return this.classObject
}

func (this *Class) Descriptor() string {
	return JVM_SIGNATURE_CLASS + this.Name() + JVM_SIGNATURE_ENDCLASS
}

func (this *Class) IsInterface() bool {
	return this.accessFlags&JVM_ACC_INTERFACE > 0
}

func (this *Class) IsArray() bool {
	return string(this.name[0]) == JVM_SIGNATURE_ARRAY
}

func (this *Class) IsAssignableFrom(class *Class) bool {
	if this == class {
		return true
	}
	// this is interface
	if this.IsInterface() { // check whether this is within class's interfaces list
		clazz := class
		for clazz != nil {
			interfaces := clazz.interfaces
			for _, interface0 := range interfaces {
				if interface0 == this {
					return true
				}
				interfaces = append(interfaces, interface0.interfaces...)
			}
			clazz = clazz.superClass
		}

	} else if this.IsArray() {
		if class.IsArray() {
			thisComponentType, ok1 := this.componentType.(*Class)
			classComponentType, ok2 := class.componentType.(*Class)
			if ok1 && ok2 { // covariant
				return thisComponentType.IsAssignableFrom(classComponentType)
			}
		}
	} else { // non-array class
		if class.IsArray() {
			if class.superClass == this {
				return true
			}
		} else {
			clazz := class
			if clazz.IsInterface() { // interface disallow
				return false
			}
			for clazz != nil {
				if clazz == this {
					return true
				}
				clazz = clazz.superClass
			}
		}
	}

	return false
}

/**
 * Find field with its name and descriptor in the class hierarchy
 * Java doesn't permit a static and instance field with same name + signature
 */
func (this *Class) FindField(name string, descriptor string) *Field {
	class := this
	for class != nil {
		for _, field := range class.fields {
			if field.name == name && field.descriptor == descriptor {
				return field
			}
		}
		class = class.superClass
	}
	return nil
}

// Only find its own defined method
func (this *Class) GetMethod(name string, descriptor string) *Method {
	for _, method := range this.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

/*
<clinit>()V method
 */
func (this *Class) GetClassInitializer() *Method {
	clinit := this.GetMethod("<clinit>", "()V")
	if clinit != nil && clinit.isStatic() { // must be static
		return clinit
	}
	return nil
}

/*
<init>(..)V methods
 */
func (this *Class) GetConstructors(publicOnly bool) []*Method {
	var constructors []*Method
	for _, method := range this.methods {
		if method.name == "<init>" && method.returnDescriptor == JVM_SIGNATURE_VOID && !method.isStatic() { // non-static
			if publicOnly {
				if (method.accessFlags & JVM_ACC_PUBLIC) > 0 {
					constructors = append(constructors, method)
				}
			} else {
				constructors = append(constructors, method)
			}
		}
	}

	return constructors
}

func (this *Class) GetConstructor(descriptor string) *Method {
	method := this.GetMethod("<init>", descriptor)
	if method != nil && method.returnDescriptor == JVM_SIGNATURE_VOID && !method.isStatic() {
		return method
	}

	return nil
}

/**
 * Find field with its name and descriptor in the class hierarchy
 * Java doesn't permit a static and instance method with same name + signature
 */
func (this *Class) FindMethod(name string, descriptor string) *Method {
	for _, method := range this.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	if this.superClass != nil {
		method := this.superClass.FindMethod(name, descriptor)
		if method != nil {
			return method
		}
	}
	for _, iface := range this.interfaces {
		method := iface.FindMethod(name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}

func (this *Class) GetDeclaredFields(publicOnly bool) []*Field {
	if !publicOnly {
		return this.fields
	}

	var publicFields []*Field
	for _, field := range this.fields {
		if (field.accessFlags & JVM_ACC_PUBLIC) > 0 {
			publicFields = append(publicFields, field)
		}
	}
	return publicFields
}

func (this *Class) inheritanceDepth() int {
	depth := 1
	for c := this.superClass; c != nil; c = c.superClass {
		depth++
	}
	return depth
}

func (this *Class) GetStaticVariable(name string, descriptor string) Value {
	field := this.FindField(name, descriptor)
	if field == nil || !field.IsStatic() {
		Fatal("Cannot find static variable %s %s in class %s", name, descriptor, this.name)
	}
	return this.staticVars[field.slot]
}

func (this *Class) SetStaticVariable(name string, descriptor string, value Value) {
	field := this.FindField(name, descriptor)
	if field == nil || !field.IsStatic() {
		Fatal("Cannot find static variable %s %s in class %s", name, descriptor, this.name)
	}
	this.staticVars[field.slot] = value
}

type Field struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
	/**
	index of instanceFields or staticFields
	for instance fields, it is the global index considering superclass hierarchy
	 */
	slot int
}

func (this *Field) IsStatic() bool {
	return (this.accessFlags & JVM_ACC_STATIC) > 0
}

func (this *Field) Signature() string {
	return this.name + ":" + this.descriptor
}

func (this *Field) Qualifier() string {
	return this.class.name + "." + this.Signature()
}

func (this *Field) defaultValue() Value {
	var t Value
	switch string(this.descriptor[0]) {
	case JVM_SIGNATURE_BYTE:
		t = Byte(0)
	case JVM_SIGNATURE_SHORT:
		t = Short(0)
	case JVM_SIGNATURE_CHAR:
		t = Char(0)
	case JVM_SIGNATURE_INT:
		t = Int(0)
	case JVM_SIGNATURE_LONG:
		t = Long(0)
	case JVM_SIGNATURE_FLOAT:
		t = Float(0.0)
	case JVM_SIGNATURE_DOUBLE:
		t = Double(0.0)
	case JVM_SIGNATURE_BOOLEAN:
		t = FALSE
	case JVM_SIGNATURE_CLASS:
		t = NULL
	case JVM_SIGNATURE_ARRAY:
		t = NULL
	default:
		Fatal("Not a valid descriptor to get a default value")
	}

	return t
}

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlerPc int
	catchType int // index of constant pool: ClassRef
}

type Method struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class

	maxStack    uint
	maxLocals   uint
	code        []uint8             //u4 code_length
	exceptions  []*ExceptionHandler //u2 exception_table_length
	localVars   []*LocalVariable
	lineNumbers []*LineNumber

	parameterDescriptors []string
	returnDescriptor     string
}

func (this *Method) isStatic() bool {
	return (this.accessFlags & JVM_ACC_STATIC) > 0
}

func (this *Method) isNative() bool {
	return (this.accessFlags & JVM_ACC_NATIVE) > 0
}

func (this *Method) isSynchronized() bool {
	return (this.accessFlags & JVM_ACC_SYNCHRONIZED) > 0
}

func (this *Method) Signature() string {
	return this.name + JVM_SIGNATURE_FUNC + strings.Join(this.parameterDescriptors, "") + JVM_SIGNATURE_ENDFUNC + this.returnDescriptor

}

func (this *Method) Qualifier() string {
	return this.class.name + "." + this.Signature()
}

type LocalVariable struct {
	method     *Method
	startPc    uint16
	length     uint16
	index      uint16
	name       string
	descriptor string
}

type LineNumber struct {
	startPc    int
	lineNumber int
}

/**
 * Runtime constant pool
 */
type Constant interface{}

type SymbolRef interface {
	Constant
	resolve()
}

type ClassRef struct {
	hostClass *Class
	className string
	class     *Class // a class ref may be non-array class, array class or an interface
}

func (this *ClassRef) ResolvedClass() *Class {
	if this.class == nil {
		this.resolve()
	}
	return this.class
}

func (this *ClassRef) resolve() {
	this.class = VM.resolveClass(this.className, this.hostClass, TRIGGER_BY_RESOLVE_CLASS_REF)
	//this.class = VM.CreateClass(this.className, this.hostClass, TRIGGER_BY_RESOLVE_CLASS_REF)
}

type MemberRef struct {
	hostClass *Class

	className  string
	class      *Class
	name       string
	descriptor string
}

func (this *MemberRef) ResolvedClass() *Class {
	if this.class == nil {
		this.resolve()
	}
	return this.class
}

func (this *MemberRef) resolve() {
	this.class = VM.resolveClass(this.className, this.hostClass, TRIGGER_BY_RESOLVE_CLASS_REF)
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
	method := class.FindMethod(this.name, this.descriptor)
	if method == nil {
		Bug("Resolve Method failed")
	}
	this.method = method
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
	class := this.ResolvedClass()
	method := class.FindMethod(this.name, this.descriptor)
	if method == nil {
		Bug("Resolve Interface Method failed")
	}
	this.method = method
}

type StringConstant struct {
	hostClass *Class
	value     string
	jstring   JavaLangString
}

func (this *StringConstant) ResolvedString() ObjectRef {
	if this.jstring.IsNull() {
		this.resolve()
	}
	return this.jstring
}

func (this *StringConstant) resolve() {
	this.jstring = VM.NewJavaLangString(this.value)
}

type UTF8Constant struct {
	hostClass *Class
	value     string
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
	hostClass  *Class
	name       string
	descriptor string
}

type MethodTypeConstant struct {
	hostClass  *Class
	descriptor string
}

type MethodHandleConstant struct {
	hostClass      *Class
	referenceKind  u1
	referenceIndex u2
}

type InvokeDynamicConstant struct {
	hostClass       *Class
	bootstrapMethod string
	name            string
	descriptor      string
}
