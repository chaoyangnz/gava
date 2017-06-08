package main

import (
	"gvm"
	"flag"
	"strings"
)



func main()  {
	var cp = flag.String("classpath", "/Users/Charvis/Dropbox/Projects/gvm-java/out/production/gvm-java", "Class path")
	flag.Parse()
	gvm.AppClasspath = strings.Split(*cp, ":")

	gvm.Start("Main")
}
