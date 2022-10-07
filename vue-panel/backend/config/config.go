package config

import (
	"io/ioutil"
	"log"
	"strings"
)

type Config struct {
	Db_pass  string
	RoutName string
}

func GetConfig() (config Config) {
	file, err := ioutil.ReadFile("./config.panel")
	if err != nil {
		log.Fatal("config.GetConfig:", err)
	}

	fileData := string(file)
	if len(fileData) == 0 {
		log.Fatal("config.GetConfig: empty file config.panel")
	}

	configData := strings.Split(fileData, ":")
	if len(configData) < 2 {
		log.Fatal("config.GetConfig: error getting config")
	}

	config = Config{
		Db_pass:  configData[0],
		RoutName: configData[1],
	}

	return
}
