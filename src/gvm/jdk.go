package gvm

import (
	"errors"
	"reflect"
)


func register(name string, fn interface{}) (err interface{})  {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(name + " is not callable.")
		}
	}()
	v := reflect.ValueOf(fn)
	v.Type().NumIn()
	NativeFunctions[name] = v
	return
}

var NativeFunctions = map[string]reflect.Value {
	"Java_GVM_print": reflect.ValueOf(Java_GVM_print),
}