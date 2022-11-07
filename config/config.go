package config

import "github.com/spf13/viper"

type Config struct {
	DB    Database `mapstructure:"db"`
	Port  string   `mapstructure:"port"`
	Salt  string   `mapstructure:"salt"`
	Token Token    `mapstructure:"token"`
}

type Database struct {
	DatabaseURL    string `mapstructure:"URL"`
	MaxConnections int    `mapstruccture:"maxConnections"`
}

type Token struct {
	Key      string `mapstructure:"Key"`
	LifeTime uint   `mapstructure:"LifeTime"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
