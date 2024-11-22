package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	MongoConnString string      `json:"MongoConnString"`
	MinioConn       MinioConfig `json:"MinioConn"`
	AppPort         string      `json:"AppPort"`
}

type MinioConfig struct {
	Endpoint        string `json:Endpoint`
	AccessKeyID     string `json:AccessKeyID`
	SecretAccessKey string `json:SecretAccessKey`
	UseSSL          bool   `json:UseSSL`
}

const fileName = "appsettings.json"

var Conf Settings = Settings{}

func LoadConfig() error {
	fmt.Println("Loading config")
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("Error reading config: %s \n%s", fileName, err.Error())
	}

	err = json.Unmarshal(bytes, &Conf)

	if err != nil {
		return fmt.Errorf("Error Unmarshaling config: %s \n%s", fileName, err.Error())
	}
	fmt.Println("Config loaded succesfuly")
	return nil
}
