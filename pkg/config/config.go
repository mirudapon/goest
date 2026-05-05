package config

type Config struct {

}

var cfg *Config

func Load() *Config {
	cfg = &Config{}
	loadEnv(".env")

	return cfg
}

func Get() *Config {
	return cfg
}

