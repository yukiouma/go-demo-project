package biz

import "github.com/google/wire"

var ProvideSet = wire.NewSet(NewCustomerUsecase)

func NewCustomerUsecase(repo CustomerRepo) *CustomerUsecase {
	return &CustomerUsecase{
		repo: repo,
	}
}
