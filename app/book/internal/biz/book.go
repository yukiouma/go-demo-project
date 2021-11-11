package biz

import (
	"context"
	"errors"
)

type Book struct {
	ID    int
	Name  string
	Saled bool
}

var ErrBookNotFound = errors.New("error: book is not found")
var ErrBookSaled = errors.New("error: this book has been saled")

type BookRepo interface {
	FindBookByID(id int) (*Book, error)
	SaveBook(book *Book) (*Book, error)
	DeleteBook(id int) error
}

type BookUsecase struct {
	repo BookRepo
}

func NewBookUsecase(repo BookRepo) *BookUsecase {
	return &BookUsecase{
		repo: repo,
	}
}

// 按id查找书
func (uc *BookUsecase) FindOneBook(ctx context.Context, id int) (*Book, error) {
	book, err := uc.repo.FindBookByID(id)
	if err != nil {
		return nil, ErrBookNotFound
	}
	return book, nil
}

// 出售一本书（更新saled字段）
func (uc *BookUsecase) SaleOneBook(ctx context.Context, id int) (*Book, error) {
	book, err := uc.FindOneBook(ctx, id)
	if err != nil {
		return nil, err
	}
	if book.Saled {
		return book, ErrBookSaled
	}
	return uc.repo.SaveBook(book)
}

// 上架一本书
func (uc *BookUsecase) NewBook(ctx context.Context, name string) (*Book, error) {
	newBook := &Book{
		Name:  name,
		Saled: false,
	}
	return uc.repo.SaveBook(newBook)
}

// 移除一本书
func (uc *BookUsecase) DeleteBook(ctx context.Context, id int) error {
	err := uc.repo.DeleteBook(id)
	if err != nil {
		return ErrBookNotFound
	}
	return nil
}
