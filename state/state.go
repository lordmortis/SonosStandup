package state

import (
	"SonosStandup/sonosAPI"
	"encoding/gob"
	"gopkg.in/errgo.v2/errors"
	"os"
)

type Data struct {
	AllTracks []string
	PlayedTracks []string
	AvailableTracks []string
	LastTrack string

	PreviousState sonosAPI.PlaybackState
	PreviousQueue int
	PreviousVolume int
	path string
}

func New(path string) Data {
	return Data{
		path: path,
	}
}

func (data *Data)Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Because(err, nil, "Could not open state data file")
	}

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		return errors.Because(err, nil, "Could not parse state data")
	}
	data.path = path

	return nil
}

func (data *Data)Save() error {
	file, err := os.Create(data.path)
	if err != nil {
		return errors.Because(err, nil, "Could not create state data file")
	}

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.Because(err, nil, "Could not write state data")
	}

	return nil
}