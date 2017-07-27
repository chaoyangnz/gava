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


	//jago.Startup("test_system_properties")
	//jago.Startup("test_athrow")
	//jago.Startup("test_static_double")
	//jago.Startup("test_shift")
	//jago.Startup("showcases/Pyramid")
	//jago.Startup("showcases/Calendar", "7", "2017")
	//jago.Startup("showcases/TreeInOrderTraverse")
	//jago.Startup("showcases/TreeLevelOrderTraverse")
	jago.Startup("showcases/EightQueens")
}

func case0()  {
	bytes, _ := ioutil.ReadFile("/Users/Chao/Dropbox/Projects/jago-showcase/build/classes/main/case1/Pyramid.class")


	classfile := jago.NewClassFile(bytes)

	classfile.Print()
}


