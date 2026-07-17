package main

import (
	"errors"
	"fmt"
)

type Course struct {
	Code string 
	Name string 
	Credit int 
	Capacity int 
	Instructor Instructor 

}

func newCourse(
	code string,
	name string, 
	credit int,
	capacity int , 
	instructor Instructor,
)(*Course, error){
	if code == ""{
		return nil, errors.New("Code should not be empty")
	}
	if name == ""{
		return nil, errors.New("Name should not be empty")
	}
	if credit <= 0 {
		return nil , errors.New("Credit should be greater than 0")
	}

	if capacity <= 0 {
		return nil, errors.New("Capacity should be greater than 0")
	}

    course := Course{
		Code : code, 
		Name : name, 
		Capacity: capacity,
		Credit: credit,
		Instructor: instructor,
	}
	return &course, nil
} 

func (c Course) display(){
	fmt.Println("-----course----")
	fmt.Printf("Code:       %s\n", c.Code)
	fmt.Printf("Name:       %s\n", c.Name)
	fmt.Printf("Credit:     %d\n", c.Credit)
	fmt.Printf("Capacity:   %d\n", c.Capacity)
	fmt.Printf("Instructor: %s\n",c.Instructor.Name)
}

func (c *Course) updateCode(code string) error {
	if code == "" {
		return errors.New("code should not be empty")
	}
	c.Code = code
	return nil
}

func (c *Course) updateName(name string) error {
	if name == "" {
		return errors.New("Name should not be empty")
	}
	c.Name = name
	return nil
}

func (c *Course) updateCredit(credit int) error {
	if credit <= 0 {
		return errors.New("credeit should be greater than 0")
	}
	c.Credit = credit 
	return nil
}

func (c *Course) updateCapacity(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity should be greater than 0")
	}
	c.Capacity = capacity
	return nil
}