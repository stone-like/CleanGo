package user

import (
	"testing"

	"github.com/stonelike/CleanGo/src/domain/entity"
	"github.com/stonelike/CleanGo/src/infra/user"
	"github.com/stretchr/testify/require"
)

func newFixtureUser() *entity.User {
	u, _ := entity.NewUser("test")

	return u
}

func Test_CreateUser(t *testing.T) {
	repo := user.NewInmem()
	s := NewService(repo)
	u := newFixtureUser()

	_, err := s.CreateUser(u.GetName())

	require.NoError(t, err)
}
