package main

type Config struct{
	NextURL *string
	PreviousURL *string
}

func GetConfig() *Config {
	config := Config{
		NextURL: nil,
		PreviousURL: nil,
	}
	return &config
}
