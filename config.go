package main

type Config struct {
	Port   string
	Baud   int
	Prompt string
}

func NewConfig(port string, baud int) *Config {
	return &Config{
		Port:   port,
		Baud:   baud,
	}
}
