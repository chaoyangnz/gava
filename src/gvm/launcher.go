package gvm

import "io/ioutil"

func Start(mainClassFile string)  {
	bytes, _ := ioutil.ReadFile(mainClassFile)
	cr := NewClassReader(bytes)
	cf := cr.ReadAsClassFile()

	cf.Print()

	cm := NewClassMirror()
	cm.Load(cf)

	run(cm)
}
