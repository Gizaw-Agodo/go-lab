package library

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"library-management/book"
	"library-management/borrow"
	"library-management/librarian"
	"library-management/member"
)

type Library struct {
	Name string
	Books map[string]*book.Book
	Members map[int]*member.Member
	Librarians map[int]*librarian.Librarian
	Borrows []*borrow.Borrow
}

func NewLibrary(name string)(*Library, error){
	if name == "" {
		return nil , errors.New("Name cannot be empty")
	}
	library := Library{
		Name: name,
		Books: make(map[string]*book.Book),
		Members: make(map[int]*member.Member),
		Librarians: make(map[int]*librarian.Librarian),
		Borrows: []*borrow.Borrow{},

	}
	return &library, nil
 
}

func (l *Library) AddBook(book *book.Book)error{
	if _, exists := l.Books[book.ISBN]; exists {
		return errors.New("the book already exists")
	}

	l.Books[book.ISBN] = book
	return nil
}

func (l *Library) AddMember(member *member.Member) error {
	if _, exists := l.Members[member.ID]; exists {
		return errors.New("member already exists")
	}

	l.Members[member.ID] = member
	return nil
}

func (l *Library) AddLibrarian(librarian *librarian.Librarian) error {
	if _, exists := l.Librarians[librarian.ID]; exists {
		return errors.New("Librarian already exists")
	}
	l.Librarians[librarian.ID] = librarian
	return nil
}


func (l *Library) FindBook(isbn string)(*book.Book, error){
	book, exists := l.Books[isbn]
	if !exists {
		return nil, errors.New("Book not found ")
	}
	return book, nil
}

func (l *Library) FindMember(id int)(*member.Member, error){
	member, exists := l.Members[id]
	if !exists {
		return nil, errors.New("Member not found")
	}
	return member, nil 
}

func (l *Library) FindLibrarian(id int)(*librarian.Librarian, error){
	librarian, exists := l.Librarians[id]
	 if !exists {
		return nil, errors.New("Librarian not found")
	 }
	 return librarian, nil
}

func (l *Library) BorrowBook(isbn string, LibrarianId int, memberId int) error {
	if _ , err := l.FindMember(memberId); err  != nil {
		return err
	}
	
	if _, err := l.FindLibrarian(LibrarianId); err != nil {
		return errors.New("Librarian not found")
	}
	
	book, err := l.FindBook(isbn)
	if err != nil {
		return err
	}

	if !book.IsAvailable() {
		return errors.New("Book is not available")
	}

	for _, borrow :=  range l.Borrows {
		if borrow.MemberId == memberId && borrow.ISBN == isbn && borrow.ReturnedAt != nil {
			return errors.New("Member already borrowed the book")
		}
	}

	newBorrow, err := borrow.NewBorrow( memberId, isbn, LibrarianId)

	if err != nil {
		return err
	}

	if err := book.Borrow(); err != nil {
		return err 
	}
	

	l.Borrows = append(l.Borrows,newBorrow)

	return nil
}

func (l *Library) ReturnBook(
	isbn string, memberId int,
) error {
	if _, err := l.FindMember(memberId); err != nil {
		return err
	}
	
	book, err := l.FindBook(isbn)
	if err != nil {
		return err 
	}

	var activeBorrow *borrow.Borrow

	for _,borrow := range l.Borrows {
		if borrow.ISBN == isbn && borrow.MemberId == memberId && borrow.ReturnedAt != nil {
			activeBorrow = borrow
			break
		}
	}
	if activeBorrow == nil {
		return errors.New("Borrow not found")
	}
	if err := activeBorrow.ReturnBook(); err != nil {
		return err 
	}

	if err := book.Return(); err != nil {
		return err 
	}

	return nil
}

func (l Library) ListBooks() {

	if len(l.Books) == 0 {
		fmt.Println("No books found.")
		return
	}

	fmt.Println("Books")
	fmt.Println("-----")

	for _, book := range l.Books {
		fmt.Println(book)
	}
}

func (l Library) ListMembers() {

	if len(l.Members) == 0 {
		fmt.Println("No members found.")
		return
	}

	fmt.Println("Members")
	fmt.Println("-------")

	for _, member := range l.Members {
		fmt.Println(member)
	}
}

func (l Library) ListBorrowRecords() {

	if len(l.Borrows) == 0 {
		fmt.Println("No borrow records found.")
		return
	}

	fmt.Println("Borrow Records")
	fmt.Println("--------------")

	for _, borrow := range l.Borrows {
		fmt.Println(borrow)
	}
}

func (l Library) ExportBooks(filename string) error {
	data, err := json.MarshalIndent(l.Books, "", "   ")
	if err != nil {
		return err 
	}
	if err := os.WriteFile(filename,data, 0644); err != nil {
		return err 
	}
	

	return nil 
}

func (l *Library) ImportBooks(filename string ) error {
	data, err := os.ReadFile(filename)
	
	if err != nil {
		return err 
	}
    
	var books map[string]*book.Book

	if err := json.Unmarshal(data, &books); err != nil {
		return err 
	}
	l.Books = books 
	return nil 
}