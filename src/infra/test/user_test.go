package repo_test

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stonelike/CleanGo/src/domain/entity"
	"github.com/stonelike/CleanGo/src/infra/user"
	"github.com/stretchr/testify/require"
)

func TestUser_Create(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		user  *entity.User
		noErr bool
	}{
		"newUser": {
			user:  entity.NewUserFromDB(entity.NewID(), "test"),
			noErr: true,
		},
	}

	//sql.Open("txdb", uuid.New().String())をすることでテストごとに干渉しあわないDBを作っている
	//t.Run()の粒度で干渉させたくないならt.Runの中に、func Test~の粒度ならfunc Testの内側どこでもで、
	//sql.Open("txdb", uuid.New().String())をすればいい...はず

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			require.NoError(t, err)
			defer db.Close()

			repo := user.NewUserPostgres(db)

			u, err := repo.Create(c.user)

			if c.noErr {
				require.NoError(t, err)
				require.Equal(t, c.user.GetId(), u.GetId())
			} else {
				require.Error(t, err)
			}
		})
	}
}

func Len2(t *testing.T, repo *user.UserPostgres) {
	//repo_test内のfixtureで最初に2ユーザーを作ってある
	list, _ := repo.List()
	require.Equal(t, 2, len(list))
}

func Len3(t *testing.T, repo *user.UserPostgres) {

	ll, _ := repo.List()
	require.Equal(t, 2, len(ll))

	u1 := entity.NewUserFromDB(entity.NewID(), "len3")
	r1, _ := repo.Create(u1)
	require.Equal(t, u1.GetId(), r1.GetId())
	list, _ := repo.List()
	require.Equal(t, 3, len(list))
}

func Len4(t *testing.T, repo *user.UserPostgres) {
	u1 := entity.NewUserFromDB(entity.NewID(), "len41")
	r1, _ := repo.Create(u1)
	require.Equal(t, u1.GetId(), r1.GetId())
	u2 := entity.NewUserFromDB(entity.NewID(), "len42")
	r2, _ := repo.Create(u2)
	require.Equal(t, u2.GetId(), r2.GetId())

	list, _ := repo.List()
	require.Equal(t, 4, len(list))
}

func TestUser_List(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		expectLen int
		fn        func(t *testing.T, repo *user.UserPostgres)
	}{
		"expectLen2": {
			expectLen: 2,
			fn:        Len2,
		},
		"expectLen3": {
			expectLen: 3,
			fn:        Len3,
		},
		"expectLen4": {
			expectLen: 4,
			fn:        Len4,
		},
	}

	//このテストではt.Runレベルで隔離したいため、Openはt.Run()内に書く
	// db, err := sql.Open("txdb", uuid.New().String())
	// require.NoError(t, err)
	// defer db.Close()

	for name, c := range cases {
		c := c
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			db, err := sql.Open("txdb", uuid.New().String())
			require.NoError(t, err)
			defer db.Close()

			repo := user.NewUserPostgres(db)
			c.fn(t, repo)

		})
	}
}
