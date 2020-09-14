package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// DBConfig struct type
type DBConfig struct {
	DB_USER     string `json:"dbUser"`
	DB_PASSWORD string `json:"dbPassword"`
	DB_NAME     string `json:"dbName"`
	PORT        string `json:"port"`
	HOST        string `json:"host"`
}

// LoadConfig to load all configs from json file
func LoadConfig() (DBConfig, error) {
	conf := DBConfig{}
	//gopath := os.Getenv("GOPATH")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("current dir", dir)
	confFilePath := dir + `/config/config.json`
	fmt.Println("path", confFilePath)
	confFile, err := os.Open(confFilePath)
	if err != nil {
		log.Fatalf("Unable to find config because of %s", err)
		return conf, err
	}

	err = json.NewDecoder(confFile).Decode(&conf)
	return conf, err
}
