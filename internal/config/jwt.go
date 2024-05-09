package config

import (
	"crypto/rsa"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

type JWTConfig struct {
	PrivateKey *rsa.PrivateKey
}

func getJWTConfig() JWTConfig {
	privateKeyPath := fatalGetString("PRIVATE_KEY_PATH")

	pemPrivateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read private key")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pemPrivateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse private key")
	}

	return JWTConfig{
		PrivateKey: privateKey,
	}
}
