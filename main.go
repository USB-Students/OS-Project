package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/USB-Students/OS_Project/config"
	"github.com/USB-Students/OS_Project/server"
)

func main() {
	addr := fmt.Sprintf("%s:%d", config.TcpHost, config.TcpPort)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	log.Printf("Server listening on %s \n", addr)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}

			go server.HandleConnection(conn, config.ResultDirectory)
		}
	}()

	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
