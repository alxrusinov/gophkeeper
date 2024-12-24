package config

// Config - type config implements interface app.Config
type Config struct {
	baseURL string
	dbURL   string
}

// ParseFlags - parsing flags and set values
func (cfg *Config) parseFlags() (err error) {
	cfg.baseURL = "localhost:8080"
	cfg.dbURL = "mongodb://localhost:3000"
	return err
}

// ParseConfigFile - parsing config file and set values
func (cfg *Config) parseConfigFile() (err error) {
	return err
}

// Run - parse config files, env and flags - initializes config
func (cfg *Config) Run() (err error) {
	err = cfg.parseConfigFile()
	err = cfg.parseFlags()

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
