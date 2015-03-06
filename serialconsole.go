package main

import (
	"bufio"
	"os"
	"github.com/tarm/serial"
	"log"
	"io"
	"time"
	"fmt"
)

func reader(s io.ReadWriteCloser) {
	reader := bufio.NewScanner(s)

	for reader.Scan() {
		fmt.Printf("%s\n", reader.Text())
	}
}

func writer(s io.ReadWriteCloser) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		if (text == "exit" || text == "quit") {
			os.Exit(0)
		}
		
		_, err := s.Write([]byte(text))
		if err != nil {
			log.Fatal(err)
		}
	}
}


func main() {
	config := NewConfig("/dev/tty.usbmodemfa131", 115200)
	ch := &serial.Config{Name: config.Port, Baud: config.Baud}
	s, err := serial.OpenPort(ch)
	if err != nil {
		log.Fatal(err)
	}

	// wait
	time.Sleep(1 * time.Second)

	go reader(s)
	writer(s)
}

