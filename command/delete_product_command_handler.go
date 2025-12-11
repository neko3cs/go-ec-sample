package command

import (
	"go-ec-sample/db"
)

type DeleteProductCommandHandler struct{}

func NewDeleteProductCommandHandler() *DeleteProductCommandHandler {
	return &DeleteProductCommandHandler{}
}

func (h *DeleteProductCommandHandler) Handle(command *DeleteProductCommand) error {
	return db.GetDB().Delete(&db.Product{}, command.Id).Error
}
