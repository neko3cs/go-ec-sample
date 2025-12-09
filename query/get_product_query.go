package query

type GetProductQuery struct {
	id uint
}

func NewGetProductQuery(id uint) *GetProductQuery {
	return &GetProductQuery{
		id: id,
	}
}
