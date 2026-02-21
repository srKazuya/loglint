package config

import (
	"log"

	"github.com/golangci/plugin-module-register/register"
	"github.com/spf13/viper"
)

func Load(path string) {
	v := viper.New()

	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		log.Printf("config not found, using defaults: %v", err)
		return
	}

	if err := v.Unmarshal(&Global); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
}


func LoadFromAny(settings any) error {
	cfg, err := register.DecodeSettings[Config](settings)
	if err != nil {
		return err
	}

	Global = cfg
	return nil
}