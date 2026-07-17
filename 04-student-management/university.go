package main 
import (
	"fmt"
	"errors"
)
type University struct {
	Name string 
	Students []Student
	Instructors []Instructor
	Courses []Course
	Enrollments []Enrollement
}

func newUniversity(name string)(*University, error){
	if name == ""{
		return nil , errors.New("Name should not be empty")
	}
	univeristy := University{
		Name : name,
	}
	return &univeristy , nil
}

func (u *University) addStudent(student Student) error {
	for _, s := range u.Students {
		if s.ID == student.ID {
			return errors.New("Student with this id already exiss")
		}
	}
	u.Students  = append(u.Students, student)
	return nil
}

func (u *University) addCourse(course Course) error {
	for _,c := range u.Courses {
		if c.Code == course.Code {
			return errors.New("course with this code already exist")
		}
	}
	u.Courses = append(u.Courses, course)

	return nil
}

func (u *University) addInstructor(instructor Instructor) error {
	for _,i := range u.Instructors {
		if i.ID == instructor.ID {
			return errors.New("Instructor with this id already exist")
		}
	}
	u.Instructors = append(u.Instructors, instructor)
	
	return nil
}

func (u University) findStudent(id int) (*Student, error) {
	for i := range u.Students {
		if u.Students[i].ID == id {
			return &u.Students[i], nil
		}
	}
	return nil, errors.New("student not found")
   
}

func (u University) displayStudents(){
	fmt.Println("-----students-----")
	for _, student := range u.Students{
		student.Display()
	}
}