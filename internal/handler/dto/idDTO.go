package dto

type GetByIdDTO struct {
	Id string `params:"id" validate:"required,uuid4"`
}
