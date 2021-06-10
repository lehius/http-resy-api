package config

func Default() *Config {
	return &Config{
		Network: &Network{
			Bind: ":8080",
		},
		Logging: &Logging{
			Level: "warn",
		},
		Store: &Store{
			DatabaseURL: "host=localhost dbname=apiserver_dev password=123456 sslmode=disable",
		},
		configFilePath: "configs/core.yaml",
	}
}
