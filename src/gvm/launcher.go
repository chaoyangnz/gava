package gvm

func Start(mainClassFile string)  {

	run(bootstrapClassLoader.load(mainClassFile))
}
