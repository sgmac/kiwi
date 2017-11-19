package main

import (
	"log"
	"testing"
)

const valid = '\u2713'
const failed = '\u2717'

func TestReadEmptyConfig(t *testing.T) {
	setup()
	config, err := readConfig()

	log.Println("When reading an empty config")
	{
		if err == nil {
			log.Printf("\t %c error should be nil.", valid)
		}

		if config != nil {
			log.Printf("\t %c config should be NOT nil.", valid)
		}

	}
	cleanup()
}

func TestReadConfig(t *testing.T) {
	setup()

	var data = `
[VPS]
    VeID = "1234"
    APIKey = "customApiKey"
`

	writeConfig("testing", []byte(data))
	config, err := readConfig()

	expected := configServ{
		VPS: serverInfo{
			VeID:   "1234",
			APIKey: "mockApiKey",
		},
	}

	log.Println("When reading a config")
	{
		if err == nil {
			log.Printf("\t %c error should be nil.", valid)
		}

		if config.VPS.APIKey != expected.VPS.APIKey {
			log.Printf("\t %c got APIKey %s expected %s\n.", failed, config.VPS.APIKey, expected.VPS.APIKey)
		}

		if config.VPS.VeID == expected.VPS.VeID {
			log.Printf("\t %c got VeID %s expected %s\n.", valid, config.VPS.VeID, expected.VPS.VeID)
		}
	}
	cleanup()
}
