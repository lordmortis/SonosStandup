package sonosAPI

import (
	"encoding/xml"
	"errors"
)

var (
	avTransportSuffix = "MediaRenderer/AVTransport/Control"
	avTransportNamespace = "urn:schemas-upnp-org:service:AVTransport:1"
)

type pauseRequest struct {
	XMLName   xml.Name `xml:"u:Pause"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
}

func (device *sonosDevice) DoPause() error {
	request := pauseRequest{
		XMLNsSoap: avTransportNamespace,
		InstanceID:  0,
	}

	data, err := device.deviceRequest(avTransportSuffix, avTransportNamespace, "Pause", request)
	if err != nil {
		return err
	}

	if data.Body.Fault != nil {
		return errors.New("Fault!")
	}

	return nil
}

type playRequest struct {
	XMLName   xml.Name `xml:"u:Play"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
	Speed int
}

func (device *sonosDevice) DoPlay() error {
	request := playRequest{
		XMLNsSoap: avTransportNamespace,
		InstanceID:  0,
		Speed: 1,
	}

	data, err := device.deviceRequest(avTransportSuffix, avTransportNamespace, "Play", request)
	if err != nil {
		return err
	}

	if data.Body.Fault != nil {
		return errors.New("Fault!")
	}

	return nil
}

type getPlaybackStateRequest struct {
	XMLName   xml.Name `xml:"u:GetTransportInfo"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
}

type getPlaybackStateResponse struct {
	CurrentTransportState string
}

func (device *sonosDevice) GetPlaybackState() (*PlaybackState, error) {
	request := getPlaybackStateRequest{
		XMLNsSoap: avTransportNamespace,
		InstanceID:  0,
	}

	data, err := device.deviceRequest(avTransportSuffix, avTransportNamespace, "GetTransportInfo", request)
	if err != nil {
		return nil, err
	}

	if data.Body.Fault != nil {
		return nil, errors.New("Fault!")
	}

	response, ok := data.Body.Content.(getPlaybackStateResponse)
	if !ok {
		return nil, errors.New("Invalid reply from server")
	}

	var state PlaybackState

	switch response.CurrentTransportState {
	case "PAUSED_PLAYBACK":
		state = PlaybackPaused
		break
	case "PLAYING":
		state = PlaybackPlaying
		break
	case "STOPPED":
		state = PlaybackStopped
		break
	}

	return &state, nil
}

type setPlaybackURIRequest struct {
	XMLName   xml.Name `xml:"u:SetAVTransportURI"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
	CurrentURI string
	CurrentURIMetaData string
}

func (device *sonosDevice) SetPlaybackURI(URI string) error {
	request := setPlaybackURIRequest{
		XMLNsSoap: avTransportNamespace,
		InstanceID:  0,
		CurrentURI: URI,
		CurrentURIMetaData: "",
	}

	data, err := device.deviceRequest(avTransportSuffix, avTransportNamespace, "SetAVTransportURI", request)
	if err != nil {
		return err
	}

	if data.Body.Fault != nil {
		return errors.New("Fault!")
	}

	return nil
}