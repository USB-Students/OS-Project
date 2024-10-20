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
	college, score, err := ProcessFilesParallel(path)
	if err != nil {
		sendMassage(conn, err.Error())
		return
	}
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

func ProcessFilesParallel(path string) (*univercity.College, float64, error) {
	files, err := fileManager.ReadDirectory(path)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading files: %v", err)
	}

	if len(files) <= 1 {
		return nil, 0, fmt.Errorf("you should have 2 or more file to process")
	}

	collegeList := CollegeList{}
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go func() {
			defer log.Printf("Go Routine %d has been processed", goroutine.GoID())
			defer wg.Done()

			college, err := decodeCollegeParallel(path, file)
			if err != nil {
				log.Println(err)
			}
			collegeList.AddCollege(college)
		}()
	}
	wg.Wait()

	list := collegeList.list

	topCollege := list[0]
	topScore := topCollege.CalculateScore()
	for _, college := range list[1:] {
		score := college.CalculateScore()
		if score > topScore {
			topCollege = college
			topScore = score
		}
	}

	return topCollege, topScore, nil
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

func decodeCollegeParallel(path, file string) (*univercity.College, error) {
	records, err := fileManager.ReadCSV(path + "/" + file + ".csv")
	if err != nil {
		return nil, fmt.Errorf("error while reading file %s: %v", path, err)
	}

	college := &univercity.College{
		Name: file,
	}

	for _, row := range records[1:] {
		student, err := decodeStudent(row)
		if err != nil {
			log.Println(err)
			continue
		}
		college.AddStudent(student)
	}
	return college, nil
}

func ProcessFilesSync(path string) (*univercity.SyncCollege, float64, error) {
	files, err := fileManager.ReadDirectory(path)
	if err != nil {
		return nil, 0, fmt.Errorf("error reading files: %v", err)
	}

	if len(files) <= 1 {
		return nil, 0, fmt.Errorf("you should have 2 or more file to process")
	}

	var collegeList []*univercity.SyncCollege

	for _, file := range files {
		college, err := decodeCollegeSync(path, file)
		if err != nil {
			log.Println(err)
		}
		collegeList = append(collegeList, college)
	}

	topCollege := collegeList[0]
	topScore := topCollege.CalculateScore()
	for _, college := range collegeList[1:] {
		score := college.CalculateScore()
		if score > topScore {
			topCollege = college
			topScore = score
		}
	}

	return topCollege, topScore, nil
}

func decodeCollegeSync(path, file string) (*univercity.SyncCollege, error) {
	records, err := fileManager.ReadCSV(path + "/" + file + ".csv")
	if err != nil {
		return nil, fmt.Errorf("error while reading file %s: %v", path, err)
	}

	college := &univercity.SyncCollege{
		Name: file,
	}

	for _, row := range records[1:] {
		student, err := decodeStudent(row)
		if err != nil {
			log.Println(err)
			continue
		}
		college.AddStudent(student)
	}
	return college, nil
}

func decodeStudent(data []string) (*univercity.Student, error) {
	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing ID: %v", err)
	}

	grade, err := strconv.ParseFloat(data[2], 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing Grade: %v", err)
	}

	return &univercity.Student{
		ID:    id,
		Name:  data[1],
		Grade: grade,
	}, nil
}
