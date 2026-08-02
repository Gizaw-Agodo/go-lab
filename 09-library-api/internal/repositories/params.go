package repositories

type ListBooksParams struct {
	Page int 
	Limit int 
	Offset int 
}

type CreateUserParams struct {
	Name string 
	Email string 
	PasswordHash string
}
