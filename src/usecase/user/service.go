package user

import (
	"github.com/pkg/errors"
	"github.com/stonelike/CleanGo/src/domain/entity"
	"github.com/stonelike/CleanGo/src/domain/repository/user"
)

type Service struct {
	repo user.Repository
}

func NewService(r user.Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateUser(name string) (*entity.User, error) {
	e, err := entity.NewUser(name)

	//ここ後でwrapするように修正
	if err != nil {
		return nil, errors.Wrap(err, "[usecase] createUserError")
	}

	return s.repo.Create(e)
}

func (s *Service) FindById(id string) (*entity.User, error) {
	u, err := s.repo.FindById(id)

	if err != nil {
		return nil, errors.Wrap(err, "[usecase] userFindByError")
	}

	return u, nil
}
