package jago

/*178 (0xB2)*/
func GETSTATIC(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()

	fieldref := c.constantPool[index].(*FieldRef)
	field := fieldref.ResolvedField()

	f.push(field.class.staticVars[field.index])
}

/*179 (0xB3)*/
func PUTSTATIC(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	value := f.pop()

	fieldref := c.constantPool[index].(*FieldRef)
	field := fieldref.ResolvedField()

	field.class.staticVars[field.index] = value
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
		VM_throw("java/lang/NullPointerException", "Cannot call method %s.%s%s on null", method.class.name, method.name, method.descriptor)
	}

	overridenMethod := objectref.Class().FindMethod(method.name, method.descriptor)

	f.push(VM_invokeMethod(overridenMethod, params...))
}

// like invokevirtual with objectref, but don't find along the inheritance
/*183 (0xB7)*/
func INVOKESPECIAL(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()

	method := c.constantPool[index].(*MethodRef).ResolvedMethod()
	params := f.loadParameters(method)

	f.push(VM_invokeMethod(method, params...))
}

/*184 (0xB8)*/
func INVOKESTATIC(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()

	methodref := c.constantPool[index].(*MethodRef)

	method := methodref.ResolvedMethod()
	params := f.loadParameters(method)

	f.push(VM_invokeMethod(method, params...))
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

	f.push(VM_invokeMethod(overridenMethod, params...))
}

/*186 (0xBA)*/
func INVOKEDYNAMIC(t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", f.opcode())
}

/*187 (0xBB)*/
func NEW(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	class := c.constantPool[index].(*ClassRef).ResolvedClass()
	objectref := class.NewObject()
	f.push(objectref)
}

// array component type
const (
	T_BOOLEAN	    = 4
	T_CHAR	        = 5
	T_FLOAT	        = 6
	T_DOUBLE	    = 7
	T_BYTE	        = 8
	T_SHORT	        = 9
	T_INT	        = 10
	T_LONG	        = 11
)

/*188 (0xBC)*/
func NEWARRAY(t *Thread, f *Frame, c *Class, m *Method) {
	atype := f.operandConst8()
	var componentDescriptor string
	switch atype {
	case T_CHAR: componentDescriptor = JVM_SIGNATURE_CHAR
	case T_BYTE: componentDescriptor = JVM_SIGNATURE_BYTE
	case T_SHORT: componentDescriptor = JVM_SIGNATURE_SHORT
	case T_INT: componentDescriptor = JVM_SIGNATURE_INT
	case T_LONG: componentDescriptor = JVM_SIGNATURE_LONG
	case T_FLOAT: componentDescriptor = JVM_SIGNATURE_FLOAT
	case T_DOUBLE: componentDescriptor = JVM_SIGNATURE_DOUBLE
	case T_BOOLEAN: componentDescriptor = JVM_SIGNATURE_BOOLEAN
	default:
		Fatal("Invalid atype value")
	}
	count := f.pop().(Int)
	arrayClass := BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + componentDescriptor, TRIGGER_BY_NEW_INSTANCE)
	arrayref := arrayClass.NewArray(count)
	f.push(arrayref)
}

/*189 (0xBD)*/
func ANEWARRAY(t *Thread, f *Frame, c *Class, m *Method) {
	index := f.operandIndex16()
	count := f.pop().(Int)

	var arrayClass *Class
	componentType := c.constantPool[index].(*ClassRef).ResolvedClass()
	if !componentType.IsArray() {
		arrayClass = BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + JVM_SIGNATURE_CLASS + componentType.Name() + JVM_SIGNATURE_ENDCLASS, TRIGGER_BY_NEW_INSTANCE)
	} else {
		arrayClass = BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + componentType.Name(), TRIGGER_BY_NEW_INSTANCE)
	}

	arrayref := arrayClass.NewArray(count)

	f.push(arrayref)
}

/*190 (0xBE)*/
func ARRAYLENGTH(t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.pop().(ArrayRef).Length())
}

/*191 (0xBF)*/
func ATHROW(t *Thread, f *Frame, c *Class, m *Method) {
	throwable := f.pop().(Reference)
	if throwable.IsNull() {
		VM_throw("java/lang/NullPointerException", "cannot throw a null throwable")
	}

	VM_Throw0(throwable, THROWN_BY_ATHROW)
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

	VM_throw("java/lang/ClassCastException", "cannot cast from " + objectref.Class().Name() + " to " + class.Name())
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
	if(T.IsAssignableFrom(S)) {
		f.push(Int(1))
		return
	}

	f.push(Int(0))
}

/*194 (0xC2)*/
func MONITORENTER(t *Thread, f *Frame, c *Class, m *Method) {
	objectref := f.pop().(Reference)
	if objectref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}

	objectref.Monitor().Enter(VM_currentThread())
}

/*195 (0xC3)*/
func MONITOREXIT(t *Thread, f *Frame, c *Class, m *Method) {
	objectref := f.pop().(Reference)

	if objectref.IsNull() {
		VM_throw("java/lang/NullPointerException", "")
	}

	objectref.Monitor().Exit(VM_currentThread())
}
