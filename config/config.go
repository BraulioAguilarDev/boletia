package config

import "github.com/spf13/viper"

var Config Settings

type Database struct {
	DBhost     string
	DBport     int
	DBuser     string
	DBname     string
	DBpassword string
	SSLmode    string
	Driver     string
}

type Settings struct {
	Period  int
	Timeout int
	Port    int

	ApiKey string
	ApiURL string

	Database Database
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
