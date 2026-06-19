package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Config struct {
	Port int    `json:"port"`
	Host string `json:"host"`
}

func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func UnmarshalJson() {
	cfg, err := loadConfig("./task10/file_config.json")
	if err != nil{
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	fmt.Println("Port: ", cfg.Port, " Host: ",cfg.Host)
}
