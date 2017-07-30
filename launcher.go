package main

import (
	"jago"
	"github.com/urfave/cli"
	"fmt"
	"os"
	"strings"
)



func main()  {
	app := cli.NewApp()
	app.Name = "Jago"
	app.Usage = "A simplified Java Virtual Machine for the educational purpose"
	app.Description = ""
	app.UsageText = "jago [-options] class [args...]"
	app.Version = "1.0.0"
	app.Author = "Chao Yang"
	app.Email = "richd.yang@gmail.com"
	app.Description = "A Java Virtual Machine demonstrating the basic features of execution engine, class loading, type/value system, exception handling, native methods etc."

	var classpath string
	var log_execution string
	var log_classload string
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "classpath, cp",
			Value: "",
			Usage: "application classpath separated by colon",
			Destination: &classpath,
		},
		cli.StringFlag{
			Name: "log:execution",
			Value: "",
			Usage: "log level of instruction execution, options: info, debug, trace",
			Destination: &log_execution,
		},
		cli.StringFlag{
			Name: "log:classload",
			Value: "",
			Usage: "log level of class loading, options: info, debug, trace",
			Destination: &log_classload,
		},
	}

	fmt.Printf("    _                   \r\n   (_) __ _  __ _  ___  \r\n   | |/ _` |/ _` |/ _ \\ \r\n   | | (_| | (_| | (_) |   version %s\r\n  _/ |\\__,_|\\__, |\\___/    \r\n |__/       |___/     \n\n\n", app.Version)
	fmt.Printf("Command: %s \n", strings.Join(os.Args, " "))

	app.Action = func(c *cli.Context) error {
		//fmt.Println(" ┬┌─┐┌─┐┌─┐")
		//fmt.Println(" │├─┤│ ┬│ │")
		//fmt.Println("└┘┴ ┴└─┘└─┘")


		args := c.Args()
		if c.NArg() < 1 {
			return nil
		}



		jagoClassName := args.Get(0)
		jagoArgs := make([]string, c.NArg()-1)
		for i := 0; i < len(args)-1; i++ {
			jagoArgs[i] = args[i+1]
		}


		if classpath != "" {
			fmt.Printf("Add a new classpath: %s\n", classpath)
			jago.SYS_CLASS_PATH += classpath
		}

		switch log_execution {
		case "info": jago.EXEC_LOG.Level = jago.INFO
		case "debug": jago.EXEC_LOG.Level = jago.DEBUG
		case "trace":  jago.EXEC_LOG.Level = jago.TRACE
		default:
			if log_execution != "" {
				fmt.Printf("Invalid log:execution option: %s\n", log_execution)
			}
		}

		switch log_classload {
		case "info": jago.CLASSLOAD_LOG.Level = jago.INFO
		case "debug": jago.CLASSLOAD_LOG.Level = jago.DEBUG
		case "trace":  jago.CLASSLOAD_LOG.Level = jago.TRACE
		default:
			if log_classload != "" {
				fmt.Printf("Invalid log:classload option: %s\n", log_classload)
			}
		}

		fmt.Println("------------------------------------------------------------\n")




		jago.VM.Startup(strings.Replace(jagoClassName, ".", "/", -1), jagoArgs...)
		return nil
	}

	app.Run(os.Args)
}

