package jago

import (
	//"reflect"
)

var SYS_CLASS_PATH = "jdk/classes:example/classes"
var EXT_CLASS_PATH = []string{"ext"}
var APP_CLASS_PATH []string
var MAIN_CLASS string

const MISC_LOG_FILE  = "log/misc.log"
const EXEC_LOG_FILE = "log/execution.log"
const CLASSLOAD_LOG_FILE = "log/classload.log"

var LOG = NewLog("misc", INFO, MISC_LOG_FILE)
var EXEC_LOG = NewLog("execution", INFO, EXEC_LOG_FILE)
var CLASSLOAD_LOG = NewLog("classload", INFO, CLASSLOAD_LOG_FILE)

var STRING_TABLE = map[string]JavaLangString{}


var BOOTSTRAP_CLASSLOADER = NewClassLoader(SYS_CLASS_PATH, nil)

var THREAD_MANAGER = &ThreadManager{}






