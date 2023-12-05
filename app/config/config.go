package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var config Config

type SQLConfig struct {
	Enable        bool   `yaml:"enable" default:"false" desc:"config:sql:enable"`
	Driver        string `yaml:"driver" default:"" desc:"config:sql:driver"`
	Host          string `yaml:"host" default:"127.0.0.1" desc:"config:sql:host"`
	Port          int    `yaml:"port" default:"3306" desc:"config:sql:port"`
	Username      string `yaml:"username" default:"root"  desc:"config:sql:username"`
	Password      string `yaml:"password" default:"" desc:"config:sql:password"`
	Database      string `yaml:"database" default:"crm_bootcamp" desc:"config:sql:database"`
	Options       string `yaml:"options" default:"" desc:"config:sql:options"`
	Connection    string `yaml:"connection" default:"" desc:"config:sql:connection"`
	AutoReconnect bool   `yaml:"autoreconnect" default:"false"  desc:"config:sql:autoreconnect"`
	StartInterval int    `yaml:"startinterval" default:"2"  desc:"config:sql:startinterval"`
	MaxError      int    `yaml:"maxerror" default:"5"  desc:"config:sql:maxerror"`
	CustomPool    bool   `yaml:"customPool" default:"5"  desc:"config:sql:customPool"`
	MaxConn       int    `yaml:"maxConn" default:"5"  desc:"config:sql:maxConn"`
	MaxIdle       int    `yaml:"maxIdle" default:"5"  desc:"config:sql:maxIdle"`
	LifeTime      int    `yaml:"lifeTime" default:"5"  desc:"config:sql:lifeTime"`
}

// please edit in here
type Config struct {
	Database struct {
		CRM     SQLConfig `yaml:"CRM"`
		Timeout int       `yaml:"TIMEOUT" default:"30000"`
	} `yaml:"database"`
	JWT struct {
		JwtAccess         string `yaml:"JWTACCESS" default:"random"`
		JwtRefresh        string `yaml:"JWTREFRESH" default:"random"`
		ExpiredJWT        int    `yaml:"EXPIREDJWT" default:"1"`
		ExpiredRefreshJWT int    `yaml:"EXPIREDREFRESHJWT" default:"24"`
	}
	Mock bool `yaml:"MOCK"`
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
