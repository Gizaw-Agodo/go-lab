package main 

import "fmt"

type Customer struct {
	ID int 
	Name string 
	Email string 
}


func (c Customer) display(){
	fmt.Println("=====customer====")
	fmt.Printf("ID:     %d\n", c.ID)
	fmt.Printf("Name:   %s\n", c.Name)
	fmt.Printf("Email:  %s\n", c.Email)
	fmt.Println("=================")
}

func (c *Customer) updateEmail(email string){
	if email == "" {
		fmt.Println("Email cannot be empty")
		return
	}

	c.Email = email
	fmt.Println("email updated successfully!")
}

func (c *Customer) updateName(name string) {
	if name == ""{
		fmt.Println("Name cannot be empty")
		return 
	}
	c.Name = name
	fmt.Println("Name updated successfully")
}

