package config

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

const (
	// defaultBaseURL - default value of base url for running server
	defaultBaseURL string = "0.0.0.0:8000"
	// defaultDBURL - default value of address for runnimg data base
	defaultDBURL string = "mongodb://mongo:27017/?connect=direct"
)

const (
	// dbURLField - configuration fields for data base
	dbURLField string = "dbURL"
	// baseURLField - cinfiguration field for base url of application
	baseURLField string = "baseURL"
)

// Define some constants for our secret keys and token expiration durations.
const (
	accessSecret         = "signature_hmac_secret_shared_key"
	refreshSecret        = "signature_hmac_secret_shared_keyh"
	accessExpire         = 24 * time.Hour
	refreshExpire        = 24 * 7 * time.Hour
	accessEncriptionKey  = "signature_hmac_secret_shared_key"
	refreshEncriptionKey = "signature_hmac_secret_shared_key"
	sigKey               = "signature_hmac_secret_shared_key"
)

// auth - struct configurations auth
type Auth struct {
	AccessSecret         string        `mapstructure:"accessSecret"`
	RefreshSecret        string        `mapstructure:"refreshSecret"`
	AccessExpire         time.Duration `mapstructure:"accessExpire"`
	RefreshExpire        time.Duration `mapstructure:"refreshExpire"`
	AccessEncriptionKey  string        `mapstructure:"accessEncriptionKey"`
	RefreshEncriptionKey string        `mapstructure:"refreshEncriptionKey"`
	SigKey               string        `mapstructure:"sigKey"`
}

// Config - type config implements interface app.Config
type Config struct {
	baseURL string `mapstructure:"baseURL"`
	dbURL   string `mapstructure:"dbURL"`
	Auth    Auth
}

func errHandler(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// Run - parse config files, env and flags - initializes config
func (cfg *Config) Run() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.SetDefault(baseURLField, defaultBaseURL)
	viper.SetDefault(dbURLField, defaultDBURL)

	viper.SetDefault("auth.accessSecret", accessSecret)
	viper.SetDefault("auth.refreshSecret", refreshSecret)
	viper.SetDefault("auth.accessExpire", accessExpire)
	viper.SetDefault("auth.refreshExpire", refreshExpire)
	viper.SetDefault("auth.accessEncriptionKey", accessEncriptionKey)
	viper.SetDefault("auth.refreshEncriptionKey", accessSecret)
	viper.SetDefault("auth.sigKey", sigKey)

	var result map[string]interface{}

	viper.ReadInConfig()

	err = viper.Unmarshal(&result)

	errHandler(err)

	err = mapstructure.Decode(result, cfg)

	errHandler(err)

	cfg.baseURL = viper.GetString(baseURLField)
	cfg.dbURL = viper.GetString(dbURLField)

	return err
}

// RunMock - method sets mocked configuration
func (cfg *Config) RunMock() {
	cfg.Auth = Auth{
		AccessSecret:         "signature_hmac_secret_shared_key",
		RefreshSecret:        "signature_hmac_secret_shared_keyh",
		AccessExpire:         24 * time.Hour,
		RefreshExpire:        24 * 7 * time.Hour,
		AccessEncriptionKey:  "signature_hmac_secret_shared_key",
		RefreshEncriptionKey: "signature_hmac_secret_shared_key",
		SigKey:               "signature_hmac_secret_shared_key",
	}
	cfg.baseURL = "0.0.0.0:8000"
	cfg.dbURL = "mongodb://mongo:27017/?connect=direct"
}

// GetBaseURL - return base url
func (cfg *Config) GetBaseURL() string {
	return cfg.baseURL
}

// GetDbURL - return database url
func (cfg *Config) GetDbURL() string {
	return cfg.dbURL
}

// NewConfig - create new config
func NewConfig() *Config {
	cfg := &Config{}

	return cfg
}
