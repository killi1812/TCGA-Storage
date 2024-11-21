package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	MongoConnString string `json:"MongoConnString"`
	MinioConnString string `json:"MinioConnString"`
	AppPort         string `json:"AppPort"`
}

const fileName = "appsettings.json"

var conf Settings = Settings{}

func GetMongoDbConnString() string {
	return conf.MongoConnString
}

func GetMinIOConnString() string {
	return conf.MinioConnString
}

func GetPort() string {
	return conf.AppPort
}

func LoadConfig() error {
	fmt.Println("Loading config")
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error reading config: %s, %s", fileName, err.Error())
	}

	err = json.Unmarshal(bytes, &conf)

	if err != nil {
		return fmt.Errorf("Error Unmarshaling config: %s, %s", fileName, err.Error())
	}
	fmt.Println("Config loaded succesfuly")
	return nil
}
