package javo

/*178 (0xB2)*/
func GETSTATIC(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()

	fieldref := c.constantPool[index].(*FieldRef)
	field := fieldref.ResolvedField()
	VM.initialize(field.class)

	f.push(field.class.staticVars[field.slot])
}

/*179 (0xB3)*/
func PUTSTATIC(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	value := f.pop()

	fieldref := c.constantPool[index].(*FieldRef)
	field := fieldref.ResolvedField()
	VM.initialize(field.class)

	field.class.staticVars[field.slot] = value
}

/*180 (0xB4)*/
func GETFIELD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	objectref := f.pop().(ObjectRef)

	f.push(f.getField(objectref, index))
}

/*181 (0xB5)*/
func PUTFIELD(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	value := f.pop()
	objectref := f.pop().(ObjectRef)

	f.putField(objectref, index, value)
}

/*182 (0xB6)*/
func INVOKEVIRTUAL(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	method := c.constantPool[index].(*MethodRef).ResolvedMethod()

	if method.isStatic() {
		Fatal("Not an instance method")
	}
	params := f.loadParameters(method)
	objectref := params[0].(Reference)
	if objectref.IsNull() {
		VM.Throw("java/lang/NullPointerException", "Cannot call method %s.%s%s on null", method.class.name, method.name, method.descriptor)
	}

	overridenMethod := objectref.Class().FindMethod(method.name, method.descriptor)

	f.push(VM.InvokeMethod(overridenMethod, params...))
}

// like invokevirtual with objectref, but don't find along the inheritance
/*183 (0xB7)*/
func INVOKESPECIAL(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()

	method := c.constantPool[index].(*MethodRef).ResolvedMethod()
	params := f.loadParameters(method)

	f.push(VM.InvokeMethod(method, params...))
}

/*184 (0xB8)*/
func INVOKESTATIC(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()

	methodref := c.constantPool[index].(*MethodRef)

	method := methodref.ResolvedMethod()
	// here, its initialization is invoked in VM.InvokeMethod(..)
	VM.initialize(method.class)
	params := f.loadParameters(method)

	f.push(VM.InvokeMethod(method, params...))
}

/*185 (0xB9)*/
func INVOKEINTERFACE(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	method := c.constantPool[index].(*InterfaceMethodRef).ResolvedMethod()

	if method.isStatic() {
		Fatal("Not an instance method")
	}
	params := f.loadParameters(method)
	// get objectref and target method
	objectref := params[0].(Reference)
	if objectref.IsNull() {
		Fatal("NullPointerException")
	}

	overridenMethod := objectref.Class().FindMethod(method.name, method.descriptor)

	f.push(VM.InvokeMethod(overridenMethod, params...))
}

/*186 (0xBA)*/
func INVOKEDYNAMIC(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}

/*187 (0xBB)*/
func NEW(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	class := c.constantPool[index].(*ClassRef).ResolvedClass()
	VM.initialize(class)
	objectref := VM.NewObject(class)
	f.push(objectref)
}

// array component type
const (
	T_BOOLEAN = 4
	T_CHAR    = 5
	T_FLOAT   = 6
	T_DOUBLE  = 7
	T_BYTE    = 8
	T_SHORT   = 9
	T_INT     = 10
	T_LONG    = 11
)

/*188 (0xBC)*/
func NEWARRAY(t *Thread, f *Frame, c *Class, m *Method) {
	atype := f.operandConst8()
	var componentType Type
	switch atype {
	case T_CHAR:
		componentType = CHAR_TYPE
	case T_BYTE:
		componentType = BYTE_TYPE
	case T_SHORT:
		componentType = SHORT_TYPE
	case T_INT:
		componentType = INT_TYPE
	case T_LONG:
		componentType = LONG_TYPE
	case T_FLOAT:
		componentType = FLOAT_TYPE
	case T_DOUBLE:
		componentType = DOUBLE_TYPE
	case T_BOOLEAN:
		componentType = BOOLEAN_TYPE
	default:
		Fatal("Invalid atype value")
	}
	count := f.pop().(Int)
	//arrayClass := VM.CreateClass(JVM_SIGNATURE_ARRAY + componentDescriptor, c, TRIGGER_BY_NEW_INSTANCE)
	//VM.initialize(arrayClass) // not mentioned in jvms
	//arrayref := VM.NewArray(arrayClass, count)
	arrayref := VM.NewArrayOfComponent(componentType, count)
	f.push(arrayref)
}

/*189 (0xBD)*/
func ANEWARRAY(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	count := f.pop().(Int)

	//var arrayClass *Class
	componentType := c.constantPool[index].(*ClassRef).ResolvedClass()
	//if !componentType.IsArray() {
	//	arrayClass = VM.CreateClass(JVM_SIGNATURE_ARRAY + JVM_SIGNATURE_CLASS + componentType.Name() + JVM_SIGNATURE_ENDCLASS, c, TRIGGER_BY_NEW_INSTANCE)
	//} else {
	//	arrayClass = VM.CreateClass(JVM_SIGNATURE_ARRAY + componentType.Name(), c, TRIGGER_BY_NEW_INSTANCE)
	//}
	//
	////VM.initialize(arrayClass) // not mentioned in jvms
	//
	//arrayref := VM.NewArray(arrayClass, count)
	arrayref := VM.NewArrayOfComponent(componentType, count)

	f.push(arrayref)
}

/*190 (0xBE)*/
func ARRAYLENGTH(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.pop().(ArrayRef).ArrayLength())
}

/*191 (0xBF)*/
func ATHROW(t *Thread, f *Frame, c *Class, m *Method) {
	throwable := f.pop().(Reference)
	if throwable.IsNull() {
		VM.Throw("java/lang/NullPointerException", "cannot throw a null throwable")
	}

	VM.Throw0(throwable, THROWN_BY_ATHROW)
}

/*192 (0xC0)*/
func CHECKCAST(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	objectref := f.pop().(Reference)

	if objectref.IsNull() {
		f.push(objectref)
		return
	}

	class := c.constantPool[index].(*ClassRef).ResolvedClass()
	if class.IsAssignableFrom(objectref.Class()) {
		f.push(objectref)
		return
	}

	VM.Throw("java/lang/ClassCastException", "cannot cast from "+objectref.Class().Name()+" to "+class.Name())
}

/*193 (0xC1)*/
func INSTANCEOF(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	objectref := f.pop().(Reference)

	if objectref.IsNull() {
		f.push(Int(0))
		return
	}

	S := objectref.Class()
	T := c.constantPool[index].(*ClassRef).ResolvedClass()
	// TODO ???
	if T.IsAssignableFrom(S) {
		f.push(Int(1))
		return
	}

	f.push(Int(0))
}

/*194 (0xC2)*/
func MONITORENTER(t *Thread, f *Frame, c *Class, m *Method) {
	objectref := f.pop().(Reference)
	if objectref.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}

	objectref.Monitor().Enter()
}

/*195 (0xC3)*/
func MONITOREXIT(t *Thread, f *Frame, c *Class, m *Method) {
	objectref := f.pop().(Reference)

	if objectref.IsNull() {
		VM.Throw("java/lang/NullPointerException", "")
	}

	objectref.Monitor().Exit()
}
