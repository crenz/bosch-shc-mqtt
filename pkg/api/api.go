package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// BoschShcAPI implements interaction with the Bosch Smart Home Controller
type BoschShcAPI interface {
	Information() (i Information, e error)
	Rooms() (r Rooms, e error)
	Subscribe() (e error)
	Unsubscribe() (e error)
	Poll() (e error)
}

type boschShcAPI struct {
	shcIPAddress string
	client       *http.Client
	pollingID    string
}

// New creates a new instance of BoschShcAPI with only an IP address
func New(shcIPAddress string, clientCertificate string, clientKey string) BoschShcAPI {
	b := boschShcAPI{}

	b.shcIPAddress = shcIPAddress
	certPem := []byte(`-----BEGIN CERTIFICATE-----
MIICtjCCAZ4CCQC0de6xedr42zANBgkqhkiG9w0BAQsFADAdMQswCQYDVQQGEwJERTEOMAwGA1UECgwFd2ViNDIwHhcNMjEwMTE1MjMyMzQ5WhcNNDgwNjAxMjMyMzQ5WjAdMQswCQYDVQQGEwJERTEOMAwGA1UECgwFd2ViNDIwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDU22OIJe+8lzn6dLo8uHS6k4ylp+GJu2KDwMAuXbN5vcRgdvY212qfcyTUDAEXqr+JSX55eqXFEB4PR/WcOKwfPoyCzbpXXKlrSQh9pqcs1k8YLpwenCxIxtqwb8gF3DpBvHsc8+ilP5Pbcha8m6nYOxoEDZLOtxTjg3oBlC1WJJLy4KiKtV7Kj7Qnd7xzSxW2kAxWN2zvFh8dxOqvLN9DdilSIWygVOiGEfPjuJU3WuCyV3JuIoWNl8gPE3p78pqOshHQk2NPgJpgmmTHOLxxhqcAZYn1TrcGIFgEo+WH1FlfTyz3pfnt9P88i2gaHU5ShUeQM1SPFGZmmAgVojbBAgMBAAEwDQYJKoZIhvcNAQELBQADggEBAI4BEk03lNKKjoc5NvkoJWjDhCT2OmekA+Lw9XmAr1eraSvoqu9CS/wwAU7lDpRMdyXpJitTtUcj4ky6N0XJ9x3+2mg9Pw4SDl2FgfNXACvH7IZzETl5lWKWwWTp7mPQ3O/zVebQp0VAFRVK/aFV6Js/lC6jQuREIguOwUQtXgYWJf9q+TvRi+yBSy2G6txkh+ETmoSZ9T69rKvDw88QTBoJyje26kpB0xZrLe/0pJUZXGm9wgNjLJFQ7iEw16okLcHwnjEx10nGuLzyGGWOv9Gtwazw6urIxkbnGxiMA2BgX69M0I4QqcxHjhFq+LgQct4SbgAxud3OCTr5M9cqcm4=
-----END CERTIFICATE-----`)
	keyPem := []byte(`-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDU22OIJe+8lzn6dLo8uHS6k4ylp+GJu2KDwMAuXbN5vcRgdvY212qfcyTUDAEXqr+JSX55eqXFEB4PR/WcOKwfPoyCzbpXXKlrSQh9pqcs1k8YLpwenCxIxtqwb8gF3DpBvHsc8+ilP5Pbcha8m6nYOxoEDZLOtxTjg3oBlC1WJJLy4KiKtV7Kj7Qnd7xzSxW2kAxWN2zvFh8dxOqvLN9DdilSIWygVOiGEfPjuJU3WuCyV3JuIoWNl8gPE3p78pqOshHQk2NPgJpgmmTHOLxxhqcAZYn1TrcGIFgEo+WH1FlfTyz3pfnt9P88i2gaHU5ShUeQM1SPFGZmmAgVojbBAgMBAAECggEAVg8xib1U1VoLLfD9z0kAoOLBDcT3khk59mz7BoQJ6WHJQPs4BupHiJokPLtxwaDeeeJGCVWGKkue66y7Z+Md7C/83XHSMjoboW3QygSUOLOZHPzTUCjyWqZTitxW8c+dmjBbUlRnLlCrNiFFghIptGwI07StM3igMHBa5sKDJf08rlMbGR58884W2/354yXgYzaRzldQe1Zm6ippCV/3QcIh0+Am60QPM77wxBUg7Fup6ogpz6eOOuRp3ODR61ubQ6PQ2Cno8gPDIAgDzTNaPPueyIk3xEcqvmWRTbcMQETCXLKMuPL+zktbjTlF1OXAMD8QtmdVr06vqIWOvw1TwQKBgQDyBmsCmX0+K63MXHJY+lAbSFotZV2E1R7cAgO2lmJlxB/KqJxTiM4xlxd1FrpN751WlLNemDEetr0ooni/a5ysp9KeHUe54z0Z7YAIsp7D2VLRn7Z8tWV+pg0JS6QKBMXRMqhUaymeRpsDLzA1yDTULnlQGn9NfMIiwSIcrIIBSQKBgQDhJc/mzqUUQsKJsv6bfRRB0/I2Hf0Jl2laHdaj7jGHBDet5VWTQdWCTzIaD3SfKnrTDGAHRw5R0j0/2aMea7B/5NpToleQTcZf6+Oq73EjMFBlI4vGWVc0yRLCe9HSJZxlFtZFHpvMfzTd37ArLbD4nAe00fDa+uX6kxoLyj4BuQKBgQCxgWeGlpA2ws5LLhpni96ow93voYJ/Y9eoQIn8pjswrDEs75EH6zRfRpNbuvmVI4Jf99u+Kx0Li7ccUF0C96gHMWbVtF/gw/sSQxA+UNMEjSWivOKYgGoaAytYf/OlrW5wShkPITF69gnGwhs3tsiYPiWXTHfFmxS/bKraIOOQkQKBgQDg+PWLPgV5/1OAYJlFKXCqbmxiYwjLIr6ky5gEaiiXL0Grw7ME2A3OjfAUVklEGiBs7rqtyvSmEZweRwxVg2n0AeufEmLI0M5eXsk5rtSwQsCyrxgl9HPNTPYv26XHhMmHwZANtQ4dAycCZVgVbSye3tpcdkNjrL1M6txl14qqEQKBgDx04ZJM4LDFFB0oulYkOnlAJgIisVrr0FMc3R1sZMpGKNy7gQUmRiQJ9cMH6Vq8ZmXZtj2eRDkW8R1r05Pw9Fr+RiuFCIraax+4fglcvUZwnQ4EHSEf7iX6SmfC3NRlY9aqCVGvvf5nqx8iaodrqJrEGaAz2reVC+YV1CpOUBH8
-----END PRIVATE KEY-----`)
	cert, err := tls.X509KeyPair([]byte(certPem), []byte(keyPem))
	if err != nil {
		log.Fatal(err)
	}
	// Add client cert and disable certificate validation (to enable communication with SHC. Ouch.)
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	b.client = &http.Client{Transport: transport}

	return &b
}

func (b *boschShcAPI) get(uri string, i interface{}) (e error) {
	url := "https://" + b.shcIPAddress + uri
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("api-version", "1.0")
	req.Header.Set("User-Agent", "bosch-shc-mqtt/0.0")
	req.Header.Set("Content-Type", "application/json")
	resp, err := b.client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Errorf("Invalid HTTP Status Code %d", resp.StatusCode)
		return fmt.Errorf("Invalid HTTP Status Code %d", resp.StatusCode)
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
