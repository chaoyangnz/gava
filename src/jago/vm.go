package jago

import (
	"time"
	"reflect"
	"fmt"
	"os"
)

func (this *Core) InvokeMethodOf(className string, methodName string, methodDescriptor string, params ... Value)  {
	class := VM.CreateClass(className, TRIGGER_BY_ACCESS_MEMBER)
	method := class.GetMethod(methodName, methodDescriptor)
	this.InvokeMethod(method, params...)
}

/*
This method is used to run a method and return value (even void method return a void value)
 */
func (this *Core) InvokeMethod(method *Method, params ... Value) Value {
	thread := this.CurrentThread()
	if method.isSynchronized() {
		var monitor *Monitor
		if method.isStatic() {
			monitor = method.class.classObject.(Reference).Monitor()
		} else {
			monitor = params[0].(Reference).Monitor()
		}
		monitor.Enter()
		defer monitor.Exit()
	}
	if !method.isNative() {
		frame := thread.NewStackFrame(method)
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
		caller := thread.current()
		thread.push(frame)
		thread.ExecuteFrame()


		if frame.exception.IsNull() { // normal return
			if method.returnDescriptor != JVM_SIGNATURE_VOID {
				return caller.pop() // if non-normal return, here will return a throwable
			}
		} else {
			// must be exception caught method
			this.Throw0(frame.exception, THROWN_BY_RETHROW) // rethrow
			return VOID
		}

	} else {
		if !method.isNative() {
			Fatal("%s Not a native method", method.Qualifier())
		}
		EXEC_LOG.Debug("\n%s\t🔸[%s]%s", repeat("\t", thread.indexOf(thread.current())), thread.name, method.Qualifier())

		fun, found := VM.FindNative(method.Qualifier())

		staticDesc := ""
		if method.isStatic() {
			staticDesc = "static"
		}
		if !found {
			Fatal( "native method %s %s is not implemented.", staticDesc, method.Qualifier())
		}

		if len(params) != fun.Type().NumIn() {
			Fatal( "The size of parameters is not adapted for native method %s %s.", staticDesc, method.Qualifier())
		}
		in := make([]reflect.Value, len(params))
		for k, param := range params {
			in[k] = reflect.ValueOf(param)
			if in[k].Kind() == reflect.Invalid {
				Bug("Native method loadParameters is nil, bug!")
			}
		}
		result := fun.Call(in)

		if  method.returnDescriptor != JVM_SIGNATURE_VOID {
			if len(result) != 0 {
				if ret, ok := result[0].Interface().(Value); ok {
					return ret
				}
			} else {
				Bug("native return wrong type or size")
			}
		}
	}

	return VOID
}

func (this *Core) CurrentThread() *Thread {
	return VM.current()
}

func (this *Core) GetClass(name string) *Class {
	return VM.CreateClass(name, TRIGGER_BY_JAVA_REFLECTION)
}

func (this *Core) GetInstanceVariable(objref ObjectRef, name string, descriptor string) Value {
	return objref.GetInstanceVariableByName(name, descriptor)
}

func (this *Core) SetInstanceVariable(objref ObjectRef, name string, descriptor string, value Value)  {
	objref.SetInstanceVariableByName(name, descriptor, value)
}

func (this *Core) GetStaticVariable(class *Class, name string, descriptor string) Value {
	field := class.FindField(name, descriptor)
	if field == nil || !field.IsStatic() {
		Fatal("Cannot find static variable %s %s in class %s", name, descriptor, class.name)
	}
	return class.staticVars[field.slot]
}

func (this *Core) SetStaticVariable(class *Class, name string, descriptor string, value Value)  {
	field := class.FindField(name, descriptor)
	if field == nil || !field.IsStatic() {
		Fatal("Cannot find static variable %s %s in class %s", name, descriptor, class.name)
	}
	class.staticVars[field.slot] = value
}

func (this *Core) CurrentTimeMillis() Long {
	return Long(time.Now().UnixNano() / int64(time.Millisecond))
}

func (this *Core) CurrentTimeNano() Long {
	return Long(time.Now().UnixNano())
}

func (this *Core) IHashCode(ref Reference) Int {
	if ref.IsNull() {
		return Int(0)
	}
	return ref.oop.header.hashCode
}

func (this *Core) Clone(obj Reference) Reference {

	var clone Reference
	if !obj.Class().IsArray() {
		clone = VM.NewObject(obj.Class()).(Reference)
	} else {
		clone = VM.NewArray(obj.Class(), obj.ArrayLength()).(Reference)
	}

	for i,value := range obj.oop.slots {
		clone.oop.slots[i] = value
	}

	return clone
}

const (
	THROWN_BY_ATHROW         = "thrown by \"athrow\"" // ATHROW instruction
	THROWN_BY_RETHROW       = "rethrown"
	THROWN_BY_VM            = "thrown by VM"
)
/*
The whole project should use panic only here !!!!!
 */
func (this *Core) Throw0(throwable Reference, thrownReason string)  {
	thread := this.CurrentThread()
	if thread.current() != nil {
		EXEC_LOG.Info("\n%s🔥Exception %s: %s at %s", repeat("\t", thread.indexOf(thread.current())+1), thrownReason, throwable.Class().name, thread.current().method.Qualifier())
	}
	panic(throwable)
}

func (this *Core) Throw(exception string, message string, args ...interface{})  {
	this.Throw0(VM.NewThrowable(exception, message, args...), THROWN_BY_VM)
}

func (this *Core) StdoutPrintf(format string, args ...interface{})  {
	fmt.Fprintf(os.Stdout, format, args...)
}

func (this *Core) StderrPrintf(format string, args ...interface{})  {
	fmt.Fprintf(os.Stderr, format, args...)
}

func (this *Core) GetTypeClass(descriptor string) JavaLangClass {
	var typeClass JavaLangClass
	switch string(descriptor[0]) {
	case JVM_SIGNATURE_CLASS: {
		fieldTypeClassName := descriptor[1:len(descriptor)-1]
		typeClass = VM.CreateClass(fieldTypeClassName, TRIGGER_BY_JAVA_REFLECTION).ClassObject()
	}
	case JVM_SIGNATURE_ARRAY: {
		fieldTypeClassName := descriptor
		typeClass = VM.CreateClass(fieldTypeClassName, TRIGGER_BY_JAVA_REFLECTION).ClassObject()
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




