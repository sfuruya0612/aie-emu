package main

import (
	"fmt"
	"os"

	"github.com/sfuruya0612/aie-emu/cmd"
	"github.com/urfave/cli"
)

const version = "20.2.1"

var (
	date      string
	hash      string
	goversion string
)

func main() {
	app := cli.NewApp()

	app.EnableBashCompletion = true
	app.Name = "aie-emu"
	app.Usage = ""

	if date != "" || hash != "" || goversion != "" {
		app.Version = fmt.Sprintf("%s %s (Build by: %s)", date, hash, goversion)
	} else {
		app.Version = version
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "profile, p",
			EnvVar: "AWS_PROFILE",
			Value:  "default",
			Usage:  "AWS credential (~/.aws/config) or read AWS_PROFILE environment variable",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "stdout",
			Usage: "",
		},
	}

	app.Action = cmd.GetIamList

	if err := app.Run(os.Args); err != nil {
		code := 1
		if c, ok := err.(cli.ExitCoder); ok {
			code = c.ExitCode()
		}
		fmt.Printf("Err: %v", err.Error())
		os.Exit(code)
	}
}
