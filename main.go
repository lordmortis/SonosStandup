package main

import (
	"SonosStandup/sonosAPI"
	"fmt"
)

func main() {
	device, err := sonosAPI.NewSonosDevice("172.17.172.69")
	if err != nil {
		fmt.Printf("Could not connect to sonos: %s\n", err)
		return
	}

	/*
	initialVolume, err := device.GetVolume()
	if err != nil {
		fmt.Printf("Could not get volume from sonos: %s\n", err)
		return
	}

	_ = initialVolume

	state, err := device.GetPlaybackState()
	if err != nil {
		fmt.Printf("Unable to get playback state: %s\n", err)
		return
	}

	_ = initialState

	err = device.DoPause()
	if err != nil {
		fmt.Printf("Could not pause: %s\n", err)
		return
	}*/

	err = device.SetPlaybackURI("http://unity-addressables.int.viewport.com.au/Standup-Stingers/Lacuna%20Coil-Our%20Truth.wav")
	if err != nil {
		fmt.Printf("Could not set media URI: %s\n", err)
		return
	}
}
