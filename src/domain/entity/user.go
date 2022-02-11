package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
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

// func (u *User) Validate() error {
// 	if u.name == "" || len(u.name) > 5 {
// 		return myerrors.Errorf(codes.InvalidRequest, entity.ErrInvalidEntity, "user")
// 	}

// 	return nil
// }

//こういうvalidatorを使うならフレームワーク使ってそこにvalidatorを押し込んだ方が良さそう、
//Entityでvalidationするなら手動でやるとかでいいっぽいけど
//myerrorsとの兼ね合いが面倒なので多分フレームワーク使うと思うからvalidationはhandler部分でよさそう

func (u *User) Validate() error {
	//エラーがなければnilを返す
	err := validation.ValidateStruct(u,
		validation.Field(
			&u.name,
			validation.Length(1, 5).Error("名前は1~5文字です"),
			validation.Required.Error("名前は必須入力です"),
		),
	)

	if err != nil {
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
