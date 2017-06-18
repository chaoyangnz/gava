package jago

func Startup(initialClassName string)  {
	thread := THREAD_MANAGER.NewThread("main")
	//systemClass := BOOTSTRAP_CLASSLOADER.CreateClass("java/lang/System").(*Class)
	//systemClass.Link()
	//systemClassClinits := systemClass.Initialize()
	//for _, clinit := range systemClassClinits { thread.enqueueFrame(NewStackFrame(clinit))}
	//
	//initializeSystemClassMethod := systemClass.GetMethod("initializeSystemClass", "()V")
	//thread.enqueueFrame(NewStackFrame(initializeSystemClassMethod))



	initialClass := BOOTSTRAP_CLASSLOADER.CreateClass(initialClassName).(*Class)
	// As per jvm specification, initial main method needs to initialize initial class
	initialClass.Link()
	initialClassClinits := initialClass.Initialize()
	for _, clinit := range initialClassClinits { thread.enqueueFrame(NewStackFrame(clinit))}

	mainMethod := initialClass.GetMethod(MAIN_METHOD_NAME, MAIN_METHOD_DESCRIPTOR)
	thread.enqueueFrame(NewStackFrame(mainMethod))

	thread.Run()
}
