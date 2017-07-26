package jago

import (
	//"reflect"
)

var SYS_CLASS_PATH = "jdk/classes:/Users/Chao/Dropbox/Projects/jago-showcase/build/classes/main:/Library/Java/JavaVirtualMachines/jdk1.8.0_152.jdk/Contents/Home/jre/lib/ext/localedata.jar:/Users/Chao/Dropbox/Projects/jago-showcase/build/classes/test"
var EXT_CLASS_PATH = []string{"ext"}
var APP_CLASS_PATH []string
var MAIN_CLASS string

const MISC_LOG_FILE  = "misc.log"
const EXEC_LOG_FILE = "execution.log"
const CLASSLOAD_LOG_FILE = "classload.log"

var LOG = NewLog("misc", DEBUG, MISC_LOG_FILE)
var EXEC_LOG = NewLog("execution", DEBUG, EXEC_LOG_FILE)
var CLASSLOAD_LOG = NewLog("classload", DEBUG, CLASSLOAD_LOG_FILE)

var STRING_TABLE = map[string]JavaLangString{}


var BOOTSTRAP_CLASSLOADER = NewClassLoader(SYS_CLASS_PATH, nil)

var THREAD_MANAGER = &ThreadManager{}






