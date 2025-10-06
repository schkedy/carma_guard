package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER      string `mapstructure:"DB_DRIVER"`
	DB_SOURCE      string `mapstructure:"DB_SOURCE"`
	SERVER_ADDRESS string `mapstructure:"SERVER_ADDRESS"`
	BOT_TOKEN      string `mapstructure:"BOT_TOKEN"`
}

// LoadConfig считывает настройки из файла или переменных окружения.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
