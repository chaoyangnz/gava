package jago

import (
	"reflect"
)

var SYS_CLASS_PATH = "jdk/classes:/Users/Chao/Dropbox/Projects/gvm-java/out/production/gvm-java"
var EXT_CLASS_PATH = []string{"ext"}
var APP_CLASS_PATH []string



var STRING_TABLE = map[string]jobject {}


var BOOTSTRAP_CLASSLOADER = NewClassLoader(SYS_CLASS_PATH, nil)

var THREAD_MANAGER = &ThreadManager{}

var NATIVE_FUNCTIONS = map[string]reflect.Value {
	"Java_GVM_println": reflect.ValueOf(Java_GVM_println),
	"Java_GVM_toUpper": reflect.ValueOf(Java_GVM_toUpper),
	"Java_java_lang_Object_registerNatives": reflect.ValueOf(Java_java_lang_Object_registerNatives),
	"Java_java_lang_System_registerNatives": reflect.ValueOf(Java_java_lang_System_registerNatives),
	"Java_java_lang_Class_registerNatives": reflect.ValueOf(Java_java_lang_System_registerNatives),
}


