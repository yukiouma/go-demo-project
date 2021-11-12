package biz

import (
	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewBookUsecase)

func NewBookUsecase(repo BookRepo) *BookUsecase {
	return &BookUsecase{
		repo: repo,
	}
}
