package config

import "github.com/spf13/viper"

var Config Settings

type Settings struct {
	ApiKey string
	ApiURL string
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}
}
