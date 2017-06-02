package main

import (
	"io/ioutil"
	"classfile"
)

func main()  {
	bytes, _ := ioutil.ReadFile("/Users/Charvis/Solution.class")
	cr := classfile.NewClassReader(bytes)
	cf := cr.ReadAsClassFile()

	cf.Print()
}
