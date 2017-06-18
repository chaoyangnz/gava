package jago

func Startup(initialClassName string)  {
	thread := THREAD_MANAGER.NewThread("main")
	initialClass := BOOTSTRAP_CLASSLOADER.CreateClass(initialClassName).(*Class)
	mainMethod := initialClass.FindMethod(MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR)
	thread.pushFrame(NewStackFrame(mainMethod))
	// As per jvm specification, initial main method needs to initialize initial class
	initialClass.Link()
	initialClass.Initialize(thread)

	thread.Run()
}
