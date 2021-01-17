package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Broker   string
	ClientID string
	Username string
	Password string
	Loglevel string
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
