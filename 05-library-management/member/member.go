package member

import (
	"errors"
	"fmt"

	"library-management/person"
)

type Member struct {
	person.Person
	Phone string
}

func NewMember(
	id int,
	name string,
	email string,
	phone string,
) (*Member, error) {

	if id <= 0 {
		return nil, errors.New("member id must be greater than zero")
	}

	if name == "" {
		return nil, errors.New("member name cannot be empty")
	}

	if email == "" {
		return nil, errors.New("member email cannot be empty")
	}

	if phone == "" {
		return nil, errors.New("member phone cannot be empty")
	}

	member := Member{
		Person: person.Person{
		ID:    id,
		Name:  name,
		Email: email,
		},

		Phone: phone,
	}

	return &member, nil
}

func (m *Member) UpdateName(name string) error {

	if name == "" {
		return errors.New("member name cannot be empty")
	}

	m.Name = name

	return nil
}

func (m *Member) UpdateEmail(email string) error {

	if email == "" {
		return errors.New("member email cannot be empty")
	}

	m.Email = email

	return nil
}

func (m *Member) UpdatePhone(phone string) error {

	if phone == "" {
		return errors.New("member phone cannot be empty")
	}

	m.Phone = phone

	return nil
}

func (m Member) String() string {
	return fmt.Sprintf(
		"Member #%d | %s | %s",
		m.ID,
		m.Name,
		m.Email,
	)
}