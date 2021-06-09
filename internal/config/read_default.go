package config

func Default() *Config {
	return &Config{
		Network: &Network{
			Bind: ":8080",
		},
		Logging: &Logging{
			Level: "warn",
		},
		configFilePath: "configs/core.yaml",
	}
}
