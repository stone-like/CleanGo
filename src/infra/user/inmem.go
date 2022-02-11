package user

import (
	"sync"

	"github.com/stonelike/CleanGo/src/codes"
	"github.com/stonelike/CleanGo/src/domain/entity"
	"github.com/stonelike/CleanGo/src/myerrors"
)

type inmemRepository struct {
	m map[string]*entity.User
	l sync.Mutex
}

func NewInmem() *inmemRepository {
	m := make(map[string]*entity.User)

	return &inmemRepository{
		m: m,
	}
}

func (i *inmemRepository) Create(e *entity.User) (*entity.User, error) {
	i.l.Lock()
	defer i.l.Unlock()

	i.m[e.GetId()] = e

	return e, nil

}

func (i *inmemRepository) FindById(id string) (*entity.User, error) {
	i.l.Lock()
	defer i.l.Unlock()

	val, ok := i.m[id]

	if !ok {
		return &entity.User{}, myerrors.Errorf(codes.Database, "userId %s is not found", id)
	}

	return val, nil

}

func (i *inmemRepository) List() ([]*entity.User, error) {
	var list []*entity.User
	return list, nil

}
