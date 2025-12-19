package command

import (
	"go-ec-sample/db"

	"gorm.io/gorm"
)

type UpdateProductCommandHandler struct {
	db *gorm.DB
}

func NewUpdateProductCommandHandler(db *gorm.DB) *UpdateProductCommandHandler {
	return &UpdateProductCommandHandler{db: db}
}

func (h *UpdateProductCommandHandler) Handle(command *UpdateProductCommand) error {
	var p db.Product
	err := h.db.First(&p, command.Id).Error
	if err != nil {
		return err
	}

	p.Name = command.Name
	p.Price = command.Price
	p.Stock = command.Stock

	return h.db.Save(p).Error
}
