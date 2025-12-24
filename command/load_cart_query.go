package command

type LoadCartCommand struct {
	userId uint
}

func NewLoadCartCommand(userId uint) *LoadCartCommand {
	return &LoadCartCommand{userId: userId}
}
