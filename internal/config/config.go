package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"runtime/debug"
)

type Config struct {
	Env            string
	DatabaseConfig DatabaseConfig
}

func Load() *Config {
	viper.SetConfigName(fmt.Sprintf("env.%v", os.Getenv("ENV")))
	viper.SetConfigType("env")
	viper.AddConfigPath("./env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Viper: failed to find config file")
	}

	return &Config{
		Env:            fatalGetString("ENV"),
		DatabaseConfig: getDatabaseConfig(),
	}
}

func fatalCheckKey(key string) {
	if !viper.IsSet(key) {
		debug.PrintStack()
		log.Fatal().Str("key", key).Msg("can't find key")
	}
}

func fatalGetString(key string) string {
	fatalCheckKey(key)
	return viper.GetString(key)
}

func fatalGetInt(key string) int {
	fatalCheckKey(key)
	return viper.GetInt(key)
}
