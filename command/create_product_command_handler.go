package command

import (
	"go-ec-sample/db"
)

type CreateProductCommandHandler struct{}

func NewCreateProductCommandHandler() *CreateProductCommandHandler {
	return &CreateProductCommandHandler{}
}

func (h *CreateProductCommandHandler) Handle(command *CreateProductCommand) error {
	p := &db.Product{
		Name:  command.Name,
		Price: command.Price,
	}
	return db.GetDB().Create(p).Error
}
