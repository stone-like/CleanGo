package entity

// //ErrNotFound not found
// var ErrNotFound = errors.New("Not found")

// //ErrInvalidEntity invalid entity
// var ErrInvalidEntity = errors.New("Invalid entity")

//Entityで共通のエラー
const (
	ErrInvalidEntity = "InvalidEntity %s"
)

//Codesを使うなら下記は使用しない

type GeneralError struct {
	Err error
}

func (g *GeneralError) Code() string {
	return "500"
}

func (g *GeneralError) Error() string {
	return g.Err.Error()
}
