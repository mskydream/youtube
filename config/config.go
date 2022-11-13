package config

import "github.com/spf13/viper"

type Config struct {
	DB       Database `mapstructure:"db"`
	Port     string   `mapstructure:"port"`
	Salt     string   `mapstructure:"salt"`
	Token    Token    `mapstructure:"token"`
	Telegram Telegram `mapstructure:"telegram"`
}

type Database struct {
	DatabaseURL    string `mapstructure:"URL"`
	MaxConnections int    `mapstruccture:"maxConnections"`
}

type Token struct {
	Key      string `mapstructure:"Key"`
	LifeTime uint   `mapstructure:"LifeTime"`
}

type Telegram struct {
	Token  string `mapstructure:"token"`
	ChatId int64  `mapstructure:"chatid"`
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
