package borrow

import "time"
import "errors"
import "fmt"

type Borrow struct {
	MemberId int 
	ISBN string 
	LibrarianId int 
	BorrowedAt time.Time
	ReturnedAt *time.Time
}

func NewBorrow(
	memberId int, 
	isbn string, 
	librarianId int,
)(*Borrow, error){
	if memberId < 0 {
		return nil, errors.New("Id shoudl be greater than 0")
	}
	if isbn == ""{
		return nil, errors.New("ISBN should not be empty")
	}
	if librarianId < 0 {
		return nil, errors.New("Librarian id should be greater than 0")
	}
	borrow := Borrow{
		MemberId: memberId,
		ISBN: isbn,
		LibrarianId: librarianId,
		BorrowedAt: time.Now(),
	}
	return &borrow, nil
}

func (b *Borrow) ReturnBook()error {
	if b.ReturnedAt != nil {
		return errors.New("The book is already returned")
	}
	now := time.Now()
	b.ReturnedAt = &now
	return nil
}

func (b Borrow) String() string {

	status := "Borrowed"

	if b.ReturnedAt != nil {
		status = "Returned"
	}

	return fmt.Sprintf(
		"Member %d | Book %s | %s",
		b.MemberId,
		b.ISBN,
		status,
	)
}