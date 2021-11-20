package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"gopkg.in/errgo.v2/errors"

	"SonosStandup/sonosAPI"
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
	rand.Seed(time.Now().UnixNano())

	device, err := sonosAPI.NewSonosDevice(configData.SonosIP)
	if err != nil {
		return errors.Because(err, nil, "could not connect to sonos")
	}

	stateData.PreviousVolume, err = device.GetVolume()
	if err != nil {
		return errors.Because(err, nil, "could not get volume from sonos")
	}

	state, err := device.GetPlaybackState()
	if err != nil {
		return errors.Because(err, nil, "could not get playback state from sonos")
	}

	stateData.PreviousState = *state
	if *state == sonosAPI.PlaybackPlaying {
		err = device.DoPause()
		if err != nil {
			return errors.Because(err, nil, "could not pause sonos")
		}
	}

	noSongs := len(stateData.AvailableTracks)

	if noSongs == 0 {
		fmt.Println("All songs played, rotating the list")
		stateData.AvailableTracks = stateData.AllTracks
		stateData.PlayedTracks = []string{}
		noSongs = len(stateData.AvailableTracks)
	}

	songIndex := int(math.Trunc(rand.Float64() * float64(noSongs)))
	if songIndex > (noSongs -1) {
		songIndex = noSongs - 1
	}

	err = device.SetPlaybackURI(stateData.AvailableTracks[songIndex])
	if err != nil {
		return errors.Because(err, nil, "could not set Media URI")
	}

	err = device.SetVolume(configData.Volume)
	if err != nil {
		return errors.Because(err, nil, "could not set Volume")
	}

	err = device.DoPlay()
	if err != nil {
		return errors.Because(err, nil, "could not play sonos")
	}

	stateData.LastTrack = stateData.AvailableTracks[songIndex]
	stateData.PlayedTracks = append(stateData.PlayedTracks, stateData.LastTrack)
	songIndex = 0
	for _, track := range stateData.AvailableTracks {
		if track != stateData.LastTrack {
			stateData.AvailableTracks[songIndex] = track
			songIndex++
		}
	}
	stateData.AvailableTracks = stateData.AvailableTracks[:songIndex]

	err = stateData.Save()
	if err != nil {
		return errors.Because(err, nil, "could not save state")
	}

	fmt.Printf("Playing %s\n", stateData.LastTrack)

	return nil
}