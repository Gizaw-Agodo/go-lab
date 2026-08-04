package models

import "time"

type Borrow struct {
	ID int64
	UserID int64
	BookID int64
	BorrowedAt time.Time
	ReturnedAt *time.Time
}