package configs

import (
	"github.com/spf13/viper"
)

type Conf struct {
	key string `mapstructure:"KEY"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf
	viper.SetConfigName("go_climate")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(path + "/.env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, err
}
