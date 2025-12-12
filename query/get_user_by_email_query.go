package query

type GetUserByEmailQuery struct {
	Email string
}

func NewGetUserByEmailQuery(email string) *GetUserByEmailQuery {
	return &GetUserByEmailQuery{
		Email: email,
	}
}
