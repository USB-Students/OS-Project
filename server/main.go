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

type CollegeList struct {
	list []*univercity.College
	mu   sync.Mutex
}

func (c *CollegeList) AddCollege(college *univercity.College) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.list = append(c.list, college)
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

	collegeList := CollegeList{}
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			records, err := fileManager.ReadCSV(path + "/" + file)
			if err != nil {
				sendMassage(conn, fmt.Sprintf("error while reading file %s: %v", file, err))
				return
			}

			college := &univercity.College{
				Name: file,
			}

			wg2 := sync.WaitGroup{}

			for _, row := range records[1:] {
				wg2.Add(1)
				go func() {
					defer wg2.Done()
					id, err := strconv.Atoi(row[0])
					if err != nil {
						sendMassage(conn, fmt.Sprintf("Error parsing ID: %v", err))
						return
					}

					grade, err := strconv.ParseFloat(row[2], 64)
					if err != nil {
						sendMassage(conn, fmt.Sprintf("Error parsing Grade: %v", err))
						return
					}

					student := univercity.Student{
						ID:    id,
						Name:  row[1],
						Grade: grade,
					}
					college.AddStudent(student)
					log.Printf("Go Routine %d has been processed", goroutine.GoID())
				}()
			}
			wg2.Wait()
			collegeList.AddCollege(college)
			log.Printf("Go Routine %d has been processed", goroutine.GoID())
		}()
	}
	wg.Wait()

	list := collegeList.list

	topCollege := list[0]
	topScore := topCollege.CalculateScore()
	for _, college := range list[1:] {
		fmt.Println(college.Name)
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
