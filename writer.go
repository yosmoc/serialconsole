package main

import (
	"bufio"
	"io"
	"os"
	"log"
)

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
