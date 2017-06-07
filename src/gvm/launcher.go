package gvm

func Start(mainClassFile string)  {
	cr := NewClassReader(mainClassFile)
	cf := cr.ReadAsClassFile()

	//cf.Print()

	cm := &JavaClass{}
	cm.Load(cf)

	run(cm)
}
