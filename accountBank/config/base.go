package config

import (
	"log"
	"os"
	"GoLearn/accountBank/helper"
)

type Config struct {
	ServiceName string `json:"service_name"`
	Database struct{
		User 		string 		`json:"user"`
		Password	string		`json:"password"`
		Database 	string		`json:"database"`
		Address		string		`json:"address"`
	}	`json:"database"`
}

func SetupConfig() Config {
	var conf Config

	// Doc file config.local.json
	configFile, err := os.Open("config.local.json")
	if err != nil {
		// Neu khong co file config.local.json thi doc file config.default.json
		configFile, err = os.Open("config.default.json")
		if err != nil {
			panic(err)
		}
		defer  configFile.Close()
	}
	defer  configFile.Close()

	//Parse du lieu file Json va bind vao Controller
	err = helper.DecodeDataFromJsonFile(configFile, &conf)
	if err != nil {
		log.Println("Khong doc duoc file config.")
		panic(err)
	}
	return conf
}

