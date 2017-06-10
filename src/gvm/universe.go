package gvm

var coreClassPath = []string{"jdk/classes", "/Users/Charvis/Dropbox/Projects/gvm-java/out/production/gvm-java"}
var extClassPath  = []string{"ext"}
var AppClasspath []string

var stringTable = make(map[string]java_lang_string)

var classCache = make(map[string] *ClassType)

var bootstrapClassLoader = &BootstrapClassLoader{}


