package main

import (
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	ConfigFile string `long:"configFile" description:"path to config.yaml file"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	parser.CommandHandler = func(command flags.Commander, args[]string) error {
		return command.Execute(args)
	}

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
