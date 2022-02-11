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

	//ここでUsecase層にてwrapしていないけど、myerrorsのスタックトレースでとってこれる
	if err != nil {
		return &entity.User{}, errors.Wrap(err, "[usecase] CreateUser")
	}

	e, err = s.repo.Create(e)
	if err != nil {
		return &entity.User{}, errors.Wrap(err, "[usecase] CreateUser")
	}

	return e, nil
}

func (s *Service) FindById(id string) (*entity.User, error) {
	u, err := s.repo.FindById(id)

	if err != nil {
		return &entity.User{}, errors.Wrap(err, "[usecase] FindById")
	}

	return u, nil
}
