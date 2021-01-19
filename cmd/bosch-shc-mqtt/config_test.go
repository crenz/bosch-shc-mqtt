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
	var c *config
	var err error

	// Reading in empty file should fail
	c, err = configFromFile("")
	if c != nil || err == nil {
		t.Errorf("configFromFile(\"\") = (%q, %q), want nil", c, err)
	}

	// Reading in non-JSON file should fail
	c, err = configFromFile("config_test.go")
	if c != nil || err == nil {
		t.Errorf("configFromFile(\"\") = (%q, %q), want nil", c, err)
	}

	// Reading in test configuration should pass
	path := "../../test/config.json"

	testConfig := config{
		MqttBrokerURL: "tcp://localhost:8883",
		MqttClientID:  "bosch-shc-mqtt",
		MqttUsername:  "user",
		MqttPassword:  "password",
		ShcIPAddress:  "192.168.0.10",
		Loglevel:      "info",
	}

	c, err = configFromFile(path)
	if err != nil {
		wd, _ := os.Getwd()
		t.Errorf("Failed to read config file %s: %v (working dir: %s)", path, err, wd)
	} else {
		compareTestData(t, "MqttBroker", c.MqttBrokerURL, testConfig.MqttBrokerURL)
		compareTestData(t, "MqttClientID", c.MqttClientID, testConfig.MqttClientID)
		compareTestData(t, "MqttUsername", c.MqttUsername, testConfig.MqttUsername)
		compareTestData(t, "MqttPassword", c.MqttPassword, testConfig.MqttPassword)
		compareTestData(t, "ShcIPAddress", c.ShcIPAddress, testConfig.ShcIPAddress)
		compareTestData(t, "Loglevel", c.Loglevel, testConfig.Loglevel)
	}
}
