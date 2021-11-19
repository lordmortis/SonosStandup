package sonosAPI

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type sonosDevice struct {
	baseURL url.URL
}

type PlaybackState uint8

const (
	PlaybackStopped PlaybackState = 0
	PlaybackPaused = 1
	PlaybackPlaying = 2
)

func (state PlaybackState)String() string {
	switch state {
	case PlaybackStopped: return "Stopped"
	case PlaybackPaused: return "Paused"
	case PlaybackPlaying: return "Playing"
	}
	return "Unknown State"
}

type Device interface {
	internalOnly()
	GetVolume() (int, error)
	GetPlaybackState() (*PlaybackState, error)
	DoPause() error
	DoPlay() error
	SetPlaybackURI(URI string) error
}

func NewSonosDevice(addressOrHostname string) (Device, error) {

	baseURL, err := url.Parse(fmt.Sprintf("http://%s:1400", addressOrHostname))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to parse URL: %s", err))
	}

	response, err := http.Get(fmt.Sprintf("%s/xml/device_description.xml", baseURL.String()))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to connect to device: %s", err))
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Unable to query device, received code %d", response.StatusCode))
	}

	device := sonosDevice{
		baseURL: *baseURL,
	}

	return &device, nil
}

func (device *sonosDevice) internalOnly() {}
