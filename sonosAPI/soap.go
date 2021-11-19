package sonosAPI

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const soapEnvNS = "http://schemas.xmlsoap.org/soap/envelope/"

type soapRequest struct {
	XMLName   xml.Name `xml:"s:Envelope"`
	XMLNsSoap string   `xml:"xmlns:s,attr"`
	XMLEncodingStyle string   `xml:"s:encodingStyle,attr"`
	Body      soapBody
}

type soapBody struct {
	XMLName xml.Name `xml:"s:Body"`
	Payload interface{}
}

type soapResponse struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body	*soapResponseBody
}

type soapResponseBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault *soapFault `xml:",omitempty"`
	Content interface{} `xml:",omitempty"`
}

func (s *soapResponseBody) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	ignoreEnd := false

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		if token == nil {
			return nil
		}



		switch elem := token.(type) {
		case xml.StartElement:
			if elem.Name.Space == soapEnvNS && elem.Name.Local == "Fault" {
				fault := soapFault{}
				err := decoder.DecodeElement(&fault, &elem)
				if err != nil {
					return errors.New("decode error")
				}
				s.Fault = &fault
			} else if elem.Name.Space == "urn:schemas-upnp-org:service:RenderingControl:1" && elem.Name.Local == "GetVolumeResponse" {
				content := getVolumeResponse{}
				err := decoder.DecodeElement(&content, &elem)
				if err != nil {
					return errors.New("decode error")
				}
				s.Content = content
			} else if elem.Name.Space == "urn:schemas-upnp-org:service:AVTransport:1" && elem.Name.Local == "PauseResponse" {
				ignoreEnd = true
			} else {
				fmt.Printf("Start: '%s' - '%s'\n", elem.Name.Space, elem.Name.Local)
			}
		case xml.EndElement:
			if elem.Name.Space == soapEnvNS && elem.Name.Local == "Body" {
				return nil
			} else if ignoreEnd {
				ignoreEnd = false
			} else {
				return errors.New("unknown end element")
			}
		}
	}
	return errors.New("OH GOD REACHED THE END")
}

type soapFault struct {
	// XMLName is the serialized name of this object.
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`

	// DetailInternal is a handle to the internal fault detail type. Do not directly access;
	// this is made public only to allow for XML deserialization.
	// Use the Detail() method instead.
	DetailInternal *soapFaultDetail `xml:"detail,omitempty"`
}

type soapFaultDetail struct {
	Content interface{} `xml:",omitempty"`
}

func soapCall(url string, namespace string, action string, payload interface{}) (*soapResponse, error) {
	arequest := soapRequest{
		XMLNsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XMLEncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
		Body: soapBody{
			Payload: payload,
		},
	}

	marshalled, err := xml.MarshalIndent(arequest, "", "\t")
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(marshalled)

	client := http.Client{}
	request, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to construct request: %s", err))
	}

	request.Header.Set("soapaction", fmt.Sprintf("%s#%s", namespace, action))

	response, err := client.Do(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to make request: %s", err))
	}

	if response.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("request failure: %d", response.StatusCode))
	}

	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not read response: %s", err))
	}

	parsedResponse := soapResponse{}

	err = xml.Unmarshal(dataBytes, &parsedResponse)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not parse response: %s", err))
	}

	return &parsedResponse, nil
}
