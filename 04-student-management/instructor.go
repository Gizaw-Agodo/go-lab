package main

import (
	"errors"
	"fmt"
)

type Instructor struct {
	ID int
	Name string
	Department string 
	Email string 
}


func newInstructor(
	id int, 
	name string, 
	department string , 
	email string, 
)(*Instructor, error){
	if id < 0 {
		return nil, errors.New("id should be greater than 0")
	}

	if name == "" {
		return nil , errors.New("Name should not be empty")
	}
	if department == "" {
		return nil , errors.New("department should not be empty")
	}
	if email == "" {
		return nil , errors.New("Email should not be empty")
	}
	instructor := Instructor{
		ID:id , 
		Name : name, 
		Department: department,
		Email: email,
	}
	return &instructor, nil
}

func (i Instructor) display(){
	fmt.Println("---------------")
	fmt.Printf("Name:        %s\n",i.Name)
	fmt.Printf("Department:  %s\n", i.Department)
	fmt.Printf("Email:       %s\n",i.Email)
	fmt.Println("---------------")
}

func (i *Instructor) updateName(name string) error {
	if name == "" {
		return errors.New("name should not be empty")
	}
	i.Name = name
	return nil
}

func (i *Instructor) updateEmail(email string) error {
	if email == ""{
		return errors.New("email should not be empty")
	}
	i.Email = email 
	return nil 
}

func (i *Instructor) updateDepartment(department string ) error {
	if department == "" {
		return errors.New("Department should not be null")
	}
	return nil
}

