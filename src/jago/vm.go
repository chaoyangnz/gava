package jago

import (
	"time"
	"reflect"
	"fmt"
	"os"
)

func VM_invokeMethod0(className string, methodName string, methodDescriptor string, params ... Value)  {
	class := BOOTSTRAP_CLASSLOADER.CreateClass(className, TRIGGER_BY_ACCESS_MEMBER)
	method := class.GetMethod(methodName, methodDescriptor)
	VM_invokeMethod(method, params...)
}

//func VM_invokeMethodWithReturn(method *Method, params ... Value) Value {
//	if method.returnDescriptor == JVM_SIGNATURE_VOID {
//		Bug("Void method don't need to return, use VM_invokeMethod() instead")
//	}
//	VM_invokeMethod(method, params...)
//	thread := VM_currentThread()
//	return thread.current().pop()
//}

/*
This method is used to run a method and return value (even void method return a void value)
 */
func VM_invokeMethod(method *Method, params ... Value) Value {
	thread := VM_currentThread()
	if !method.isNative() {
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
		caller := thread.current()
		thread.push(frame)
		thread.ExecuteFrame()


		if frame.exception.IsNull() { // normal return
			if method.returnDescriptor != JVM_SIGNATURE_VOID {
				return caller.pop() // if non-normal return, here will return a throwable
			}
		} else {
			// must be exception caught method
			VM_Throw0(frame.exception, THROWN_BY_RETHROW) // rethrow
			return VOID
		}

	} else {
		if !method.isNative() {
			Fatal("%s Not a native method", method.Qualifier())
		}
		EXEC_LOG.Debug("\n%s\t🔸[%s]%s", repeat("\t", thread.indexOf(thread.current())), thread.name, method.Qualifier())

		fun, found := findNative(method.Qualifier())

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

func VM_currentThread() *Thread {
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

func VM_clone(obj Reference) Reference {

	var clone Reference
	if !obj.Class().IsArray() {
		clone = obj.Class().NewObject().(Reference)
	} else {
		clone = obj.Class().NewArray(obj.Length()).(Reference)
	}

	for i,value := range obj.oop.values {
		clone.oop.values[i] = value
	}

	return clone
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

const (
	THROWN_BY_ATHROW         = "thrown by \"athrow\"" // ATHROW instruction
	THROWN_BY_RETHROW       = "rethrown"
	THROWN_BY_VM            = "thrown by VM"
)
/*
The whole project should use panic only here !!!!!
 */
func VM_Throw0(throwable Reference, thrownReason string)  {
	thread := VM_currentThread()
	if thread.current() != nil {
		EXEC_LOG.Info("\n%s🔥Exception %s: %s at %s", repeat("\t", thread.indexOf(thread.current())+1), thrownReason, throwable.Class().name, thread.current().method.Qualifier())
	}
	panic(throwable)
}

func VM_throw(exception string, message string, args ...interface{})  {
	VM_Throw0(NewThrowable(exception, message, args...), THROWN_BY_VM)
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

var system_properties = map[string]string {
	"java.version":                "1.8.0_152-ea",
	"java.home":                   "TODO",
	"java.specification.name":     "Java Platform API Specification",
	"java.specification.version":  "1.8",
	"java.specification.vendor":   "Oracle Corporation",

	"java.vendor":                 "Oracle Corporation",
	"java.vendor.url":             "http://java.oracle.com/",
	"java.vendor.url.bug":         "http://bugreport.sun.com/bugreport/",

	"java.vm.name":                "Jago 64-Bit VM",
	"java.vm.version":             "1.0.0",
	"java.vm.vendor":              "Chao Yang",//"Oracle Corporation",
	"java.vm.info":                "mixed mode",
	"java.vm.specification.name":  "Java Virtual Machine Specification",
	"java.vm.specification.version": "1.8",
	"java.vm.specification.vendor": "Oracle Corporation",

	"java.runtime.name":           "Java(TM) SE Runtime Environment",
	"java.runtime.version":        "1.8.0_152-ea-b05",

	"java.class.version":          "52.0",
	"java.class.path":             "TODO",

	"java.io.tmpdir":              "TODO",
	"java.library.path":           "TODO",
	"java.ext.dirs":               "TODO",
	"java.endorsed.dirs":          "TODO",
	"java.awt.graphicsenv":        "sun.awt.CGraphicsEnvironment",
	"java.awt.printerjob":         "sun.lwawt.macosx.CPrinterJob",
	"awt.toolkit":                 "sun.lwawt.macosx.LWCToolkit",

	"path.separator":              ":",
	"line.separator":              "\n",
	"file.separator":              "/",
	"file.encoding":               "UTF-8",
	"file.encoding.pkg":           "sun.io",

	"sun.stdout.encoding":  "UTF-8",
	"sun.stderr.encoding":  "UTF-8",

	"os.name":                     "Mac OS X",
	"os.arch":                     "x86_64",
	"os.version":                  "10.12.5",

	"user.name":                   "TODO",
	"user.home":                   "TODO",
	"user.country":                "TODO",
	"user.language":               "TODO",
	"user.timezone":               "TODO",
	"user.dir":                    "TODO",

	"sun.java.launcher":           "SUN_STANDARD",
	"sun.java.command":            "TODO",
	"sun.boot.library.path":       "TODO",
	"sun.boot.class.path":         "TODO",
	"sun.os.patch.level":          "unknown",
	"sun.jnu.encoding":            "UTF-8",
	"sun.management.compiler":     "HotSpot 64-Bit Tiered Compilers",
	"sun.arch.data.model":         "64",
	"sun.cpu.endian":              "little",
	"sun.io.unicode.encoding":     "UnicodeBig",
	"sun.cpu.isalist":             "",

	"http.nonProxyHosts":          "local|*.local|169.254/16|*.169.254/16",
	"ftp.nonProxyHosts":           "local|*.local|169.254/16|*.169.254/16",
	"socksNonProxyHosts":          "local|*.local|169.254/16|*.169.254/16",
	"gopherProxySet":              "false",
}

func VM_systemProperties() map[string]string {
	return system_properties
}

func VM_setSystemProperty(prop string, value string)  {
	system_properties[prop] = value
}


