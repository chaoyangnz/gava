package jago

const MAIN_METHOD_NAME = "main"
const MAIN_METHOD_DESCRIPTOR = "([Ljava/lang/String;)V"

func Startup(initialClassName string, args... string)  {
	RegisterNatives()
	thread := THREAD_MANAGER.NewThread("main")

	//systemClass := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/System")
	////systemClass.Initialize(thread)
	//initializeSystemClassMethod := systemClass.GetMethod("initializeSystemClass", "()V")
	//VM_invokeMethod(thread, initializeSystemClassMethod)


	//systemClass := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/System")
	//systemClassClinits := systemClass.Initialize()
	//for _, clinit := range systemClassClinits { thread.enqueueFrame(NewStackFrame(clinit))}
	////
	//initializeSystemClassMethod := systemClass.GetMethod("initializeSystemClass", "()V")
	//thread.enqueueFrame(NewStackFrame(initializeSystemClassMethod))



	initialClass := BOOTSTRAP_CLASSLOADER.CreateClass(initialClassName)
	//initialClass.Initialize(thread)
	// As per jvm specification, initial main method needs to initialize initial class
	//initialClass.Link0()
	//initialClassClinits := initialClass.Initialize()
	//for _, clinit := range initialClassClinits { thread.enqueueFrame(NewStackFrame(clinit))}

	mainMethod := initialClass.GetMethod(MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR)
	argsArr := BOOTSTRAP_CLASSLOADER.CreateClass("[Ljava/lang/String;").NewArray(Int(len(args)))
	for i, arg := range args {
		argsArr.SetElement(Int(i), NewJavaLangString(arg))
	}
	VM_invokeMethod(thread, mainMethod, argsArr)
}
