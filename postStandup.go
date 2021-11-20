package main

import "SonosStandup/sonosAPI"

type PostStandupCommand struct {
}

var postStandupCommand PostStandupCommand

func init() {
	_, err := parser.AddCommand("postStandup",
		"Post Standup",
		"Resume whatever the sonos was doing before standup", &postStandupCommand)
	if err != nil {
		panic(err)
	}
}

func (x *PostStandupCommand)Execute(args []string) error {
	device, err := sonosAPI.NewSonosDevice(configData.SonosIP)
	if err != nil {
		return errors.Because(err, nil, "could not connect to sonos")
	}

	err = device.SetVolume(stateData.PreviousVolume)
	if err != nil {
		return errors.Because(err, nil, "could not set previous volume")
	}

	//TODO: fix the queue thing, resume playing
	return nil
}