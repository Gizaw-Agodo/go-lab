package library
import (
	"fmt"
	"errors"

	"library-management/book"
	"library-management/borrow"
	"library-management/member"
	"library-management/librarian"
)

type Library struct {
	Name string
	Books map[string]*book.Book
	Members map[string]*member.Member
	Librarians map[string]*librarian.Librarian
	Borrows []*borrow.Borrow
}