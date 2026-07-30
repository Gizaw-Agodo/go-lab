package dto

type CreateBookRequest struct {
	Title  string `json:"title" validate:"required,min=3,max=100"`
	Author string `json:"author" validate:"required,min=3,max=100"`
}

type UpdateBookRequest struct {
	Title  string `json:"title" validate:"required,min=3,max=100"`
	Author string `json:"author" validate:"required,min=3,max=100"`
}