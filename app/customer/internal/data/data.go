package data

import (
	"frame/app/customer/internal/biz"
	"frame/app/customer/internal/conf"

	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewBD, NewCustomerRepo)

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

func NewCustomerRepo(db fakeDB) biz.CustomerRepo {
	return &customerRepo{
		db: db,
	}
}
