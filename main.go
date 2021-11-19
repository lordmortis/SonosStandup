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

	initialVolume, err := device.GetVolume()
	if err != nil {
		fmt.Printf("Could not get volume from sonos: %s\n", err)
		return
	}

	fmt.Printf("Volume is %v\n", initialVolume)

	err = device.DoPlay()
	if err != nil {
		fmt.Printf("Could not play: %s\n", err)
		return
	}

	state, err := device.GetPlaybackState()
	if err != nil {
		fmt.Printf("Unable to get playback state: %s\n", err)
		return
	}

	fmt.Printf("State is %s\n", *state)

	err = device.DoPause()
	if err != nil {
		fmt.Printf("Could not pause: %s\n", err)
		return
	}

	state, err = device.GetPlaybackState()
	if err != nil {
		fmt.Printf("Unable to get playback state: %s\n", err)
		return
	}

	fmt.Printf("State is %s\n", *state)
}
