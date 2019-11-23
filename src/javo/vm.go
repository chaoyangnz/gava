package javo

import (
	"github.com/orcaman/concurrent-map"
	"sync"
	"reflect"
	"strconv"
	"os"
	"path"
)

type SystemSettings map[string]string

func (this SystemSettings) SetSystemSetting(key string, value string) {
	this[key] = value
}

func (this SystemSettings) GetSystemSetting(key string) string {
	return this[key]
}

type Javo struct {
	SystemSettings

	*ExecutionEngine
	*MethodArea
	*Heap
	*OS

	*LoggerFactory
	*Logger
}

var VM = NewVM()

func NewVM() *Javo {
	vm := &Javo{}
	javoHome := os.Getenv("JAVO_HOME")
	vm.SystemSettings = map[string]string{
		"log.base":              path.Join(javoHome, "log"),
		"log.level.threads":     strconv.Itoa(WARN),
		"log.level.thread":      strconv.Itoa(WARN),
		"log.level.classloader": strconv.Itoa(WARN),
		"log.level.io":          strconv.Itoa(WARN),
		"log.level.misc":        strconv.Itoa(WARN),

		"classpath.system":      path.Join(javoHome, "jdk/classes"),
		"classpath.extension":   "",
		"classpath.application": "",
	}

	vm.LoggerFactory = &LoggerFactory{}

	return vm
}

// Before vm initialization, all lot of system settings can be set.
func (this *Javo) Init() {
	natives := make(map[string]reflect.Value)

	threadsLogLevel, _ := strconv.Atoi(this.GetSystemSetting("log.level.threads"))
	ioLogLevel, _ := strconv.Atoi(this.GetSystemSetting("log.level.io"))
	this.ExecutionEngine = &ExecutionEngine{
		make([]Instruction, JVM_OPC_MAX+1),
		natives,
		cmap.New(),
		this.NewLogger("threads", threadsLogLevel, "threads.log"),
		this.NewLogger("io", ioLogLevel, "io.log")}
	this.RegisterInstructions()
	this.RegisterNatives()

	this.Heap = &Heap{}

	classloaderLogLevel, _ := strconv.Atoi(this.GetSystemSetting("log.level.classloader"))
	systemClasspath := VM.GetSystemSetting("classpath.system")
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

	miscLogLevel, _ := strconv.Atoi(this.GetSystemSetting("log.level.misc"))
	this.Logger = this.LoggerFactory.NewLogger("misc", miscLogLevel, "misc.log")
}

const MAIN_METHOD_NAME = "main"
const MAIN_METHOD_DESCRIPTOR = "([Ljava/lang/String;)V"

var VM_WG = &sync.WaitGroup{}

func (this *Javo) Startup(initialClassName string, args ... string) {

	// bootstrap thread don't run in a new go routine, just in Go startup routine
	VM.RunBootstrapThread(func() {
		// welcome to the Java world
		// the Java journey starts here
		VM.InvokeMethodOf("java/lang/System", "initializeSystemClass", "()V")

		// Use AppClassLoader to load initial class
		systemClassLoaderObject := VM.InvokeMethodOf("java/lang/ClassLoader", "getSystemClassLoader", "()Ljava/lang/ClassLoader;").(JavaLangClassLoader)
		initialClass := VM.createClass(initialClassName, systemClassLoaderObject, TRIGGER_BY_ACCESS_MEMBER)
		mainMethod := initialClass.FindMethod(MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR)

		// initial a thread
		VM.NewThread("main",
			func() {
				// As per jvm specification, initial main method needs to initialize initial class
				argsArr := VM.NewArrayOfName("[Ljava/lang/String;", Int(len(args)))
				for i, arg := range args {
					argsArr.SetArrayElement(Int(i), VM.NewJavaLangString(arg))
				}
				VM.InvokeMethod(mainMethod, argsArr)
			},
			func() {
				VM.exitDaemonThreads()
			}).start()
	})

	VM_WG.Wait()
}
