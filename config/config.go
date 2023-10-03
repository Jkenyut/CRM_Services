package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
)

type SQLConfig struct {
	Enable        bool   `yaml:"enable" default:"false" desc:"config:sql:enable" json:"Enable,omitempty"`
	Driver        string `yaml:"driver" default:"" desc:"config:sql:driver" json:"Driver,omitempty"`
	Host          string `yaml:"host" default:"127.0.0.1" desc:"config:sql:host" json:"Host,omitempty"`
	Port          int    `yaml:"port" default:"3306" desc:"config:sql:port" json:"Port,omitempty"`
	Username      string `yaml:"username" default:"root"  desc:"config:sql:username" json:"Username,omitempty"`
	Password      string `yaml:"password" default:"root" desc:"config:sql:password" json:"Password,omitempty"`
	Database      string `yaml:"database" default:"mydb" desc:"config:sql:database" json:"Database,omitempty"`
	Options       string `yaml:"options" default:"" desc:"config:sql:options" json:"Options,omitempty"`
	Connection    string `yaml:"connection" default:"" desc:"config:sql:connection" json:"Connection,omitempty"`
	AutoReconnect bool   `yaml:"autoreconnect" default:"false"  desc:"config:sql:autoreconnect" json:"AutoReconnect,omitempty"`
	StartInterval int    `yaml:"startinterval" default:"2"  desc:"config:sql:startinterval" json:"StartInterval,omitempty"`
	MaxError      int    `yaml:"maxerror" default:"5"  desc:"config:sql:maxerror" json:"MaxError,omitempty"`
	CustomPool    bool   `yaml:"customPool" default:"5"  desc:"config:sql:customPool" json:"CustomPool,omitempty"`
	MaxConn       int    `yaml:"maxConn" default:"5"  desc:"config:sql:maxConn" json:"MaxConn,omitempty"`
	MaxIdle       int    `yaml:"maxIdle" default:"5"  desc:"config:sql:maxIdle" json:"MaxIdle,omitempty"`
	LifeTime      int    `yaml:"lifeTime" default:"5"  desc:"config:sql:lifeTime" json:"LifeTime,omitempty"`
}

type ConfigYaml struct {
	Database SQLConfig `yaml:"database" default:"database"`
	Redis    string    `yaml:"Key" default:"Feedback"`
}

func main() {
	// Get the default value for the Redis field.
	//config := ConfigYaml{}
	fmt.Println(reflect.ValueOf(ConfigYaml{}))
	// Marshal the struct to YAML
	config := "haha"
	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Println("Error marshaling YAML:", err)
		return
	}
	f, err := os.Create("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(string(yamlData))
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("YAML data has been written")
}
