package user

import "github.com/stonelike/CleanGo/src/domain/entity"

type UserDao struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

func (u UserDao) ConvertToEntity() *entity.User {
	return entity.NewUserFromDB(u.Id, u.Name)
}

func ConvertDaosToEntities(daos []UserDao) []*entity.User {
	es := make([]*entity.User, len(daos))

	for i, v := range daos {
		es[i] = v.ConvertToEntity()
	}

	return es
}
