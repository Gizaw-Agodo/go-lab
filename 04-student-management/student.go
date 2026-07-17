package main

import (
	"errors"
	"fmt"
)

type Student struct {
	ID int 
	Name string
	Age int 
	Email string 
	Program string 
}

func newStudent(
	id int, 
	name string, 
	age int, 
	email string, 
	program string,
)(*Student, error){
	if id <0 {
		return nil, errors.New("Id should be greater than 0")
	}
	if name == ""{
		return nil, errors.New("Name should not be empty")
	}

	if age < 16 {
		return nil, errors.New("Age should be greater than 16")
	}

	if email == ""{
		return nil , errors.New("Email should not be empty")
	}

	if program == "" {
		return nil, errors.New("Program should not be empty")
	}

	student := Student{
		ID : id , 
		Name : name, 
		Email: email, 
		Program : program, 
		Age : age,
	}

	return &student, nil

}

func (s Student) Display(){
	fmt.Println("------student-----")
	fmt.Printf("ID:     %d\n",s.ID)
	fmt.Printf("Name:   %s\n", s.Name)
	fmt.Printf("Age:    %d\n", s.Age)
	fmt.Printf("Email:  %s\n", s.Email)
	fmt.Printf("Program:%s\n", s.Program)
	fmt.Println("------------------")
}


func (s *Student) updateName(name string) error {
	if name == "" {
		return errors.New("Name should not be empty")
	}
	s.Name = name
	return nil
}

func (s *Student) updateAge(age int) error {
	if age < 16 {
		return errors.New("Age should be greater than 16")
	}
	s.Age = age 
	return nil 
}

func (s *Student) updateEmail(email string ) error {
	if email == "" {
		return errors.New("email should not be empty")
	}
	return nil
}

func (s *Student) updateProgram(program string) error {
	if (program == ""){
		return errors.New("Program should not be empty")
	}
	return nil
}