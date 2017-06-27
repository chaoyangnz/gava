package jago

import (
	//"reflect"
)

var SYS_CLASS_PATH = "jdk/classes:/Users/Chao/Dropbox/Projects/gvm-java/out/production/gvm-java"
var EXT_CLASS_PATH = []string{"ext"}
var APP_CLASS_PATH []string



var STRING_TABLE = map[string]JavaLangString{}


var BOOTSTRAP_CLASSLOADER = NewClassLoader(SYS_CLASS_PATH, nil)

var THREAD_MANAGER = &ThreadManager{}




