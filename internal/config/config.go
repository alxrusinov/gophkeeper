package config

import (
	"github.com/spf13/pflag"
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

// Config - type config implements interface app.Config
type Config struct {
	baseURL string
	dbURL   string
}

// Run - parse config files, env and flags - initializes config
func (cfg *Config) Run() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.SetDefault(baseURLField, defaultBaseURL)
	viper.SetDefault(dbURLField, defaultDBURL)

	viper.ReadInConfig()

	pflag.String(baseURLField, defaultBaseURL, "address for running server")
	pflag.String(dbURLField, defaultDBURL, "address for running data base")

	pflag.Parse()

	err = viper.BindPFlags(pflag.CommandLine)

	if err != nil {
		return err
	}

	cfg.baseURL = viper.GetString(baseURLField)
	cfg.dbURL = viper.GetString(dbURLField)

	return err
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
