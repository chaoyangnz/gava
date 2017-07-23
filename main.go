package main

import (
	//"flag"
	//"strings"
	"jago"
	"io/ioutil"
	"flag"
	"strings"
)



func main()  {
	var cp = flag.String("classpath", "/Users/Charvis/Dropbox/Projects/gvm-java/out/production/gvm-java", "Class path")
	flag.Parse()
	jago.APP_CLASS_PATH = strings.Split(*cp, ":")

	jago.Startup("Main")
}

//func main() {
//	testClassFileModule()
//}

func testClassFileModule()  {
	bytes, _ := ioutil.ReadFile("/Users/Chao/Dropbox/Projects/gvm-java/out/production/gvm-java/Main.class")


	classfile := jago.NewClassFile(bytes)

	classfile.Print()
}
