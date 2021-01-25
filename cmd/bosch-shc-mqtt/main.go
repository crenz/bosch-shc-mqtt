package main

import (
	"flag"

	"github.com/crenz/bosch-shc-mqtt/pkg/api"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	//	"github.com/crenz/bosch-shc-mqtt/pkg/api"
)

func main() {
	log.SetLevel(configDefaultLogLevel)
	log.SetOutput(colorable.NewColorableStdout())

	pConfigFile := flag.String("c", "/Users/rnz9fe/Documents/01_Personal/github/bosch-shc-mqtt/test/config.json", "path to configuration file")
	flag.Parse()

	var c *config = &config{}
	var err error
	if len(*pConfigFile) > 0 {
		c, err = configFromFile(*pConfigFile)
		if err != nil {
			log.Errorf("Error reading config file %s: %v", *pConfigFile, err)
			return
		}
	}

	log.SetLevel(c.getLogLevel())
	api := api.New(c.ShcIPAddress, c.getClientCertificateString(), c.getClientKeyString())
	//	x, _ := api.Rooms()
	//	spew.Dump(x)
	if err = api.Subscribe(); err != nil {
		log.Errorf("Error subscribing: %v", err)
	}
	if err = api.Poll(); err != nil {
		log.Errorf("Error polling: %v", err)
	}
	if err = api.Unsubscribe(); err != nil {
		log.Errorf("Error unsubscribing: %v", err)
	}
}
