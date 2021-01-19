package api

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const uriInformation = ":8446/smarthome/public/information"

// BoschShcAPI implements interaction with the Bosch Smart Home Controller
type BoschShcAPI interface {
	Information() (i Information, e error)
}

type boschShcAPI struct {
	shcIPAddress string
}

// New creates a new instance of BoschShcAPI with only an IP address
func New(shcIPAddress string) BoschShcAPI {
	b := boschShcAPI{}

	b.shcIPAddress = shcIPAddress

	return &b
}

func (b *boschShcAPI) get(uri string, i interface{}) (e error) {
	url := "https://" + b.shcIPAddress + uri
	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return readErr
	}

	err = json.Unmarshal(body, &i)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	// Disable certificate validation. Ouch.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}
