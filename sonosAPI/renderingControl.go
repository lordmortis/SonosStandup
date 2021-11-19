package sonosAPI

import (
	"encoding/xml"
	"errors"
	"fmt"
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
	url := fmt.Sprintf("%s/%s", device.baseURL.String(), renderingSuffix)

	request := getVolumeRequest{
		XMLNsSoap: renderingSoapNamespace,
		InstanceID:  0,
		Channel:   "Master",
	}

	data, err := soapCall(url, renderingSoapNamespace, "GetVolume", request)
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