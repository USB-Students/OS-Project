package univercity

import (
	"fmt"
	"sync"
)

type Student struct {
	ID    int
	Name  string
	Grade float64
}

type College struct {
	Name          string
	students      []*Student
	studentsCount int
	totalGrade    float64
	mu            sync.Mutex
}

func (c *College) AddStudent(student *Student) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.students = append(c.students, student)
	c.studentsCount++
	c.totalGrade += student.Grade
}

func (c *College) CalculateScore() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	avg := c.totalGrade / float64(c.studentsCount)

	topStudents := 0
	for _, student := range c.students {
		if student.Grade > avg {
			topStudents++
		}
	}

	return float64(topStudents) / float64(len(c.students)) * 100
}

func (c *College) String() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return fmt.Sprintf("Name: %s, Student Count: %v, Total Grade: %f, Average Grade: %f",
		c.Name, c.studentsCount, c.totalGrade, c.totalGrade/float64(c.studentsCount))
}

type SyncCollege struct {
	Name          string
	students      []*Student
	studentsCount int
	totalGrade    float64
}

func (c *SyncCollege) AddStudent(student *Student) {
	c.students = append(c.students, student)
	c.studentsCount++
	c.totalGrade += student.Grade
}

func (c *SyncCollege) CalculateScore() float64 {
	avg := c.totalGrade / float64(c.studentsCount)

	topStudents := 0
	for _, student := range c.students {
		if student.Grade > avg {
			topStudents++
		}
	}

	return float64(topStudents) / float64(len(c.students)) * 100
}

func (c *SyncCollege) String() string {
	return fmt.Sprintf("Name: %s, Student Count: %v, Total Grade: %f, Average Grade: %f",
		c.Name, c.studentsCount, c.totalGrade, c.totalGrade/float64(c.studentsCount))
}
