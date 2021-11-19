package sonosAPI

import (
	"encoding/xml"
	"errors"
	"fmt"
)

var (
	avTransportSuffix = "MediaRenderer/AVTransport/Control"
	avTransportNamespace = "urn:schemas-upnp-org:service:AVTransport:1"
)

type doPauseRequest struct {
	XMLName   xml.Name `xml:"u:Pause"`
	XMLNsSoap string   `xml:"xmlns:u,attr"`
	InstanceID int
}

func (device *sonosDevice) DoPause() error {
	url := fmt.Sprintf("%s/%s", device.baseURL.String(), avTransportSuffix)

	request := doPauseRequest{
		XMLNsSoap: avTransportNamespace,
		InstanceID:  0,
	}

	data, err := soapCall(url, avTransportNamespace, "Pause", request)
	if err != nil {
		return err
	}

	if data.Body.Fault != nil {
		return errors.New("Fault!")
	}

	return nil
}