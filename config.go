package main

type Config struct {
	Serial SerialConfig `yaml:"serial"`
}

type SerialConfig struct {
	Port   string `yaml:"port"`
	Baud   int    `yaml:"baudrate"`
}

	}
}
