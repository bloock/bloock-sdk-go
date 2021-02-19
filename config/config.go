package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Constants struct {
	Env   Env
	Cloud Cloud
}

type Env string

type Cloud struct {
	Azure Azure
}

type Azure struct {
	Test     bool
	Host     string
	PathProd string
	PathTest string
}

func EnvVariables() Constants {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var configuration Constants
	if err := viper.Unmarshal(&configuration); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return configuration
}
