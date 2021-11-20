package sonosAPI

import (
	"encoding/xml"
	"errors"
)

var (
	renderingSuffix = "MediaRenderer/RenderingControl/Control"
	renderingSoapNamespace = "urn:schemas-upnp-org:service:RenderingControl:1"
)

type getVolumeRequest struct {
	XMLName   xml.Name `xml:"u:GetVolume"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
	Channel string
}

type getVolumeResponse struct {
	CurrentVolume int
}

func (device *sonosDevice) GetVolume() (int, error) {
	request := getVolumeRequest{
		XMLNsSoap: renderingSoapNamespace,
		InstanceID:  0,
		Channel:   "Master",
	}

	data, err := device.deviceRequest(renderingSuffix, renderingSoapNamespace, "GetVolume", request)
	if err != nil {
		return -1, err
	}

	if data.Body.Fault != nil {
		return -1, errors.New("Fault!")
	}

	response, ok := data.Body.Content.(getVolumeResponse)
	if !ok {
		return -1, errors.New("Invalid reply from server")
	}

	return response.CurrentVolume, nil
}

type setVolumeRequest struct {
	XMLName   xml.Name `xml:"u:SetVolume"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
	Channel string
	DesiredVolume int
}

func (device *sonosDevice) SetVolume(volume int) error {
	request := setVolumeRequest{
		XMLNsSoap: renderingSoapNamespace,
		InstanceID:  0,
		Channel:   "Master",
		DesiredVolume: volume,
	}

	data, err := device.deviceRequest(renderingSuffix, renderingSoapNamespace, "SetVolume", request)
	if err != nil {
		return err
	}

	if data.Body.Fault != nil {
		return errors.New("fault")
	}

	return nil
}