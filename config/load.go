package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type ProxyConfig struct {
	HTTPProxy  string `toml:"http_proxy"`
	HTTPSProxy string `toml:"https_proxy"`
}

type SourceConfig struct {
}

type TargetConfig struct {
}

type Config struct {
	Proxy  ProxyConfig  `toml:"proxy"`
	Source SourceConfig `toml:"source"`
	Target TargetConfig `toml:"target"`
}

func LoadConfig(path string) (*Config, error) {
	configFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if tomlError := toml.Unmarshal(configFile, &config); tomlError != nil {
		return nil, tomlError
	}
	return &config, nil
}
