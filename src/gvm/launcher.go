package gvm

func Start(mainClassFile string)  {

	run(ofClassType(mainClassFile))
	print(typeCache)
}
