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

	_ = initialVolume

	err = device.DoPause()
	if err != nil {
		fmt.Printf("Could not pause: %s\n", err)
		return
	}
}
