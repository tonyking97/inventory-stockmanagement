package config

import (
	"encoding/json"
	"log"
	"os"
)

var config *Configuration = nil

func Init() *Configuration{
	if config == nil {
		workingDir, _ := os.Getwd();
		file, err := os.Open( workingDir + "/config/config.json")
		if err != nil {
			log.Fatal("Config File Not Found.")
		}
		defer file.Close()
		decoder := json.NewDecoder(file)
		config = &Configuration{}
		err = decoder.Decode(config)
		if err!= nil {
			log.Fatal(err)
		}
	}
	return config
}
