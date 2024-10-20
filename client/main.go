package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func MakeClient(host string, port int) {
	serverAddress := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server at ", serverAddress)

	_, err = fmt.Fprintln(conn, "Start processing")
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	fmt.Println("The request has been sent to the server")
	fmt.Println("The client is waiting")

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	fmt.Println("Response from server:", response)

	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
