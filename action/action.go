package action

type Result interface{}

type Action interface {
	Act() Result
}
