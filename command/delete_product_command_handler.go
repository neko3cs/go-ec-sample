package command

import (
	"go-ec-sample/db"

	"gorm.io/gorm"
)

type DeleteProductCommandHandler struct {
	db *gorm.DB
}

func NewDeleteProductCommandHandler(db *gorm.DB) *DeleteProductCommandHandler {
	return &DeleteProductCommandHandler{db: db}
}

func (h *DeleteProductCommandHandler) Handle(command *DeleteProductCommand) error {
	return h.db.Delete(&db.Product{}, command.Id).Error
}
