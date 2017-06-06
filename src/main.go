package main

import (
	"gvm"
	"flag"
	"strings"
)

var classpath []string

func main()  {
	var cp = flag.String("classpath", "/Library/Java/JavaVirtualMachines/jdk1.8.0_131.jdk/Contents/Home/jre/lib/rt.jar", "Class path")
	flag.Parse()
	classpath = strings.Split(*cp, ":")

	gvm.Start("Main.class")
}
