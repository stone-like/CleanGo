package user

import "github.com/stonelike/CleanGo/src/domain/entity"

type UseCase interface {
	CreateUser(name string) (*entity.User, error)
	FindById(id string) (*entity.User, error)
}
