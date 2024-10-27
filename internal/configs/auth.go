package configs

import "time"

type hash struct {
	Cost int `yaml:"cost"`
}
type token struct {
	ExpiresIn int    `yaml:"expires_in"`
	SignedKey string `yaml:"signed_key"`
}

func (t *token) GetExpiresInDuration() time.Duration {
	return time.Duration(t.ExpiresIn) * time.Hour
}

type Auth struct {
	Hash  hash  `yaml:"hash"`
	Token token `yaml:"token"`
}
