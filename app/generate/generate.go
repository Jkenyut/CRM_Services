package main

import (
	"crm_service/app/config"
	"fmt"
	"github.com/mcuadros/go-defaults"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	c := new(config.Config)
	defaults.SetDefaults(c)
	d, _ := yaml.Marshal(&c)
	// To write to a file
	err := os.WriteFile("config.yaml", d, 0644)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	os.Exit(0)
}
