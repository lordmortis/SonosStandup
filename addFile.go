package main

import (
	"gopkg.in/errgo.v2/errors"
	"net/url"
)

type AddFileCommand struct {
}

var addFileCommand AddFileCommand

func init () {
	_, err := parser.AddCommand(
		"addFile",
		"Add File",
		"Add Audio file to playback",
		&addFileCommand)
	if err != nil {
		panic(err)
	}
}

func (x *AddFileCommand)Execute(args[]string) error {
	if len(args) == 0 {
		return errors.New("no url provided")
	}

	_, err := url.Parse(args[0])
	if err != nil {
		return errors.Because(err, nil, "invalid song URL")
	}

	for _, item := range stateData.AllTracks {
		if item == args[0] {
			return errors.New("Song already in list")
		}
	}

	stateData.AllTracks = append(stateData.AllTracks, args[0])
	stateData.AvailableTracks = append(stateData.AllTracks, args[0])
	err = stateData.Save()
	if err != nil {
		return errors.Because(err, nil, "unable to save state data")
	}

	return nil
}