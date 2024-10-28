package configs

import (
	"fmt"
	"os"

	"github.com/ChatService/configs"
	"gopkg.in/yaml.v3"
)

type ConfigFilePath string

type Config struct {
	Database Database `yaml:"database"`
	Auth     Auth     `yaml:"auth"`
	Http     Http     `yaml:"http"`
}

func NewConfig(filePath ConfigFilePath) (Config, error) {
	var (
		configBytes = configs.DefaultConfigBytes
		config      = Config{}
		err         error
	)

	if filePath != "" {
		configBytes, err = os.ReadFile(string(filePath))
		if err != nil {
			return Config{}, fmt.Errorf("failed to read YAMl file: %v", err)
		}
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal YAML: %v", err)
	}
	return config, nil
}
