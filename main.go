package main

import (
	"SonosStandup/state"
	"gopkg.in/errgo.v2/errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/jessevdk/go-flags"

	"SonosStandup/config"
)

type Options struct {
	ConfigFile string `long:"configFile" description:"path to config.yaml file"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)
var configData *config.Config
var stateData state.Data

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	parser.CommandHandler = func(command flags.Commander, args[]string) error {
		var err error

		configData, err = config.Load(&options.ConfigFile)
		if err != nil {
			return errors.Because(err, nil, "unable to parse config file")
		}

		statePath := filepath.Join(configData.StatePath, "state.data")
		if _, err = os.Stat(statePath); err == nil {
			err = stateData.Load(filepath.Join(configData.StatePath, "state.data"))
			if err != nil {
				return errors.Because(err, nil, "unable to read state")
			}
		} else {
			stateData = state.New(statePath)
		}

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
