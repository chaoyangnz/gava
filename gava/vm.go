package gava

import (
	"os"
	"path"
	"reflect"
	"strconv"
	"sync"
)

type SystemProperties map[string]string

func (this SystemProperties) SetSystemProperty(key string, value string) {
	this[key] = value
}

func (this SystemProperties) GetSystemProperty(key string) string {
	return this[key]
}

type Gava struct {
	SystemProperties

	*ExecutionEngine
	*MethodArea
	*Heap
	*OS

	*LoggerFactory
	*Logger
}

var VM = NewVM()

func NewVM() *Gava {
	vm := &Gava{}
	gavaHome := os.Getenv("GAVA_HOME")
	vm.SystemProperties = map[string]string{
		"log.base":              path.Join(gavaHome, "log"),
		"log.level.threads":     strconv.Itoa(WARN),
		"log.level.thread":      strconv.Itoa(WARN),
		"log.level.classloader": strconv.Itoa(WARN),
		"log.level.io":          strconv.Itoa(WARN),
		"log.level.misc":        strconv.Itoa(WARN),

		"classpath.system":      path.Join(gavaHome, "jdk/classes"),
		"classpath.extension":   "",
		"classpath.application": "",
	}

	vm.LoggerFactory = &LoggerFactory{}

	return vm
}

// Before vm initialization, all lot of system settings can be set.
func (this *Gava) Init() {
	natives := make(map[string]reflect.Value)

	threadsLogLevel, _ := strconv.Atoi(this.GetSystemProperty("log.level.threads"))
	ioLogLevel, _ := strconv.Atoi(this.GetSystemProperty("log.level.io"))
	this.ExecutionEngine = &ExecutionEngine{
		make([]Instruction, JVM_OPC_MAX+1),
		natives,
		sync.Map{},
		this.NewLogger("threads", threadsLogLevel, "threads.log"),
		this.NewLogger("io", ioLogLevel, "io.log")}
	this.RegisterInstructions()
	this.RegisterNatives()

	this.Heap = &Heap{}

	classloaderLogLevel, _ := strconv.Atoi(this.GetSystemProperty("log.level.classloader"))
	systemClasspath := VM.GetSystemProperty("classpath.system")
	this.MethodArea = &MethodArea{
		make(map[NL]*Class),
		make(map[NL]*Class),
		make(map[string]JavaLangString),
		&BootstrapClassLoader{
			NewClassPath(systemClasspath),
			this.NewLogger("classloader", classloaderLogLevel, "classloader.log"),
		},
	}

	this.OS = &OS{}

	miscLogLevel, _ := strconv.Atoi(this.GetSystemProperty("log.level.misc"))
	this.Logger = this.LoggerFactory.NewLogger("misc", miscLogLevel, "misc.log")
}

const MAIN_METHOD_NAME = "main"
const MAIN_METHOD_DESCRIPTOR = "([Ljava/lang/String;)V"

var VM_WG = &sync.WaitGroup{}

func (this *Gava) Startup(initialClassName string, args ...string) {

	// bootstrap thread don't run in a new go routine, just in Launch startup routine
	this.RunBootstrapThread(func() {
		// welcome to the Java world
		// the Java journey starts here
		VM.InvokeMethodOf("java/lang/System", "initializeSystemClass", "()V")

		// Use AppClassLoader to load initial class
		systemClassLoaderObject := VM.InvokeMethodOf("java/lang/ClassLoader", "getSystemClassLoader", "()Ljava/lang/ClassLoader;").(JavaLangClassLoader)
		initialClass := VM.createClass(initialClassName, systemClassLoaderObject, TRIGGER_BY_ACCESS_MEMBER)
		mainMethod := initialClass.FindMethod(MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR)
		if mainMethod == nil {
			Fatal("main method not found")
		}

		// initial a thread
		this.NewThread("main",
			func() {
				// As per jvm specification, initial main method needs to initialize initial class
				argsArr := VM.NewArrayOfName("[Ljava/lang/String;", Int(len(args)))
				for i, arg := range args {
					argsArr.SetArrayElement(Int(i), VM.NewJavaLangString(arg))
				}
				this.InvokeMethod(mainMethod, argsArr)
			},
			func() {
				this.exitDaemonThreads()
			}).start()
	})

	VM_WG.Wait()
}
