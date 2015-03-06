package main

import (
	"bufio"
	"fmt"
	"io"
)

func reader(s io.ReadWriteCloser) {
	reader := bufio.NewScanner(s)

	for reader.Scan() {
		fmt.Printf("%s\n", reader.Text())
	}
}
