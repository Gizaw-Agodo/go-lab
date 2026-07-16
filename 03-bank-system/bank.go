package main

import "fmt"

type Bank struct {
	Name string
}

func (b Bank) Welcome() {
	fmt.Println("===================================")
	fmt.Printf(" Welcome to %s\n", b.Name)
	fmt.Println("===================================")
}

func (b Bank) Goodbye() {
	fmt.Println()
	fmt.Println("===================================")
	fmt.Printf(" Thank you for using %s\n", b.Name)
	fmt.Println("===================================")
}

func (b Bank) DisplayInfo() {
	fmt.Println("========== Bank Information ==========")
	fmt.Printf("Bank Name : %s\n", b.Name)
	fmt.Println("======================================")
}
