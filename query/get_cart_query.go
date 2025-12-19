package query

type GetCartQuery struct {
	userId uint
}

func NewGetCartQuery(userId uint) *GetCartQuery {
	return &GetCartQuery{userId: userId}
}
