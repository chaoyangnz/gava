package main

import (
	"jago"
	"github.com/urfave/cli"
	"fmt"
	"os"
	"strings"
	"strconv"
	"path"
)

func main() {
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
	var nologo bool = false
	var log_thread string
	var log_classloader string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "classpath, cp",
			Value:       "",
			Usage:       "application classpath separated by colon",
			Destination: &classpath,
		},
		cli.BoolFlag{
			Name:        "nologo, nl",
			Usage:       "don't show logo",
			Destination: &nologo,
		},
		cli.StringFlag{
			Name:        "log:thread",
			Value:       "",
			Usage:       "log level of instruction execution in a thread context, options: info, debug, trace",
			Destination: &log_thread,
		},
		cli.StringFlag{
			Name:        "log:classloader",
			Value:       "",
			Usage:       "log level of class loading, options: info, debug, trace",
			Destination: &log_classloader,
		},
	}

	app.Action = func(c *cli.Context) error {
		if !nologo {
			fmt.Printf("    _                   \r\n   (_) __ _  __ _  ___  \r\n   | |/ _` |/ _` |/ _ \\ \r\n   | | (_| | (_| | (_) |   version %s\r\n  _/ |\\__,_|\\__, |\\___/    \r\n |__/       |___/     \n", app.Version)
		}
		fmt.Printf("\n\nCommand: %s \n", strings.Join(os.Args, " "))

		args := c.Args()
		if c.NArg() < 1 {
			return nil
		}

		jago_home := os.Getenv("JAGO_HOME")
		if jago_home == "" {
			fmt.Println("Please set JAGO_HOME environment variable first.")
			return nil
		} else if !path.IsAbs(jago_home) {
			fmt.Println("JAGO_HOME needs to be absolute path.")
			return nil
		}

		jago.VM.Init()

		jagoClassName := args.Get(0)
		jagoArgs := make([]string, c.NArg()-1)
		for i := 0; i < len(args)-1; i++ {
			jagoArgs[i] = args[i+1]
		}

		if classpath != "" {
			fmt.Printf("Add a new classpath: %s\n", classpath)
			jago.VM.SetSystemSetting("classpath.application", classpath)
		}

		switch log_thread {
		case "error":
			jago.VM.SetSystemSetting("log.level.thread", strconv.Itoa(jago.ERROR))
		case "warn":
			jago.VM.SetSystemSetting("log.level.thread", strconv.Itoa(jago.WARN))
		case "info":
			jago.VM.SetSystemSetting("log.level.thread", strconv.Itoa(jago.INFO))
		case "debug":
			jago.VM.SetSystemSetting("log.level.thread", strconv.Itoa(jago.DEBUG))
		case "trace":
			jago.VM.SetSystemSetting("log.level.thread", strconv.Itoa(jago.TRACE))
		default:
			if log_thread != "" {
				fmt.Printf("Invalid log:thread option: %s\n", log_thread)
			}
		}

		switch log_classloader {
		case "error":
			jago.VM.SetSystemSetting("log.level.classloader", strconv.Itoa(jago.ERROR))
		case "warn":
			jago.VM.SetSystemSetting("log.level.classloader", strconv.Itoa(jago.WARN))
		case "info":
			jago.VM.SetSystemSetting("log.level.classloader", strconv.Itoa(jago.INFO))
		case "debug":
			jago.VM.SetSystemSetting("log.level.classloader", strconv.Itoa(jago.DEBUG))
		case "trace":
			jago.VM.SetSystemSetting("log.level.classloader", strconv.Itoa(jago.TRACE))
		default:
			if log_classloader != "" {
				fmt.Printf("Invalid log:classloader option: %s\n", log_classloader)
			}
		}

		fmt.Println("------------------------------------------------------------\n")

		jago.VM.Startup(strings.Replace(jagoClassName, ".", "/", -1), jagoArgs...)
		return nil
	}

	app.Run(os.Args)
}
