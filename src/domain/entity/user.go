package entity

import (
	in "github.com/stonelike/CleanGo/src/domain/entity/internal"
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

	if err := u.Validate(); err != nil {
		return nil, err
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
		return &in.GeneralError{
			Err: in.ErrInvalidEntity,
		}
	}

	return nil
}

func (u *User) GetId() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}
