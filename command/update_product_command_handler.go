package command

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"
)

type UpdateProductCommandHandler struct{}

func NewUpdateProductCommandHandler() *UpdateProductCommandHandler {
	return &UpdateProductCommandHandler{}
}

func (h *UpdateProductCommandHandler) Handle(command *UpdateProductCommand) error {
	var p domain.Product
	err := db.GetDB().First(&p, command.Id).Error
	if err != nil {
		return err
	}

	p.Name = command.Name
	p.Price = command.Price

	return db.GetDB().Save(p).Error
}
