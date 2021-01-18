package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	MqttBrokerUrl string
	MqttClientID  string
	MqttUsername  string
	MqttPassword  string
	Loglevel      string
}

func configFromFile(path string) (config *Config, e error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}
	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
