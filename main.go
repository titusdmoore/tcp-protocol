package main

import (
	// "bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const PORT = ":6379"

func main() {
	ln, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Panicf("Unable to listent to port: %v", err)
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Panicf("Unable to establish connections: %v", err)
	}
	defer conn.Close()

	fmt.Printf("Listening on Port %s\n", PORT)

	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Panicf("Unable to read from client: %v", err)
		}

		conn.Write([]byte("+OK\r\n"))
	}
}
