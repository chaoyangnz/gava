package jago

import (
	"github.com/orcaman/concurrent-map"
	"path/filepath"
	"os"
	"os/user"
	"sync"
	"reflect"
	"strings"
)

var SYS_CLASS_PATH = "jdk/classes:example/classes"
var EXT_CLASS_PATH = []string{"ext"}
var APP_CLASS_PATH []string
var MAIN_CLASS string

const MISC_LOG_FILE  = "log/misc.log"
const EXEC_LOG_FILE = "log/execution.log"
const CLASSLOAD_LOG_FILE = "log/classload.log"

var LOG = NewLog("misc", DEBUG, MISC_LOG_FILE)
var EXEC_LOG = NewLog("execution", INFO, EXEC_LOG_FILE)
var CLASSLOAD_LOG = NewLog("classload", INFO, CLASSLOAD_LOG_FILE)

type StringPool map[string]JavaLangString

func (this StringPool) GetStringInPool(str string) (JavaLangString, bool) {
	strObj, found := this[str];
	if strObj == nil {
		strObj = NULL
	}
	return strObj, found
}

func (this StringPool) InternString(stringobj JavaLangString) JavaLangString {
	str := stringobj.toNativeString()

	if strObj, found := this[str]; found {
		return strObj
	} else {
		this[str] = stringobj
		return stringobj
	}
}

type SystemProperties map[string]string

func (this SystemProperties) InitSystemProperties() SystemProperties {
	m := map[string]string {
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

	for k, v := range m {
		this.SetSystemProperty(k, v)
	}

	currentPath, _ := filepath.Abs(filepath.Dir(os.Args[0])+"/..")
	this.SetSystemProperty("sun.java.command", strings.Join(os.Args, " "))
	this.SetSystemProperty("java.home", currentPath)
	this.SetSystemProperty("java.class.path", SYS_CLASS_PATH)

	this.SetSystemProperty("java.io.tmpdir", SYS_CLASS_PATH)
	this.SetSystemProperty("java.library.path", SYS_CLASS_PATH)
	this.SetSystemProperty("java.ext.dirs", SYS_CLASS_PATH)
	this.SetSystemProperty("java.endorsed.dirs", SYS_CLASS_PATH)

	user, _ := user.Current()
	this.SetSystemProperty("user.name", user.Name)
	this.SetSystemProperty("user.home", user.HomeDir)
	this.SetSystemProperty("user.country", "NZ")
	this.SetSystemProperty("user.language", "en")
	this.SetSystemProperty("user.timezone", "")
	this.SetSystemProperty("user.dir", user.HomeDir)
	this.SetSystemProperty("sun.boot.library.path", "")
	this.SetSystemProperty("sun.boot.class.path", "")

	return this
}

func (this SystemProperties) ListSystemProperties() map[string]string {
	return this
}

func (this SystemProperties) SetSystemProperty(prop string, value string)  {
	this[prop] = value
}

func (this SystemProperties) GetSystemProperty(prop string) string {
	return this[prop]
}

const MAX_THREADS = 1000

type Core struct {}

type Universe struct {
	SystemProperties
	NativeMethodRegistry
	*Heap
	*ThreadManager
	StringPool
	*ClassLoader
	*Core
}

var VM = &Universe{
	make(map[string]string),
	make(map[string]reflect.Value),
	&Heap{},
	&ThreadManager{cmap.New()},
	make(map[string]JavaLangString),
	&ClassLoader{
			NewClassPath(SYS_CLASS_PATH),
			cmap.New(),//make(map[string]*Class),
			nil,
			0},
	&Core{},
}

const MAIN_METHOD_NAME = "main"
const MAIN_METHOD_DESCRIPTOR = "([Ljava/lang/String;)V"

var VM_WG = &sync.WaitGroup{}
func (this *Universe) Startup(initialClassName string, args... string)  {

	VM.InitSystemProperties()
	VM.RegisterNatives()

	// bootstrap thread don't run in a new go routine, just in Go startup routine
	VM.RunBootstrapThread(func() {
		// welcome to the Java world
		// the Java journey starts here
		VM.InvokeMethodOf("java/lang/System", "initializeSystemClass", "()V")

		// initial a thread
		VM.NewThread("main", func() {
			// As per jvm specification, initial main method needs to initialize initial class
			argsArr := VM.NewArrayOfName("[Ljava/lang/String;", Int(len(args)))
			for i, arg := range args {
				argsArr.SetArrayElement(Int(i), VM.NewJavaLangString(arg))
			}
			VM.InvokeMethodOf(initialClassName, MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR, argsArr)
		},
			func() {
				VM.exitDaemonThreads()
			}).start()
	})

	VM_WG.Wait()
}






