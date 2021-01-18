package main

import (
	"os"
	"testing"
)

func compareTestData(t *testing.T, field string, got string, want string) {
	if got != want {
		t.Errorf("Config file differs from test data. Read %q %q, expected %q", field, got, want)
	}

}

func TestConfigFromFile(t *testing.T) {
	var c *Config
	var err error

	c, err = configFromFile("")
	if c != nil || err == nil {
		t.Errorf("configFromFile(\"\") = (%q, %q), want nil", c, err)
	}

	path := "../../test/config.json"

	testConfig := Config{
		MqttBrokerUrl: "tcp://localhost:8883",
		MqttClientID:  "bosch-shc-mqtt",
		MqttUsername:  "user",
		MqttPassword:  "password",
		Loglevel:      "info",
	}

	c, err = configFromFile(path)
	if err != nil {
		wd, _ := os.Getwd()
		t.Errorf("Failed to read config file %s: %v (working dir: %s)", path, err, wd)
	} else {
		compareTestData(t, "MqttBroker", c.MqttBrokerUrl, testConfig.MqttBrokerUrl)
		compareTestData(t, "MqttClientID", c.MqttClientID, testConfig.MqttClientID)
		compareTestData(t, "MqttUsername", c.MqttUsername, testConfig.MqttUsername)
		compareTestData(t, "MqttPassword", c.MqttPassword, testConfig.MqttPassword)
		compareTestData(t, "Loglevel", c.Loglevel, testConfig.Loglevel)
	}
}
