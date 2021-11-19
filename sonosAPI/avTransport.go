package sonosAPI

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strconv"
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

type seekRequest struct {
	XMLName   xml.Name `xml:"u:Seek"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
	Unit string
	Target string
}

func (device *sonosDevice) DoSeek(seekType SeekType, seekValue int) error {
	request := seekRequest{
		XMLNsSoap: avTransportNamespace,
		InstanceID:  0,
	}

	switch seekType {
	case SeekTrackNumber:
		request.Unit = "TRACK_NR"
		request.Target = strconv.Itoa(seekValue)
		break
	case SeekTime:
		request.Unit = "REL_TIME"
		request.Target = seekValueFormatter(seekValue)
		break
	case SeekTimeDelta:
		request.Unit = "TIME_DELTA"
		if seekValue > 0 {
			request.Target = fmt.Sprintf("+%s", seekValueFormatter(seekValue))
		} else {
			request.Target = fmt.Sprintf("-%s", seekValueFormatter(seekValue))
		}
		break
	}

	data, err := device.deviceRequest(avTransportSuffix, avTransportNamespace, "Seek", request)
	if err != nil {
		return err
	}

	if data.Body.Fault != nil {
		return errors.New("Fault!")
	}

	return nil
}

func seekValueFormatter(seekValue int) string {
	seconds := seekValue % 60
	seekValue -= seconds * 60
	seekValue = seekValue / 60
	minutes := seekValue % 60
	seekValue -= minutes * 60
	hours := minutes / 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}