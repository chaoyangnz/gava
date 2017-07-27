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
	//case2()
	//testAthrow()
	//print_system_properties()
	test_static_double()
	//test_tree_traverse()
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
	jago.Startup("case2/Calendar", "7", "2017")
}

func print_system_properties()  {
	jago.Startup("test_system_properties")
}

func testAthrow()  {
	jago.Startup("test_athrow")
}

func test_tree_traverse()  {
	jago.Startup("_102_binary_tree_level_order_traversal/Tests")
}

func test_static_double()  {
	jago.Startup("test_static_double")
}


