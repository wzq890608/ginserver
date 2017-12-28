package config

import (
	"os"
	"log"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
)

var (
	DEBUG = true
)

type DatabaseConfig struct {
	Type          string
	Host          string
	Port          int
	User          string
	Password      string
	DatabaseName  string
	MaxIdleConns  int
	MaxOpenConns  int
	IsDevelopment bool
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type Web struct {
	Port int
}

type Configuration struct {
	Database DatabaseConfig
	Redis    RedisConfig
	Web      Web
}

func LoadConfig(path string) *Configuration {
	var file []byte
	var err error
	var dir string
	if DEBUG {
		dir = "D:\\GoProject\\src\\ginserver\\"
	} else {
		ex, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		dir = filepath.Dir(ex) + string(os.PathSeparator)
	}
	file, err = ioutil.ReadFile(dir + string(os.PathSeparator) + getEnv() + ".json")
	if err != nil {
		log.Fatalf("Reda config file error %s", err.Error())
	}
	var config *Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error parsing JSON: %s", err.Error())
	}

	return config
}

func getEnv() string {
	if DEBUG {
		return "dev"
	}
	return "prod"
}
