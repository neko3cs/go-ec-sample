package command

import (
	"go-ec-sample/db"

	"gorm.io/gorm"
)

type CreateProductCommandHandler struct {
	db *gorm.DB
}

func NewCreateProductCommandHandler(db *gorm.DB) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{db: db}
}

func (h *CreateProductCommandHandler) Handle(command *CreateProductCommand) error {
	p := &db.Product{
		Name:  command.Name,
		Price: command.Price,
		Stock: command.Stock,
	}
	return h.db.Create(p).Error
}
