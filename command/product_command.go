package command

import (
	"go-ec-sample/db"
	"go-ec-sample/domain"
)

type ProductCommand struct{}

func NewProductCommand() *ProductCommand {
	return &ProductCommand{}
}

func (c *ProductCommand) InsertProduct(p *domain.Product) error {
	return db.GetDB().Create(p).Error
}

func (c *ProductCommand) UpdateProduct(p *domain.Product) error {
	return db.GetDB().Save(p).Error
}
