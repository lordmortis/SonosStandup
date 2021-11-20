package main

import "fmt"

type InfoCommand struct {
}

var infoCommand InfoCommand

func init () {
	_, err := parser.AddCommand(
		"info",
		"Info",
		"Get Info about the current standup state",
		&infoCommand)
	if err != nil {
		panic(err)
	}
}

func (x *InfoCommand)Execute(args[]string) error {
	fmt.Printf("Last Played Track: %s\n", stateData.LastTrack)
	fmt.Printf("Volume before last playback: %d\n", stateData.PreviousVolume)
	fmt.Printf("State before last playback: %s\n", stateData.PreviousState)
	fmt.Printf("Position in queue before last playback: %d\n", stateData.PreviousQueue)

	fmt.Printf("Tracks available to be Played:\n")
	for _, track := range stateData.AvailableTracks {
		fmt.Printf("\t%s\n", track)
	}

	fmt.Printf("All tracks:\n")
	for _, track := range stateData.AllTracks {
		fmt.Printf("\t%s\n", track)
	}

	return nil
}