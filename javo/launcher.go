package javo

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/profile"
	"github.com/urfave/cli"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func Launch() {
	app := cli.NewApp()
	app.Name = "Javo"
	app.Usage = "A simplified Java Virtual Machine for the educational purpose"
	app.Description = ""
	app.UsageText = "javo [-options] class [args...]"
	app.Version = "1.0.0"
	app.Author = "Chao Yang"
	app.Email = "chao@yang.to"
	app.Description = "A Java Virtual Machine demonstrating the basic features of execution engine, class loading, type/value system, exception handling, native methods etc."

	var classpath string
	var noLogo = false
	var logThread string
	var logClassloader string
	var debug = false
	var trace = false
	var profiling = false
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "classpath, cp",
			Value:       "",
			Usage:       "application classpath separated by colon",
			Destination: &classpath,
		},
		cli.BoolFlag{
			Name:        "noLogo, nl",
			Usage:       "don't show logo",
			Destination: &noLogo,
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "debug mode",
			Destination: &debug,
		},
		cli.BoolFlag{
			Name:        "trace, t",
			Usage:       "trace mode",
			Destination: &trace,
		},
		cli.StringFlag{
			Name:        "log:thread",
			Value:       "",
			Usage:       "log level of instruction execution in a thread context, options: info, debug, trace",
			Destination: &logThread,
		},
		cli.StringFlag{
			Name:        "log:classloader",
			Value:       "",
			Usage:       "log level of class loading, options: info, debug, trace",
			Destination: &logClassloader,
		},
		cli.BoolFlag{
			Name:        "profile, p",
			Usage:       "profile javo",
			Destination: &profiling,
		},
	}

	app.Action = func(c *cli.Context) error {
		if !noLogo {
			color.HiMagenta("               __       ___   ____    ____  ______   \n              |  |     /   \\  \\   \\  /   / /  __  \\  \n              |  |    /  ^  \\  \\   \\/   / |  |  |  | \n        .--.  |  |   /  /_\\  \\  \\      /  |  |  |  |  version %s\n        |  `--'  |  /  _____  \\  \\    /   |  `--'  | \n         \\______/  /__/     \\__\\  \\__/     \\______/  ", app.Version)
		}
		commandEcho := fmt.Sprintf("Command: %s", strings.Join(os.Args, " "))
		color.Cyan("\n\n%s\n", commandEcho)

		args := c.Args()
		if c.NArg() < 1 {
			return nil
		}

		javoHome := os.Getenv("JAVO_HOME")
		if javoHome == "" {
			color.Yellow("Please set JAVO_HOME environment variable first.")
			return nil
		} else if !path.IsAbs(javoHome) {
			color.Yellow("JAVO_HOME needs to be absolute path.")
			return nil
		}

		javaClassName := args.Get(0)
		javaClassName = strings.TrimSuffix(javaClassName, ".class")
		javaArgs := make([]string, c.NArg()-1)
		for i := 0; i < len(args)-1; i++ {
			javaArgs[i] = args[i+1]
		}

		if classpath != "" {
			color.Blue("Add a new classpath: %s\n", classpath)
			VM.SetSystemProperty("classpath.application", classpath)
		}

		switch logThread {
		case "error":
			VM.SetSystemProperty("log.level.thread", strconv.Itoa(ERROR))
		case "warn":
			VM.SetSystemProperty("log.level.thread", strconv.Itoa(WARN))
		case "info":
			VM.SetSystemProperty("log.level.thread", strconv.Itoa(INFO))
		case "debug":
			VM.SetSystemProperty("log.level.thread", strconv.Itoa(DEBUG))
		case "trace":
			VM.SetSystemProperty("log.level.thread", strconv.Itoa(TRACE))
		default:
			if logThread != "" {
				color.Yellow("Invalid log:thread option: %s\n", logThread)
			}
		}

		switch logClassloader {
		case "error":
			VM.SetSystemProperty("log.level.classloader", strconv.Itoa(ERROR))
		case "warn":
			VM.SetSystemProperty("log.level.classloader", strconv.Itoa(WARN))
		case "info":
			VM.SetSystemProperty("log.level.classloader", strconv.Itoa(INFO))
		case "debug":
			VM.SetSystemProperty("log.level.classloader", strconv.Itoa(DEBUG))
		case "trace":
			VM.SetSystemProperty("log.level.classloader", strconv.Itoa(TRACE))
		default:
			if logClassloader != "" {
				color.Yellow("Invalid log:classloader option: %s\n", logClassloader)
			}
		}

		if debug {
			VM.SetSystemProperty("log.level.classloader", strconv.Itoa(DEBUG))
			VM.SetSystemProperty("log.level.thread", strconv.Itoa(DEBUG))
			VM.SetSystemProperty("log.level.io", strconv.Itoa(DEBUG))
			VM.SetSystemProperty("log.level.misc", strconv.Itoa(DEBUG))
		}

		if trace {
			VM.SetSystemProperty("log.level.classloader", strconv.Itoa(TRACE))
			VM.SetSystemProperty("log.level.thread", strconv.Itoa(TRACE))
			VM.SetSystemProperty("log.level.io", strconv.Itoa(TRACE))
			VM.SetSystemProperty("log.level.misc", strconv.Itoa(TRACE))
		}

		if profiling {
			// CPU profiling
			defer profile.Start(func(p *profile.Profile) {
				path := "./profile"
				profile.ProfilePath(path)(p)
				profile.CPUProfile(p)
				color.Yellow("Run go tool pprof --pdf /usr/local/bin/javo %s/cpu.pprof > %s/cpu.pdf", path, path)
			}).Stop()
			// Memory profiling
			//defer profile.Start(profile.MemProfile).Stop()
		}

		color.Green(strings.Repeat("·", len(commandEcho)) + "\n")

		VM.Init()
		VM.Startup(javaNameToBinaryName0(javaClassName), javaArgs...)
		return nil
	}

	t := time.Now()
	app.Run(os.Args)
	color.Yellow("\n\n%dms elapsed", time.Now().Sub(t).Milliseconds())
}
