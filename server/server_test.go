package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

var (
	host      = "localhost"
	port      = 8001
	directory = "../data"
)

func TestCompareSyncAndParallel(t *testing.T) {
	start := time.Now()
	college, _, err := ProcessFilesParallel(directory)
	if err != nil {
		t.Error(err)
		return
	}
	elapsed := time.Since(start)
	log.Printf("parallel elapsed: %d\n", elapsed.Nanoseconds())
	log.Println(college.String(), "score:", college.CalculateScore())

	start = time.Now()
	syncCollege, _, err := ProcessFilesSync(directory)
	if err != nil {
		t.Error(err)
		return
	}
	elapsed = time.Since(start)
	fmt.Printf("sync elapsed: %d\n", elapsed.Nanoseconds())
	log.Println(syncCollege.String(), "score:", syncCollege.CalculateScore())
}

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
				continue
			}

			go HandleConnection(conn, directory)
		}
	}()

	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
