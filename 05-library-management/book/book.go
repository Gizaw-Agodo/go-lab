package book

import (
	"fmt"
	"errors"
)

type Book struct {
	ISBN string
	Name string 
	Author string 
	Category string 
	TotalCopy int 
	Available int 
}

func Newbook(
	isbn string, 
	name string, 
	author string, 
	category string, 
	totalCopy int, 
)(*Book, error){
	if isbn == "" {
		return nil, errors.New("ISBN cannot be empty")
	}
	if name == ""{
		return nil, errors.New("Name cannot be empty")
	}
	if author == ""{
		return nil, errors.New("Author cannot be empty")
	}
	if category == "" {
		return nil, errors.New("Category cannot be empty")
	}
	if totalCopy == 0 {
		return nil , errors.New("Total copy should be greater than 0")
	}
	newBook := Book {
		ISBN: isbn,
		Name: name,
		Author: author,
		Category: category,
		TotalCopy: totalCopy,
		Available: totalCopy,
	}

	return &newBook, nil
}

func (b *Book) Borrow() error {
	if b.Available == 0 {
		return errors.New("No book available")
	}
	b.TotalCopy -= 1
	return nil
}
 
func (b *Book) Return() error {
	if b.Available == b.TotalCopy {
		return errors.New("All books already exist")
	}
	b.Available += 1
	return nil
}

func (b Book) IsAvailable() bool {
	return b.Available > 0
}

func (b *Book) UpdateCategory(category string) error {
	if category == "" {
		return errors.New("Category cannot be empty")
	}
	b.Category  = category
	return nil
}

func (b Book) String() string {
	return fmt.Sprintf(
		"ISBN: %s | Name: %s | Author: %s | Available: %d/%d",
		b.ISBN,
		b.Name,
		b.Author,
		b.Available,
		b.TotalCopy,
	)
} 

