package book

import (
	"testing"
)

func TestBorrowBookSuccess(t *testing.T) {

	// Arrange
	book, err := Newbook(
		"9780134190440",
		"The Go Programming Language",
		"Alan Donovan",
		"Programming",
		3,
	)

	if err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	// Act
	err = book.Borrow()
	

	// Assert
	if err != nil {
		t.Fatalf("Borrow() returned an unexpected error: %v", err)
	}

	if book.Available != 2 {
		t.Errorf(
			"expected AvailableCopies to be 2, got %d",
			book.Available,
		)
	}
}

func TestBorrowBookUnavailable(t *testing.T) {

	book, err := Newbook(
		"123",
		"Go",
		"Alan Donovan",
		"Programming",
		1,
	)

	if err != nil {
		t.Fatalf("failed to create book: %v", err)
	}

	if err := book.Borrow();  err != nil {
		t.Fatalf("unexpected error while borrowing first copy: %v", err)
	}

	if err := book.Borrow(); err == nil {
		t.Fatal("expected Borrow() to return an error, got nil")
	}

	expected := 0

	if book.Available != expected {
		t.Errorf(
			"expected AvailableCopies to remain %d, got %d",
			expected,
			book.Available,
		)
	}
}