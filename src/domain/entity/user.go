package entity

import (
	"github.com/stonelike/CleanGo/src/codes"
	entity "github.com/stonelike/CleanGo/src/domain/entity/internal"
	"github.com/stonelike/CleanGo/src/myerrors"
)

type User struct {
	id   string
	name string
}

func NewUser(name string) (*User, error) {
	u := &User{
		id:   NewID(),
		name: name,
	}

	//エラーのときはnilを返すことを避けた方がいいみたい、errに気づかないで*User(nil)に対して操作をしてしまったとき対策
	if err := u.Validate(); err != nil {
		return &User{}, err
	}

	return u, nil
}

func NewUserFromDB(id, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func (u *User) Validate() error {
	if u.name == "" || len(u.name) > 5 {
		return myerrors.Errorf(codes.InvalidRequest, entity.ErrInvalidEntity, "user")
	}

	return nil
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}
