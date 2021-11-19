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

type SonosDevice interface {
	internalOnly()
	GetVolume() (int, error)
	DoPause() error
	DoPlay() error
}

func NewSonosDevice(addressOrHostname string) (SonosDevice, error) {

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

func (device *sonosDevice) DoPlay() error {
	return errors.New("NOT IMPLEMENTED")
}