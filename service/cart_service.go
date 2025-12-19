package service

import (
	"encoding/json"
	"go-ec-sample/command"
	"go-ec-sample/db"
	"go-ec-sample/domain"
	"go-ec-sample/viewmodel"

	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
)

type CartService struct{}

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) LoadCart(session sessions.Session) (*domain.Cart, error) {
	raw := session.Get("cart")
	if raw == nil {
		return domain.NewCart(), nil
	}
	var view *viewmodel.CartView
	if err := json.Unmarshal([]byte(raw.(string)), &view); err != nil {
		return nil, err
	}
	return view.ToDomainModel(), nil
}

func (s *CartService) SaveCart(session sessions.Session, cart *domain.Cart) error {
	view := viewmodel.NewCartView(cart)
	b, err := json.Marshal(view)
	if err != nil {
		return err
	}
	session.Set("cart", string(b))
	return session.Save()
}

func (s *CartService) Checkout(cart *domain.Cart) error {
	return db.GetDB().Transaction(func(tx *gorm.DB) error {
		h := command.NewPurchaseCartCommandHandler(tx)
		for _, item := range cart.Items() {
			cmd := command.NewPurchaseCartCommand(item.ProductId(), item.Quantity())
			err := h.Handle(cmd)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
