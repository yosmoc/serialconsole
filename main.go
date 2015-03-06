package main

import (
	"flag"
	"github.com/tarm/serial"
	"log"
	"time"
)

var (
	confFile = flag.String("conf", "config.yml", "config file path")
)

func main() {
	flag.Parse()

	config, err := ParseConf(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	ch := &serial.Config{Name: config.Serial.Port, Baud: config.Serial.Baud}
	s, err := serial.OpenPort(ch)
	if err != nil {
		log.Fatal(err)
	}

	// wait
	time.Sleep(1 * time.Second)

	go reader(s)
	writer(s)
}
