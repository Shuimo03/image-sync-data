package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type ProxyConfig struct {
	HTTPProxy  string `toml:"http_proxy"`
	HTTPSProxy string `toml:"https_proxy"`
}

type TargetConfig struct {
	Server   string `toml:"server"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type Config struct {
	Proxy  ProxyConfig  `toml:"proxy"`
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
