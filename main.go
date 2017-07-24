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
	var cp = flag.String("classpath", "/Users/Chao/Dropbox/Projects/jago-showcase/build/classes/main", "Class path")
	flag.Parse()
	jago.APP_CLASS_PATH = strings.Split(*cp, ":")


	//case0()
	//case1()
	case2()
}

func case0()  {
	bytes, _ := ioutil.ReadFile("/Users/Chao/Dropbox/Projects/jago-showcase/build/classes/main/case1/Pyramid.class")


	classfile := jago.NewClassFile(bytes)

	classfile.Print()
}

func case1()  {
	jago.Startup("case1/Pyramid")
}

func case2()  {
	jago.Startup("case2/Calendar", jago.NewJavaLangString("7"), jago.NewJavaLangString("2007"))
}


