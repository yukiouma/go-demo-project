package service

import (
	"frame/app/customer/internal/biz"

	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewCustomerService)

func NewCustomerService(usecase *biz.CustomerUsecase) *CustomerService {
	return &CustomerService{
		usecase: usecase,
	}
}
