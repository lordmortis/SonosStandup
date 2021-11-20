package main

import (
	"SonosStandup/sonosAPI"
	"gopkg.in/errgo.v2/errors"
)

type RunStandupCommand struct {
}

var runStandupCommand RunStandupCommand

func init() {
	_, err := parser.AddCommand("runStandup", "Run Standup Music", "Run Standup Music", &runStandupCommand)
	if err != nil {
		panic(err)
	}
}

func (x *RunStandupCommand)Execute(args []string) error {
	device, err := sonosAPI.NewSonosDevice(configData.SonosIP)
	if err != nil {
		return errors.Because(err, nil, "could not connect to sonos")
	}

	initialVolume, err := device.GetVolume()
	if err != nil {
		return errors.Because(err, nil, "could not get volume from sonos")
	}

	_ = initialVolume
	//TODO: write initial volume to state file

	state, err := device.GetPlaybackState()
	if err != nil {
		return errors.Because(err, nil, "could not get playback state from sonos")
	}

	if *state == sonosAPI.PlaybackPlaying {
		//TODO: write that we were playing to statefile
		err = device.DoPause()
		if err != nil {
			return errors.Because(err, nil, "could not pause sonos")
		}
	} else {
		//TODO: write that we were stopped to statefile
	}

	//TODO: Pick the song / reset the songlist
	err = device.SetPlaybackURI("http://unity-addressables.int.viewport.com.au/Standup-Stingers/Darude-Sandstorm.flac")
	if err != nil {
		return errors.Because(err, nil, "could not set Media URI")
	}

	//TODO: Get volume from config file
	err = device.SetVolume(50)
	if err != nil {
		return errors.Because(err, nil, "could not set Media URI")
	}

	err = device.DoPlay()
	if err != nil {
		return errors.Because(err, nil, "could not play sonos")
	}

	return nil
}