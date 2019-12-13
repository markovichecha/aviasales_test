package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config represents the .yml file context
type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		DBName   string `yaml:"dbname"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

// NewConfig creates a Config from the .yml file
func NewConfig(file string) Config {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
