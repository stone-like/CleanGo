package user

import (
	"database/sql"

	"github.com/go-gorp/gorp"

	"github.com/stonelike/CleanGo/src/domain/entity"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) Create(e *entity.User) (*entity.User, error) {
	dbmap := &gorp.DbMap{Db: r.db, Dialect: gorp.PostgresDialect{}}

	t := dbmap.AddTableWithName(UserDao{}, "users")
	t.ColMap("Id").Rename("id")
	t.ColMap("Name").Rename("name")

	err := dbmap.Insert(&UserDao{Id: e.GetId(), Name: e.GetName()})
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (r *UserPostgres) FindById(id string) (*entity.User, error) {

	dbmap := &gorp.DbMap{Db: r.db, Dialect: gorp.PostgresDialect{}}

	var dao UserDao
	err := dbmap.SelectOne(dao, `select id,name user where id=$1`, id)

	if err != nil {
		return nil, err
	}
	return dao.ConvertToEntity(), nil

}

func (r *UserPostgres) List() ([]*entity.User, error) {

	dbmap := &gorp.DbMap{Db: r.db, Dialect: gorp.PostgresDialect{}}

	var userList []UserDao
	_, err := dbmap.Select(&userList, `select * from users`)

	if err != nil {
		return nil, err
	}

	return ConvertDaosToEntities(userList), nil

}
