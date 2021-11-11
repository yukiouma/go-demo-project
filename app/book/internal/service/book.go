package service

import (
	"context"
	"frame/api/book/v1"
	"frame/app/book/internal/biz"
)

type BookService struct {
	usecase *biz.BookUsecase
}

var _ book.BookHttpServer = new(BookService)

func (s *BookService) FindOneBook(q *book.QueryBookReq) (*book.BookRes, error) {
	res, err := s.usecase.FindOneBook(context.Background(), q.ID)
	if err != nil {
		return nil, err
	}
	return &book.BookRes{
		Data: &book.Book{
			ID:    res.ID,
			Name:  res.Name,
			Saled: res.Saled,
		},
		Message: "Getting book successfully",
	}, nil
}

func (s *BookService) SaleOneBook(q *book.QueryBookReq) (*book.BookRes, error) {
	res, err := s.usecase.SaleOneBook(context.Background(), q.ID)
	if err != nil {
		return nil, err
	}
	return &book.BookRes{
		Data: &book.Book{
			ID:    res.ID,
			Name:  res.Name,
			Saled: res.Saled,
		},
		Message: "Saling book successfully",
	}, nil
}

func (s *BookService) NewOneBook(q *book.SaveBookReq) (*book.BookRes, error) {
	res, err := s.usecase.NewBook(context.Background(), q.Name)
	if err != nil {
		return nil, err
	}
	return &book.BookRes{
		Data: &book.Book{
			ID:    res.ID,
			Name:  res.Name,
			Saled: res.Saled,
		},
		Message: "Putting a book on the shelf successfully",
	}, nil
}

func (s *BookService) DeleteOneBook(q *book.QueryBookReq) (*book.DeleteBookRes, error) {
	err := s.usecase.DeleteBook(context.Background(), q.ID)
	if err != nil {
		return nil, err
	}
	return &book.DeleteBookRes{
		Message: "Deleting a book successfully",
	}, nil
}
