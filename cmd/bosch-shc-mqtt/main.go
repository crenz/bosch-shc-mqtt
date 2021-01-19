package main

import (
	"flag"

	"github.com/crenz/bosch-shc-mqtt/pkg/api"
	"github.com/davecgh/go-spew/spew"
	"github.com/mattn/go-colorable"
	log "github.com/sirupsen/logrus"
	//	"github.com/crenz/bosch-shc-mqtt/pkg/api"
)

const defaultLogLevel = log.ErrorLevel

func getLogLevel(llString string) log.Level {
	level, error := log.ParseLevel(llString)
	if error != nil {
		log.Errorf("Unknown log level %s, using error", llString)
		return defaultLogLevel
	}
	return level
}

func main() {
	log.SetLevel(defaultLogLevel)
	log.SetOutput(colorable.NewColorableStdout())

	pConfigFile := flag.String("c", "", "path to configuration file")
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

	if c.Loglevel != "" {
		log.SetLevel(getLogLevel(c.Loglevel))
	}

	api := api.New(c.ShcIPAddress)
	spew.Dump(api.Information())
}
