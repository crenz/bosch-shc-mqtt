package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	MqttBrokerURL string
	MqttClientID  string
	MqttUsername  string
	MqttPassword  string
	ShcIPAddress  string
	Loglevel      string
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
