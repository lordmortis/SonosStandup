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
		request.Target = timeValueFormatter(seekValue)
		break
	case SeekTimeDelta:
		request.Unit = "TIME_DELTA"
		if seekValue > 0 {
			request.Target = fmt.Sprintf("+%s", timeValueFormatter(seekValue))
		} else {
			request.Target = fmt.Sprintf("-%s", timeValueFormatter(seekValue))
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

type getPositionInfoRequest struct {
	XMLName   xml.Name `xml:"u:GetPositionInfo"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
}

type getPositionInfoResponse struct {
	Track int
	TrackDuration string
	TrackMetadata string
	RelTime string
}

func (device *sonosDevice) GetPositionInfo() (*PlaybackPosition, error) {
	request := getPositionInfoRequest{
		XMLNsSoap: avTransportNamespace,
		InstanceID:  0,
	}

	data, err := device.deviceRequest(avTransportSuffix, avTransportNamespace, "GetPositionInfo", request)
	if err != nil {
		return nil, err
	}

	if data.Body.Fault != nil {
		return nil, errors.New("Fault!")
	}

	response, ok := data.Body.Content.(getPositionInfoResponse)
	if !ok {
		return nil, errors.New("Invalid reply from server")
	}

	playbackPosition := PlaybackPosition{
		TrackNo: response.Track,
		TrackDuration: timeValueParser(response.TrackDuration),
		TrackPosition: timeValueParser(response.RelTime),
	}

	return &playbackPosition, nil
}

type getMediaInfoRequest struct {
	XMLName   xml.Name `xml:"u:GetMediaInfo"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
}

type getMediaInfoResponse struct {
	NrTracks int
	CurrentURI string
	NextURI string
}

func (device *sonosDevice) GetMediaInfo() (*MediaInfo, error) {
	request := getMediaInfoRequest{
		XMLNsSoap: avTransportNamespace,
		InstanceID:  0,
	}

	data, err := device.deviceRequest(avTransportSuffix, avTransportNamespace, "GetMediaInfo", request)
	if err != nil {
		return nil, err
	}

	if data.Body.Fault != nil {
		return nil, errors.New("Fault!")
	}

	response, ok := data.Body.Content.(getMediaInfoResponse)
	if !ok {
		return nil, errors.New("Invalid reply from server")
	}

	mediaInfo := MediaInfo{
		NumberOfTracks: response.NrTracks,
		CurrentURI: response.CurrentURI,
		NextURI: response.NextURI,
	}

	return &mediaInfo, nil
}
