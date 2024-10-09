package dto

type LoginRequest struct {
	ID       string `json:"id" validate:"required,min=1,max=255"`
	Password string `json:"password" validate:"required,min=1,max=255"`
}
