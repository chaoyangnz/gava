package gvm

import (
	"reflect"
)

type (
	// these types are for vm internal use only
	// basicaly they are mapped to 10 types mentioned in JVM specification
	t_any       interface{}
		//isPrimitive()   bool
		//isReference()   bool
		//defaultValue()  t_any


	t_byte      int8
	t_char      uint16
	t_short     int16
	t_int       int32
	t_long      int64
	t_boolean   int32
	t_float     float32
	t_double    float64
	t_object    *JavaObject
	t_array     *JavaArray
)

const (
	java_true           = t_boolean(0x1)
	java_false          = t_boolean(0x0)
)

func checkType(value t_any) t_any {
	t := ""
	switch value.(type) {
	case t_byte:    t = "B" //byte
	case t_char:    t = "C" //char
	case t_double:  t = "D" //double
	case t_float:   t= "F" //float
	case t_int:     t= "I" //int
	case t_long:    t="J" //long
	case t_short:   t= "S" //short
	case t_boolean: t = "Z" //boolean
	case t_object:  t = "L" //reference
	case t_array:   t = "[" //array
	default:
		if value == nil {
			t = "nil"
		} else {
			t = reflect.TypeOf(value).String()
			if t[:8] != "gvm.java" {
				fatal("Not a valid vm type")
			}
		}
	}

	all("Check vm type: %s\n", t)
	return value
}

func defaultValue(descriptor string) t_any {
	ch := descriptor[:1]
	var value t_any
	switch ch {
	case "B": value = t_byte(0) //byte
	case "C": value = t_char(0) //char
	case "D": value = t_double(0.0) //double
	case "F": value = t_float(0.0) //float
	case "I": value = t_int(0) //int
	case "J": value = t_long(0) //long
	case "S": value = t_short(0) //short
	case "Z": value = t_boolean(java_false) //boolean
	case "L": value = t_object(nil) //reference
	case "[": value = t_array(nil) //array
	default:
		fatal("Not a valid vm type")
	}
	return value
}



//func (this t_byte) isPrimitive() bool  {
//	return true
//}
//
//func (this t_byte) isReference() bool  {
//	return false
//}
//
//func (this t_byte) defaultValue() t_any {
//	return t_byte(0)
//}
//
//func (this t_char) isPrimitive() bool  {
//	return true
//}
//
//func (this t_char) isReference() bool  {
//	return false
//}
//
//func (this t_char) defaultValue() t_any {
//	return t_char(0)
//}
//
//func (this t_short) isPrimitive() bool  {
//	return true
//}
//
//func (this t_short) isReference() bool  {
//	return false
//}
//
//func (this t_short) defaultValue() t_any {
//	return t_short(0)
//}
//
//func (this t_int) isPrimitive() bool  {
//	return true
//}
//
//func (this t_int) isReference() bool  {
//	return false
//}
//
//func (this t_int) defaultValue() t_any {
//	return t_int(0)
//}
//
//func (this t_long) isPrimitive() bool  {
//	return true
//}
//
//func (this t_long) isReference() bool  {
//	return false
//}
//
//func (this t_long) defaultValue() t_any {
//	return t_long(0)
//}
//
//func (this t_boolean) isPrimitive() bool  {
//	return true
//}
//
//func (this t_boolean) isReference() bool  {
//	return false
//}
//
//func (this t_boolean) defaultValue() t_any {
//	return java_false
//}
//
//func (this t_float) isPrimitive() bool  {
//	return true
//}
//
//func (this t_float) isReference() bool  {
//	return false
//}
//
//func (this t_float) defaultValue() t_any {
//	return t_float(0.0)
//}
//
//func (this t_double) isPrimitive() bool  {
//	return true
//}
//
//func (this t_double) isReference() bool  {
//	return false
//}
//
//func (this t_double) defaultValue() t_any {
//	return t_double(0.0)
//}
//
//func (this *JavaObject) isPrimitive() bool  {
//	return false
//}
//
//func (this *JavaObject) isReference() bool  {
//	return true
//}
//
//func (this  *JavaObject) defaultValue() t_any {
//	return (*JavaObject)(nil)
//}
//
//func (this *JavaArray) isPrimitive() bool  {
//	return false
//}
//
//func (this *JavaArray) isReference() bool  {
//	return true
//}
//
//func (this  *JavaArray) defaultValue() t_any {
//	return (*JavaArray)(nil)
//}