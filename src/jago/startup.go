package jago

const MAIN_METHOD_NAME = "main"
const MAIN_METHOD_DESCRIPTOR = "([Ljava/lang/String;)V"

func Startup(initialClassName string, args... string)  {
	defer func() { // finally check uncaught exception
		r := recover()
		if r != nil {
			throwable, ok := r.(Reference)
			if ok {
				handleUncaughtException(throwable)
			}
		}
	}()

	RegisterNatives()

	// initial a thread
	thread := THREAD_MANAGER.NewThread("main")

	systemClass := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/System")
	initializeSystemClassMethod := systemClass.GetMethod("initializeSystemClass", "()V")
	VM_invokeMethod(thread, initializeSystemClassMethod)

	// As per jvm specification, initial main method needs to initialize initial class
	initialClass := BOOTSTRAP_CLASSLOADER.CreateClass(initialClassName)
	mainMethod := initialClass.GetMethod(MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR)
	argsArr := BOOTSTRAP_CLASSLOADER.CreateClass("[Ljava/lang/String;").NewArray(Int(len(args)))
	for i, arg := range args {
		argsArr.SetElement(Int(i), NewJavaLangString(arg))
	}
	VM_invokeMethod(thread, mainMethod, argsArr)
}
