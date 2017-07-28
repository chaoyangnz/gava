package jago

import (
	"time"
	"reflect"
	"fmt"
	"os"
)

func _isClinit(method *Method) bool {
	return method.name == "<clinit>" && method.descriptor == "()V"
}

func VM_invokeMethod(className string, methodName string, methodDescriptor string, params ... Value)  {
	class := BOOTSTRAP_CLASSLOADER.CreateClass(className, TRIGGER_BY_ACCESS_MEMBER)
	method := class.GetMethod(methodName, methodDescriptor)
	VM_invokeMethod0(method, params...)
}

func VM_invokeMethod0(method *Method, params ... Value)  {
	thread := VM_getCurrentThread()
	if !method.isNative() {
		current := thread.current()
		frame := NewStackFrame(method)
		i := 0
		for _, param := range params {
			frame.storeVar(uint(i), param)

			switch param.(type) {
			case Long, Double:
				i += 2
			default:
				i += 1
			}
		}
		thread.pushFrames(frame)
		thread.RunTo(current)
	} else {
		if !method.isNative() {
			Fatal("%s Not a native method", method.Qualifier())
		}
		EXEC_LOG.Info("\n%s\t🍎%s", repeat("\t", thread.indexOf(thread.current())), method.Qualifier())

		fun, found := findNative(method.Qualifier())

		staticDesc := ""
		if method.isStatic() {
			staticDesc = "static"
		}
		if !found {
			Fatal( "native method %s %s is not implemented.", staticDesc, method.Qualifier())
		}

		if len(params) != fun.Type().NumIn() {
			Fatal( "The number of loadParameters is not adapted for native method %s %s.", staticDesc, method.Qualifier())
		}
		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
			if in[k].Kind() == reflect.Invalid {
				Bug("Native method loadParameters is nil, bug!")
			}
		}
		result := fun.Call(in)


		if len(result) == 0 {
			return
		}

		if method.returnDescriptor != JVM_SIGNATURE_VOID {
			if ret, ok := result[0].Interface().(Value); ok {
				thread.current().push(ret)
			} else {
				Bug("native return wrong type")
			}

		}
	}
}

func VM_getCurrentThread() *Thread {
	return THREAD_MANAGER.current()
}

func VM_getClass(name string) *Class {
	return BOOTSTRAP_CLASSLOADER.CreateClass(name, TRIGGER_BY_JAVA_REFLECTION)
}

func VM_getInstanceVariable(objref ObjectRef, name string, descriptor string) Value {
	return objref.GetInstanceVariableByName(name, descriptor)
}

func VM_setInstanceVariable(objref ObjectRef, name string, descriptor string, value Value)  {
	objref.SetInstanceVariableByName(name, descriptor, value)
}

func VM_getStaticVariable(class *Class, name string, descriptor string) Value {
	field := class.FindField(name, descriptor)
	if field == nil || !field.IsStatic() {
		Fatal("Cannot find static variable %s %s in class %s", name, descriptor, class.name)
	}
	return class.staticVars[field.index]
}

func VM_setStaticVariable(class *Class, name string, descriptor string, value Value)  {
	field := class.FindField(name, descriptor)
	if field == nil || !field.IsStatic() {
		Fatal("Cannot find static variable %s %s in class %s", name, descriptor, class.name)
	}
	class.staticVars[field.index] = value
}

func VM_currentTimeMillis() Long {
	return Long(time.Now().UnixNano() / int64(time.Millisecond))
}

func VM_currentTimeNano() Long {
	return Long(time.Now().UnixNano())
}

func VM_iHashCode(ref Reference) Int {
	if ref.IsNull() {
		return Int(0)
	}
	return ref.oop.header.hashCode
}

func VM_intern_String(stringobj JavaLangString) JavaLangString {
	str := stringobj.toNativeString()

	if stringObject, ok := STRING_TABLE[str]; ok {
		return stringObject
	} else {
		STRING_TABLE[str] = stringobj
		return stringobj
	}
}

/*
The whole project should use panic only here !!!!!
 */
func VM_Throw0(throwable Reference)  {
	panic(throwable)
}

func VM_throw(exception string, message string, args ...interface{})  {
	VM_Throw0(NewThrowable(exception, message, args...))
}

func VM_stdoutPrintf(format string, args ...interface{})  {
	fmt.Fprintf(os.Stdout, format, args...)
}

func VM_stderrPrintf(format string, args ...interface{})  {
	fmt.Fprintf(os.Stderr, format, args...)
}

func VM_getTypeClass(descriptor string) JavaLangClass {
	var typeClass JavaLangClass
	switch string(descriptor[0]) {
	case JVM_SIGNATURE_CLASS: {
		fieldTypeClassName := descriptor[1:len(descriptor)-1]
		typeClass = BOOTSTRAP_CLASSLOADER.CreateClass(fieldTypeClassName, TRIGGER_BY_JAVA_REFLECTION).ClassObject()
	}
	case JVM_SIGNATURE_ARRAY: {
		fieldTypeClassName := descriptor
		typeClass = BOOTSTRAP_CLASSLOADER.CreateClass(fieldTypeClassName, TRIGGER_BY_JAVA_REFLECTION).ClassObject()
	}
	case JVM_SIGNATURE_BYTE: typeClass = BYTE_TYPE.ClassObject()
	case JVM_SIGNATURE_SHORT: typeClass = SHORT_TYPE.ClassObject()
	case JVM_SIGNATURE_CHAR: typeClass = CHAR_TYPE.ClassObject()
	case JVM_SIGNATURE_INT: typeClass = INT_TYPE.ClassObject()
	case JVM_SIGNATURE_LONG: typeClass = LONG_TYPE.ClassObject()
	case JVM_SIGNATURE_FLOAT: typeClass = FLOAT_TYPE.ClassObject()
	case JVM_SIGNATURE_DOUBLE: typeClass = DOUBLE_TYPE.ClassObject()
	case JVM_SIGNATURE_BOOLEAN: typeClass = BOOLEAN_TYPE.ClassObject()
	default:
		Fatal("type %s is not a unsupported type", descriptor)
	}

	return typeClass
}


