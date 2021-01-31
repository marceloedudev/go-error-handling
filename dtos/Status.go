package dtos

type StatusJSON struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
