package jago

func Startup(initialClassName string)  {
	initialClass := BOOTSTRAP_CLASSLOADER.CreateClass(initialClassName).(*Class)
	Execute(initialClass)
}
