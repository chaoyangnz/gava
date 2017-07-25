package jago

import (
	//"reflect"
)

var SYS_CLASS_PATH = "jdk/classes:/Users/Chao/Dropbox/Projects/jago-showcase/build/classes/main"
var EXT_CLASS_PATH = []string{"ext"}
var APP_CLASS_PATH []string
var MAIN_CLASS string

const TRACE_LOG_FILE = "trace.log"
const LOG_LEVEL = INFO


var LOG = NewLog(TRACE_LOG_FILE)

var STRING_TABLE = map[string]JavaLangString{}


var BOOTSTRAP_CLASSLOADER = NewClassLoader(SYS_CLASS_PATH, nil)

var THREAD_MANAGER = &ThreadManager{}






