package data

import (
	"frame/app/book/internal/conf"

	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewBookRepo, NewBD)

func NewBD(conf *conf.Conf) (fakeDB, error) {
	db := make(fakeDB)
	err := db.Dial(
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewBookRepo(db fakeDB) *BookRepo {
	return &BookRepo{
		db: db,
	}
}
