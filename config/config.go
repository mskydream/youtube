package config

import "github.com/spf13/viper"

type Config struct {
	DB   Database `mapstructure:"db"`
	Port string   `mapstructure:"port"`
}

type Database struct {
	DatabaseURL    string `mapstructure:"URL"`
	MaxConnections int    `mapstruccture:"maxConnections"`
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
