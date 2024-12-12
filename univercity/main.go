package univercity

import (
	"fmt"
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
}

func (c *College) AddStudent(student *Student) {
	c.students = append(c.students, student)
	c.studentsCount++
	c.totalGrade += student.Grade
}

func (c *College) CalculateScore() float64 {
	avg := c.totalGrade / float64(c.studentsCount)

	topStudents := 0
	for _, student := range c.students {
		if student.Grade > avg {
			topStudents++
		}
	}
	return float64(topStudents) / float64(c.studentsCount) * 100
}

func (c *College) String() string {
	return fmt.Sprintf("Name: %s, Student Count: %v, Total Grade: %f, Average Grade: %f",
		c.Name, c.studentsCount, c.totalGrade, c.totalGrade/float64(c.studentsCount))
}
