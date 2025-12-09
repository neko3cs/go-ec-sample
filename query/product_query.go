package query

import "go-ec-sample/domain"

type ProductQuery struct{}

func NewProductQuery() *ProductQuery {
	return &ProductQuery{}
}

func (q *ProductQuery) FindAll() ([]domain.Product, error) {
	// DB未接続のため、ダミーデータを返す
	products := []domain.Product{
		{ID: 1, Name: "Apple", Price: 120},
		{ID: 2, Name: "Banana", Price: 80},
		{ID: 3, Name: "Strawberry", Price: 300},
	}
	return products, nil
}
