package service

import (
	"context"
	v1 "frame/api/book/v1"
	"frame/app/book/internal/biz"
)

type BookService struct {
	usecase *biz.BookUsecase
	v1.UnimplementedBookServiceServer
}

var _ v1.BookServiceServer = new(BookService)

func (s *BookService) FindBook(ctx context.Context, in *v1.FindBookRequest) (*v1.BookReply, error) {
	res, err := s.usecase.FindOneBook(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.BookReply{
		Data: &v1.Book{
			Id:   int64(res.ID),
			Name: res.Name,
			SaleInfo: &v1.SaleInfo{
				SaledAt:      res.SaleInfo.SaledAt.String(),
				CustomerId:   int64(res.SaleInfo.CustomerId),
				CustomerName: res.SaleInfo.CustomerName,
			},
		},
		Message: "Getting book successfully",
	}, nil
}

func (s *BookService) SaleBook(ctx context.Context, in *v1.SaleBookRequest) (*v1.BookReply, error) {
	res, err := s.usecase.SaleOneBook(ctx, int(in.Id), int(in.CustomerId))
	if err != nil {
		return nil, err
	}
	return &v1.BookReply{
		Data: &v1.Book{
			Id:   int64(res.ID),
			Name: res.Name,
			SaleInfo: &v1.SaleInfo{
				SaledAt:      res.SaleInfo.SaledAt.String(),
				CustomerId:   int64(res.SaleInfo.CustomerId),
				CustomerName: res.SaleInfo.CustomerName,
			},
		},
		Message: "Saling book successfully",
	}, nil
}

func (s *BookService) NewBook(ctx context.Context, in *v1.NewBookRequest) (*v1.BookReply, error) {
	res, err := s.usecase.NewBook(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	var (
		saledAt      string
		customerId   int64
		customerName string
	)
	if res.SaleInfo != nil {
		saledAt = res.SaleInfo.SaledAt.String()
		customerId = int64(res.SaleInfo.CustomerId)
		customerName = res.SaleInfo.CustomerName
	}
	return &v1.BookReply{
		Data: &v1.Book{
			Id:   int64(res.ID),
			Name: res.Name,
			SaleInfo: &v1.SaleInfo{
				SaledAt:      saledAt,
				CustomerId:   customerId,
				CustomerName: customerName,
			},
		},
		Message: "Putting a book on the shelf successfully",
	}, nil
}

func (s *BookService) DeleteBook(ctx context.Context, in *v1.DeleteBookRequest) (*v1.DeleteBookReply, error) {
	err := s.usecase.DeleteBook(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteBookReply{
		Message: "Deleting a book successfully",
	}, nil
}
