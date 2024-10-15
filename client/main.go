package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func MakeClient(host string, port int) {
	// Define the server address and port
	serverAddress := fmt.Sprintf("%s:%d", host, port)

	// Establish a connection to the server
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server at ", serverAddress)

	// Send a request to the server
	_, err = fmt.Fprintln(conn, "Start processing")
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}

	fmt.Println("The request has been sent to the server")
	fmt.Println("The client is waiting")

	// Read the response from the server
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	// Print the server's response
	fmt.Println("Response from server:", response)

	// Wait for user input to close the client
	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
