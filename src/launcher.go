package main

import (
	"io/ioutil"
	"classfile"
	"flag"
	"strings"
)

var classpath []string

func main()  {
	var cp = flag.String("classpath", "/Library/Java/JavaVirtualMachines/jdk1.8.0_131.jdk/Contents/Home/jre/lib/rt.jar", "Class path")
	flag.Parse()
	classpath = strings.Split(*cp, ":")

	bytes, _ := ioutil.ReadFile("Main.class")
	cr := classfile.NewClassReader(bytes)
	cf := cr.ReadAsClassFile()

	//cf.Print()

	cm := classfile.NewClassMirror(cf)
	cm.FindMethod(classfile.SignatureOf("main:V()"))
	print(1)
}
