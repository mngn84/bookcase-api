package apiserver

// Config
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: `toml:"bind_addr"`,
		LogLevel: `toml:"log_level"`,
	}
}
