package main 
import (
	"errors"
	"fmt"
)
type Enrollement struct {
	CourseCode string
	StudentId int
	Grade float64
}

func newEnrollement(courseCode string, studentId int)(*Enrollement, error){
	if courseCode == "" {
		return nil, errors.New("invalid course code")
	}
	if studentId < 0 {
		return nil , errors.New("invalid student id ")
	}
	enrollement := Enrollement{
		CourseCode: courseCode,
		StudentId: studentId,
		Grade: 0,
	}
	return &enrollement, nil
}

func (e *Enrollement) assignGrade(grade float64) error {
	if grade < 0 || grade > 100 {
		return errors.New("invalde grade")
	}
	e.Grade = grade
	return nil
}

func (e Enrollement) Passed() bool {
	return e.Grade >= 50
}


func (e Enrollement) Display() {

	fmt.Println("-------- Enrollment --------")
	fmt.Printf("Student ID : %d\n", e.StudentId)
	fmt.Printf("Course     : %s\n", e.CourseCode)
	fmt.Printf("Grade      : %.2f\n", e.Grade)
	fmt.Printf("Passed     : %t\n", e.Passed())
	fmt.Println("----------------------------")
}