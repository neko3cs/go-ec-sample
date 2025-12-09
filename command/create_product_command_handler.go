package command

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"
)

type CreateProductCommandHandler struct{}

func NewCreateProductCommandHandler() *CreateProductCommandHandler {
	return &CreateProductCommandHandler{}
}

func (h *CreateProductCommandHandler) Handle(command *CreateProductCommand) error {
	p := &domain.Product{
		Name:  command.Name,
		Price: command.Price,
	}
	return db.GetDB().Create(p).Error
}
