package data

import (
	"frame/app/book/internal/biz"
	"frame/app/book/internal/conf"

	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewBD, NewBookRepo)

func NewBD(conf *conf.ConfDB) fakeDB {
	db := make(fakeDB)
	err := db.Dial(
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
	)
	if err != nil {
		panic(err)
	}
	return db
}

func NewBookRepo(db fakeDB) biz.BookRepo {
	return &bookRepo{
		db: db,
	}
}
