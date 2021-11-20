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
	PlaybackStopped PlaybackState = iota
	PlaybackPaused
	PlaybackPlaying
)

func (state PlaybackState)String() string {
	switch state {
	case PlaybackStopped: return "Stopped"
	case PlaybackPaused: return "Paused"
	case PlaybackPlaying: return "Playing"
	}
	return "Unknown State"
}

type SeekType uint8

const (
	SeekTrackNumber SeekType = iota
	SeekTime
	SeekTimeDelta
)

func (seekType SeekType)String() string {
	switch seekType {
	case SeekTrackNumber: return "Seek to Track Number"
	case SeekTime: return "Seek to Time"
	case SeekTimeDelta: return "Seek to Delta Time"
	}
	return "Unknown Seek Type"
}

type Device interface {
	internalOnly()
	GetVolume() (int, error)
	SetVolume(int) error
	GetPlaybackState() (*PlaybackState, error)
	DoPause() error
	DoPlay() error
	SetPlaybackURI(URI string) error
	DoSeek(seekType SeekType, seekValue int) error
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
