package main

import (
	"SonosStandup/sonosAPI"
	"gopkg.in/errgo.v2/errors"
)

type TestCommand struct {
}

var testCommand TestCommand

func init () {
	_, err := parser.AddCommand(
		"test",
		"Test",
		"Run a test command",
		&testCommand)
	if err != nil {
		panic(err)
	}
}

func (x *TestCommand)Execute(args[]string) error {
	device, err := sonosAPI.NewSonosDevice(configData.SonosIP)
	if err != nil {
		return errors.Because(err, nil, "could not connect to sonos")
	}

	position, err := device.GetPositionInfo()
	if err != nil {
		return errors.Because(err, nil, "could not set previous volume")
	}

	_ = position

	return nil
}