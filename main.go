package main

import (
	"flag"
	"strings"
	"jago"
)



func main()  {
	var cp = flag.String("classpath", "/Users/Charvis/Dropbox/Projects/gvm-java/out/production/gvm-java", "Class path")
	flag.Parse()
	jago.APP_CLASS_PATH = strings.Split(*cp, ":")

	jago.Startup("Main")
}
