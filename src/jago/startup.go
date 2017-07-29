package jago

import "sync"

const MAIN_METHOD_NAME = "main"
const MAIN_METHOD_DESCRIPTOR = "([Ljava/lang/String;)V"

var VM_WG = &sync.WaitGroup{}

func Startup(initialClassName string, args... string)  {

	RegisterNatives()

	// bootstrap thread don't run in a new go routine, just in Go startup routine
	THREAD_MANAGER.RunBootstrapThread(func() {
		// welcome to the Java world
		// the Java journey starts here
		VM_invokeMethod0("java/lang/System", "initializeSystemClass", "()V")

		// initial a thread
		THREAD_MANAGER.NewThread("main", func() {
			// As per jvm specification, initial main method needs to initialize initial class
			argsArr := NewArray("[Ljava/lang/String;", Int(len(args)))
			for i, arg := range args {
				argsArr.SetElement(Int(i), NewJavaLangString(arg))
			}
			VM_invokeMethod0(initialClassName, MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR, argsArr)
		},
		func() {
			THREAD_MANAGER.exitDaemonThreads()
		}).start()
	})

	VM_WG.Wait()
}
