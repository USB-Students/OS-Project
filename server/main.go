package server

import (
	"fmt"
	"github.com/USB-Students/OS_Project/goroutine"
	"log"
	"net"
	"strconv"
	"sync"

	fileManager "github.com/USB-Students/OS_Project/file"
	"github.com/USB-Students/OS_Project/univercity"
)

func HandleConnection(conn net.Conn, path string) {
	defer conn.Close()
	log.Println("Processing request from", conn.RemoteAddr())
	college, score := processFiles(conn, path)
	sendMassage(conn, fmt.Sprintf("%s, Score: %f", college.String(), score))
}

func processFiles(conn net.Conn, path string) (*univercity.College, float64) {
	files, err := fileManager.ReadDirectory(path)
	if err != nil {
		sendMassage(conn, fmt.Sprintf("Error reading files: %v", err))
		return nil, 0
	}

	if len(files) <= 1 {
		sendMassage(conn, "you should have 2 or more file to process")
		return nil, 0
	}

	getColleges := make(chan *univercity.College, len(files))
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go func() {
			records, err := fileManager.ReadCSV(path + "/" + file)
			if err != nil {
				sendMassage(conn, fmt.Sprintf("error while reading file %s: %v", file, err))
				wg.Done()
				return
			}

			college := &univercity.College{
				Name: file,
			}

			wg2 := sync.WaitGroup{}

			for i, row := range records {
				wg2.Add(1)
				go func() {
					if i == 0 {
						wg2.Done()
						return
					}

					// Convert the fields to the appropriate types
					id, err := strconv.Atoi(row[0])
					if err != nil {
						sendMassage(conn, fmt.Sprintf("Error parsing ID: %v", err))
						wg2.Done()
						return
					}

					grade, err := strconv.ParseFloat(row[2], 64)
					if err != nil {
						sendMassage(conn, fmt.Sprintf("Error parsing Grade: %v", err))
						wg2.Done()
						return
					}

					// Create a new Student struct and append it to the list
					student := univercity.Student{
						ID:    id,
						Name:  row[1],
						Grade: grade,
					}
					college.AddStudent(student)
					routineID := goroutine.GoID()
					log.Printf("Go Routine %d has been processed", routineID)
					wg2.Done()
				}()
			}
			wg2.Wait()
			getColleges <- college
			routineID := goroutine.GoID()
			log.Printf("Go Routine %d has been processed", routineID)
			wg.Done()
		}()
	}

	wg.Wait()

	if len(getColleges) < len(files) {
		sendMassage(conn, "error in reading college a college data")
		return nil, 0
	}

	var list []*univercity.College

loop:
	for {
		college := <-getColleges
		list = append(list, college)

		if len(list) == len(files) {
			break loop
		}
	}

	topScore := 0.0
	topCollege := list[0]
	for _, college := range list[:1] {
		score := college.CalculateScore()
		if score > topScore {
			topCollege = college
			topScore = score
		}
	}

	return topCollege, topScore
}

func sendMassage(conn net.Conn, message string) {
	log.Println("Sending massage to ", conn.RemoteAddr())
	_, err := fmt.Fprintln(conn, message)
	if err != nil {
		log.Printf("Error sending massage: %v", err)
		return
	}
	log.Println("The message was send")
}
