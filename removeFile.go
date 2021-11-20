package main

import "gopkg.in/errgo.v2/errors"

type RemoveFilecommand struct {

}

var removeFilecommand RemoveFilecommand

func init () {
	_, err := parser.AddCommand(
		"removeFile",
		"Remove File",
		"Remove Audio file from playback",
		&removeFilecommand)
	if err != nil {
		panic(err)
	}
}

func (x *RemoveFilecommand)Execute(args[]string) error {
	if len(args) == 0 {
		return errors.New("no url provided")
	}

	found := false
	index := 0
	for _, track := range stateData.AllTracks {
		if track != args[0] {
			stateData.AllTracks[index] = track
			index++
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("Song not in list")
	}
	stateData.AllTracks = stateData.AllTracks[:index]

	index = 0
	for _, track := range stateData.PlayedTracks {
		if track != args[0] {
			stateData.PlayedTracks[index] = track
			index++
		} else {
			found = true
		}
	}
	stateData.PlayedTracks = stateData.PlayedTracks[:index]

	index = 0
	for _, track := range stateData.AvailableTracks {
		if track != args[0] {
			stateData.AvailableTracks[index] = track
			index++
		} else {
			found = true
		}
	}
	stateData.AvailableTracks = stateData.AvailableTracks[:index]

	err := stateData.Save()
	if err != nil {
		return errors.Because(err, nil, "could not save state")
	}

	return nil
}