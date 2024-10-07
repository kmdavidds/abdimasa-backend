package dto

type CreateDetailRequest struct {
	Slug  string `json:"slug" validate:"required,min=1,max=255"`
	Value string `json:"value" validate:"required"`
}

type GetDetailByIDRequest struct {
	ID uint8 `param:"id" validate:"required,min=1,max=255"`
}

type GetDetailBySlugRequest struct {
	Slug string `param:"slug" validate:"required,min=1,max=255"`
}

type UpdateDetailRequest struct {
	ID    uint8  `json:"id" validate:"required,min=1,max=255"`
	Slug  string `json:"slug" validate:"min=1,max=255"`
	Value string `json:"value" validate:"min=1"`
}

type DeleteDetailRequest struct {
	ID uint8 `param:"id" validate:"required,min=1,max=255"`
}
