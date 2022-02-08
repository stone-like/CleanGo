package user

import "github.com/stonelike/CleanGo/src/domain/entity"

type Repository interface {
	FindById(id string) (*entity.User, error)
	Create(u *entity.User) (*entity.User, error)
	List() ([]*entity.User, error)
}
