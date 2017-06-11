package gvm

var coreClassPath = []string{"jdk/classes", "/Users/Charvis/Dropbox/Projects/gvm-java/out/production/gvm-java"}
var extClassPath  = []string{"ext"}
var AppClasspath []string

var stringTable = map[string]java_lang_string {}

var typeCache = map[string]Type {
	"B": BYTE_TYPE,
	"C": CHAR_TYPE,
	"S": SHORT_TYPE,
	"I": INT_TYPE,
	"J": LONG_TYPE,
	"F": FLOAT_TYPE,
	"D": DOUBLE_TYPE,
	"Z": BOOLEAN_TYPE,
	"V": VOID_TYPE,
}

var bootstrapClassLoader = &BootstrapClassLoader{}


