package main

import (
	"fmt"
	"library-management/book"
	"library-management/librarian"
	"library-management/library"
	"library-management/member"
)
func main(){
	lib,err := library.NewLibrary("John f kenedy")
	if err != nil {
		panic(err)
	}
	goBook, err := book.Newbook(
		"12ss2233",
		"The go programming language", 
		"Alan Donovan",
		"programming",
		3,
	)
	if err != nil {
		panic (err)
	}

	cleanCode, err := book.Newbook(
		"9780132350884",
		"Clean Code",
		"Robert C. Martin",
		"Programming",
		2,
	)
	if err != nil {
		panic (err)
	}

	if err := lib.AddBook(goBook); err != nil {
		panic(err)
	}
	if err := lib.AddBook(cleanCode); err != nil {
		panic(err)
	}

	alice, err := member.NewMember(
		1,
		"Alice",
		"alice@example.com",
		"0911111111",
	)
	if err != nil {
		panic(err)
	}

	bob, err := member.NewMember(
		2,
		"Bob",
		"bob@example.com",
		"0922222222",
	)
	if err != nil {
		panic(err)
	}
	// add members to library 
	if err := lib.AddMember(alice); err != nil {
		panic(err)
	}
	if err := lib.AddMember(bob); err != nil {
		panic(err)
	}

	admin, err := librarian.NewLibrarian(
		1,
		"John",
		"john@library.com",
	)
	if err != nil {
		panic(err)
	}
	
	// add librarian to the library 
	if err := lib.AddLibrarian(admin); err != nil {
		panic(err)
	}

	if err := lib.BorrowBook( "9780134190440",1,1,); err != nil {
		fmt.Println(err)
	}

	if err := lib.BorrowBook("9780132350884",1,1,); err != nil {
		fmt.Println(err)
	}
	
}