package service

import (
	"frame/app/book/internal/biz"

	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(NewBookService)

func NewBookService(usecase *biz.BookUsecase) *BookService {
	return &BookService{
		usecase: usecase,
	}
}
