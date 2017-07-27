package jago

/*178 (0xB2)*/
func GETSTATIC(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	index := f.index16()

	fieldref := c.constantPool[index].(*FieldRef)
	field := fieldref.ResolvedField()

	f.push(field.class.staticVars[field.index])
	f.nextPc()
}

/*179 (0xB3)*/
func PUTSTATIC(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	index := f.index16()
	value := f.pop()

	fieldref := c.constantPool[index].(*FieldRef)
	field := fieldref.ResolvedField()

	field.class.staticVars[field.index] = value
	f.nextPc()
}

/*180 (0xB4)*/
func GETFIELD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	index := f.index16()
	objectref := f.pop().(ObjectRef)

	f.push(f.getField(objectref, index))
	f.nextPc()
}

/*181 (0xB5)*/
func PUTFIELD(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	index := f.index16()
	value := f.pop()
	objectref := f.pop().(ObjectRef)

	f.putField(objectref, index, value)
	f.nextPc()
}

/*182 (0xB6)*/
func INVOKEVIRTUAL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	pc := f.pc
	index := f.index16()
	method := c.constantPool[index].(*MethodRef).ResolvedMethod()

	if method.isStatic() {
		Fatal("Not an instance method")
	}
	params := f.params(method)
	objectref := params[0].(Reference)
	if objectref.IsNull() {
		VM_throw("java/lang/NullPointerException", "Cannot call method %s.%s%s on null", method.class.name, method.name, method.descriptor)
	}

	overridenMethod := objectref.Class().FindMethod(method.name, method.descriptor)

	VM_invokeMethod(t, overridenMethod, params...)
	if pc == f.pc { // otherwise, may be offset due to exception caught
		f.nextPc()
	}
}

// like invokevirtual with objectref, but don't find along the inheritance
/*183 (0xB7)*/
func INVOKESPECIAL(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	pc := f.pc
	index := f.index16()

	method := c.constantPool[index].(*MethodRef).ResolvedMethod()
	params := f.params(method)

	VM_invokeMethod(t, method, params...)
	if pc == f.pc { // otherwise, may be offset due to exception caught
		f.nextPc()
	}
}

/*184 (0xB8)*/
func INVOKESTATIC(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	pc := f.pc
	index := f.index16()

	methodref := c.constantPool[index].(*MethodRef)

	method := methodref.ResolvedMethod()
	params := f.params(method)

	VM_invokeMethod(t, method, params...)
	if pc == f.pc { // otherwise, may be offset due to exception caught
		f.nextPc()
	}
}

/*185 (0xB9)*/
func INVOKEINTERFACE(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	pc := f.pc
	index := f.index16()
	method := c.constantPool[index].(*InterfaceMethodRef).ResolvedMethod()

	if method.isStatic() {
		Fatal("Not an instance method")
	}
	params := f.params(method)
	// get objectref and target method
	objectref := params[0].(Reference)
	if objectref.IsNull() {
		Fatal("NullPointerException")
	}

	overridenMethod := objectref.Class().FindMethod(method.name, method.descriptor)


	VM_invokeMethod(t, overridenMethod, params...)
	if pc == f.pc { // otherwise, may be offset due to exception caught
		f.nextPc()
	}
}

/*186 (0xBA)*/
func INVOKEDYNAMIC(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	Fatal("Not implemented for opcode %d\n", opcode)
}

/*187 (0xBB)*/
func NEW(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	index := f.index16()
	class := c.constantPool[index].(*ClassRef).ResolvedClass()
	objectref := class.NewObject()
	f.push(objectref)
	f.nextPc()
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
func NEWARRAY(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	atype := f.const8(1)
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
	arrayClass := BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + componentDescriptor)
	arrayref := arrayClass.NewArray(count)
	f.push(arrayref)
	f.nextPc()
}

/*189 (0xBD)*/
func ANEWARRAY(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	index := f.index16()
	count := f.pop().(Int)

	var arrayClass *Class
	componentType := c.constantPool[index].(*ClassRef).ResolvedClass()
	if !componentType.IsArray() {
		arrayClass = BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + JVM_SIGNATURE_CLASS + componentType.Name() + JVM_SIGNATURE_ENDCLASS)
	} else {
		arrayClass = BOOTSTRAP_CLASSLOADER.CreateClass(JVM_SIGNATURE_ARRAY + componentType.Name())
	}

	arrayref := arrayClass.NewArray(count)

	f.push(arrayref)
	f.nextPc()
}

/*190 (0xBE)*/
func ARRAYLENGTH(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	f.push(f.pop().(ArrayRef).Length())
	f.nextPc()
}

/*191 (0xBF)*/
func ATHROW(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	throwable := f.pop().(Reference)
	if throwable.IsNull() {
		VM_throw("java/lang/NullPointerException", "cannot throw a null throwable")
	}

	f.nextPc()
	doThrow(throwable)
}

/*192 (0xC0)*/
func CHECKCAST(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	index := f.index16()
	objectref := f.pop().(Reference)

	if objectref.IsNull() {
		f.push(objectref)
		f.nextPc()
		return
	}

	class := c.constantPool[index].(*ClassRef).ResolvedClass()
	if class.IsAssignableFrom(objectref.Class()) {
		f.push(objectref)
		f.nextPc()
		return
	}

	f.nextPc()
	VM_throw("java/lang/ClassCastException", "cannot cast from " + objectref.Class().Name() + " to " + class.Name())
}

/*193 (0xC1)*/
func INSTANCEOF(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	index := f.index16()
	objectref := f.pop().(Reference)

	if objectref.IsNull() {
		f.push(Int(0))
		f.nextPc()
		return
	}

	S := objectref.Class()
	T := c.constantPool[index].(*ClassRef).ResolvedClass()
	// TODO ???
	if(T.IsAssignableFrom(S)) {
		f.push(Int(1))
		f.nextPc()
		return
	}

	f.push(Int(0))
	f.nextPc()
}

/*194 (0xC2)*/
func MONITORENTER(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	/*objectref := */f.pop()

	//LOG.Warn("Not implemented for opcode %d\n", opcode)
	f.nextPc()
}

/*195 (0xC3)*/
func MONITOREXIT(opcode uint8, t *Thread, f *Frame, c *Class, m *Method) {
	/*objectref := */f.pop()

	//LOG.Warn("Not implemented for opcode %d\n", opcode)
	f.nextPc()
}
