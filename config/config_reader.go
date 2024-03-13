package config

import "github.com/pelletier/go-toml"

func ReadConfig(filePath string) (*Config, error) {
	var config Config

	file, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := file.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
