package model

// Reusable config goes here
type AppConfig struct {
	Encryption Encryption
	Jwt        JwtConfig
	Auth       AuthConfig
}

type Encryption struct {
	Cost           uint8
	MasterPassword string `mapstructure:"mp"`
}

type AuthConfig struct {
	ExcludedMethods []string
}
