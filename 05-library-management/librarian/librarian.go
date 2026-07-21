package librarian

import (
	"errors"
	"fmt"
	"library-management/person"
)

type Librarian struct {
	person.Person
}

func NewLibrarian(
	id int,
	name string,
	email string,
) (*Librarian, error) {

	if id <= 0 {
		return nil, errors.New("librarian id must be greater than zero")
	}

	if name == "" {
		return nil, errors.New("librarian name cannot be empty")
	}

	if email == "" {
		return nil, errors.New("librarian email cannot be empty")
	}

	librarian := Librarian{
		Person : person.Person{
		ID:    id,
		Name:  name,
		Email: email,
		},
		
	}

	return &librarian, nil
}

func (l *Librarian) UpdateName(name string) error {

	if name == "" {
		return errors.New("librarian name cannot be empty")
	}

	l.Name = name

	return nil
}

func (l *Librarian) UpdateEmail(email string) error {

	if email == "" {
		return errors.New("librarian email cannot be empty")
	}

	l.Email = email

	return nil
}

func (l Librarian) String() string {
	return fmt.Sprintf(
		"Librarian #%d | %s | %s",
		l.ID,
		l.Name,
		l.Email,
	)
}