package main

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type config struct {
	MqttBrokerURL         string
	MqttClientID          string
	MqttUsername          string
	MqttPassword          string
	ShcIPAddress          string
	Loglevel              string
	ClientCertificateFile string
	ClientKeyFile         string
	ClientCertificate     string
	ClientKey             string
}

func configFromFile(path string) (c *config, e error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}
	var ctemp config
	err = json.Unmarshal(data, &ctemp)
	if err != nil {
		return nil, err
	}
	return &ctemp, nil
}

const configDefaultLogLevel = log.ErrorLevel

func (c *config) getLogLevel() log.Level {
	if c.Loglevel == "" {
		return configDefaultLogLevel
	}

	level, error := log.ParseLevel(c.Loglevel)
	if error != nil {
		log.Errorf("Unknown log level %s, using error", c.Loglevel)
		return configDefaultLogLevel
	}
	return level
}

func (c *config) getClientCertificateString() (cert string) {
	return c.ClientCertificate
}

func (c *config) getClientKeyString() (key string) {
	return c.ClientKey
}
