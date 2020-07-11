package store

type Config struct {
	Hostname string `toml:"host_name"`
	HostPort int64 `toml:"host_port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	DatabaseName string `toml:"database_name"`
}

func NewConfig() *Config {
	return &Config{}
}