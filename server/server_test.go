package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
)

var (
	host      = "localhost"
	port      = 2000
	directory = "./data"
)

func TestMakeServer(t *testing.T) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on port %d \n", port)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				//t.Errorf("Error accepting connection: %v", err)
				continue
			}

			go HandleConnection(conn, directory)
		}
	}()

	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
