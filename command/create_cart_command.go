package command

type CreateCartCommand struct {
	UserId uint
}

func NewCreateCartCommand(userId uint) *CreateCartCommand {
	return &CreateCartCommand{UserId: userId}
}
