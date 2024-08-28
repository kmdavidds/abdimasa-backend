package dto

type CreateRemarkRequest struct {
	Name       string `json:"name" validate:"required,min=1,max=255"`
	Occupation string `json:"occupation" validate:"required,min=1,max=255"`
	Gender     string `json:"gender" validate:"required,min=1,max=255"`
}
