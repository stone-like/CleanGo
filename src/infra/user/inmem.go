package user

import (
	"sync"

	"errors"

	"github.com/stonelike/CleanGo/src/domain/entity"
	in "github.com/stonelike/CleanGo/src/infra/internal"
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

	err := errors.New("User Not Found")
	if !ok {
		return nil, &in.InfraError{
			Message:      "inmem error",
			OrignalError: err,
		}
	}

	return val, nil

}

func (i *inmemRepository) List() ([]*entity.User, error) {
	var list []*entity.User
	return list, nil

}
