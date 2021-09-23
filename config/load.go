package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

var (
	global *Config
)

func C() *Config {
	if global == nil {
		panic("Load config first")
	}
	return global
}

func LoadConfigFromToml(filepath string) error {
	cfg := newConfig()
	if _, err := toml.DecodeFile(filepath, cfg); err != nil {
		return err
	}
	return cfg.InitGlobal()
}

func LoadConfigFromEnv() error {
	cfg := newConfig()
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("load config from env, %s", err.Error())
	}
	return cfg.InitGlobal()
}
