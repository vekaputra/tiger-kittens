package config

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseConfig DatabaseConfig
	EmailConfig    EmailConfig
	Env            string
	IsAllowCORS    bool
	IsEnableDebug  bool
	JWTConfig      JWTConfig
	Port           int
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
		DatabaseConfig: getDatabaseConfig(),
		EmailConfig:    getEmailConfig(),
		Env:            fatalGetString("ENV"),
		IsAllowCORS:    fatalGetBool("ALLOW_CORS"),
		IsEnableDebug:  fatalGetBool("ENABLE_DEBUG"),
		JWTConfig:      getJWTConfig(),
		Port:           fatalGetInt("PORT"),
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

func fatalGetBool(key string) bool {
	fatalCheckKey(key)
	return viper.GetBool(key)
}

func fatalGetDuration(key string, unit time.Duration) time.Duration {
	fatalCheckKey(key)
	return viper.GetDuration(key) * unit
}
