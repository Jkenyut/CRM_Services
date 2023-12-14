package config

import (
	"fmt"
	"github.com/Jkenyut/libs-numeric-go/libs_config"
	"gopkg.in/yaml.v3"
	"os"
)

var config Config

// please edit in here
type Config struct {
	Listener libs_config.Listener `yaml:"listener,inline"`
	Database struct {
		CRM     libs_config.SQLConfig `yaml:"crm"`
		Timeout int                   `yaml:"timeout" default:"30000"`
	} `yaml:"database"`
	JWT    libs_config.JWTConfig `yaml:"jwt"`
	KeyAES string                `yaml:"keyAES" default:"!a@b3b$n$j6(KQM1"`
	Mock   bool                  `yaml:"mock"`
}

func init() {
	// Read YAML file
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Sprintf("Error reading YAML file: %s\n", err)
	}

	// Unmarshal the YAML file into a Config struct
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Sprintf("Error parsing YAML file: %s\n", err)
	}
}
func GetConfig() *Config {
	return &config
}
