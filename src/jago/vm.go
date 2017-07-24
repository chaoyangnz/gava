package jago

import (
	"time"
	"reflect"
)

func _isClinit(method *Method) bool {
	return method.name == "<clinit>" && method.descriptor == "()V"
}

func VM_invokeMethod(thread *Thread, method *Method, params ... Value)  {
	if !method.isNative() {
		current := thread.peekFrame()
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
		LOG.Info("\n%s\t🍎%s", __indent(thread, thread.peekFrame()), method.Qualifier())

		fun, found := findNative(method.Qualifier())

		staticDesc := ""
		if method.isStatic() {
			staticDesc = "static"
		}
		if !found {
			Fatal( "native method %s %s is not implemented.", staticDesc, method.Qualifier())
		}

		if len(params) != fun.Type().NumIn() {
			Fatal( "The number of params is not adapted for native method %s %s.", staticDesc, method.Qualifier())
		}
		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
			if in[k].Kind() == reflect.Invalid {
				Bug("Native method params is nil, bug!")
			}
		}
		result := fun.Call(in)


		if len(result) == 0 {
			return
		}

		if method.returnDescriptor != JVM_SIGNATURE_VOID {
			thread.peekFrame().push(result[0].Interface().(Value))
		}
	}
}

func VM_getCurrentThread() *Thread {
	//TODO
	return THREAD_MANAGER.currentThread
}

func VM_getClass(name string) *Class {
	return BOOTSTRAP_CLASSLOADER.CreateClass(name)
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

//func VM_getCallerClass() JavaLangClass {
//
//}
