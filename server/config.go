package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"os"
)

type (
	Config struct {
		Port int
	}
)

func LoadConfig(r io.Reader) (*Config, error) {
	var config Config
	meta, err := toml.DecodeReader(r, &config)
	if err != nil {
		return nil, err
	}

	if !meta.IsDefined("port") {
		return nil, fmt.Errorf("undefined port in config")
	}

	return &config, nil
}

func LoadConfigFile(filepath string) (*Config, error) {
	if _, err := os.Stat(filepath); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath, os.O_RDONLY, 666)
	if err != nil {
		return nil, err
	}

	config, err := LoadConfig(file)
	if err != nil {
		return nil, err
	}

	return config, nil
}
