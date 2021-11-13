package biz

import (
	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewBookUsecase)

func NewBookUsecase(
	repo BookRepo,
	customer CustomerClient,
) *BookUsecase {
	return &BookUsecase{
		repo:     repo,
		customer: customer,
	}
}
