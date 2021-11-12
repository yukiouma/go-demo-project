package service

import (
	"context"
	v1 "frame/api/book/v1"
	"frame/app/book/internal/biz"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookService struct {
	usecase *biz.BookUsecase
	v1.UnimplementedBookServiceServer
}

var _ v1.BookServiceServer = new(BookService)

func (s *BookService) FindBook(ctx context.Context, in *v1.FindBookRequest) (*v1.BookReply, error) {
	res, err := s.usecase.FindOneBook(context.Background(), int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.BookReply{
		Data: &v1.Book{
			Id:      int64(res.ID),
			Name:    res.Name,
			SaledAt: timestamppb.New(res.SaledAt),
		},
		Message: "Getting book successfully",
	}, nil
}

func (s *BookService) SaleBook(ctx context.Context, in *v1.SaleBookRequest) (*v1.BookReply, error) {
	res, err := s.usecase.SaleOneBook(context.Background(), int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.BookReply{
		Data: &v1.Book{
			Id:      int64(res.ID),
			Name:    res.Name,
			SaledAt: timestamppb.New(res.SaledAt),
		},
		Message: "Saling book successfully",
	}, nil
}

func (s *BookService) NewBook(ctx context.Context, in *v1.NewBookRequest) (*v1.BookReply, error) {
	res, err := s.usecase.NewBook(context.Background(), in.Name)
	if err != nil {
		return nil, err
	}
	return &v1.BookReply{
		Data: &v1.Book{
			Id:      int64(res.ID),
			Name:    res.Name,
			SaledAt: timestamppb.New(res.SaledAt),
		},
		Message: "Putting a book on the shelf successfully",
	}, nil
}

func (s *BookService) DeleteBook(ctx context.Context, in *v1.DeleteBookRequest) (*v1.DeleteBookReply, error) {
	err := s.usecase.DeleteBook(context.Background(), int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.DeleteBookReply{
		Message: "Deleting a book successfully",
	}, nil
}
