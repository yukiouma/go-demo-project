package biz

import (
	"context"
	"errors"
	"time"
)

type SaleInfo struct {
	SaledAt      time.Time
	CustomerId   int
	CustomerName string
}

type Book struct {
	ID       int
	Name     string
	SaleInfo *SaleInfo
}

var ErrBookNotFound = errors.New("error: book is not found")
var ErrBookSaled = errors.New("error: this book has been saled")

type BookRepo interface {
	FindBookByID(id int) (*Book, error)
	SaveBook(book *Book) (*Book, error)
	DeleteBook(id int) error
}

type BookUsecase struct {
	repo     BookRepo
	customer CustomerClient
}

// 按id查找书
func (uc *BookUsecase) FindOneBook(ctx context.Context, id int) (*Book, error) {
	book, err := uc.repo.FindBookByID(id)
	if err != nil {
		return nil, ErrBookNotFound
	}
	if !book.SaleInfo.SaledAt.IsZero() || book.SaleInfo.CustomerId > 0 {
		customer, err := uc.customer.FindCustomer(ctx, book.SaleInfo.CustomerId)
		if err != nil {
			return book, err
		}
		book.SaleInfo.CustomerName = customer.Name
	}
	return book, nil
}

// 出售一本书
func (uc *BookUsecase) SaleOneBook(ctx context.Context, id, customerId int) (*Book, error) {
	book, err := uc.FindOneBook(ctx, id)
	if err != nil {
		return nil, err
	}
	customer, err := uc.customer.FindCustomer(ctx, customerId)
	if err != nil {
		return nil, err
	}
	if !book.SaleInfo.SaledAt.IsZero() {
		return nil, ErrBookSaled
	}
	book.SaleInfo.CustomerId = customer.Id
	book.SaleInfo.CustomerName = customer.Name
	book.SaleInfo.SaledAt = time.Now()
	return uc.repo.SaveBook(book)
}

// 上架一本书
func (uc *BookUsecase) NewBook(ctx context.Context, name string) (*Book, error) {
	newBook := &Book{
		Name: name,
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
