package service

import (
	"context"
	v1 "frame/api/customer/v1"
	"frame/app/customer/internal/biz"
)

type CustomerService struct {
	usecase *biz.CustomerUsecase
	v1.UnimplementedCustomerServiceServer
}

var _ v1.CustomerServiceServer = new(CustomerService)

func (s *CustomerService) FindCustomer(ctx context.Context, in *v1.FindCustomerRequest) (*v1.CustomerReply, error) {
	customer, err := s.usecase.FindCustomer(int(in.GetId()))
	if err != nil {
		return nil, err
	}
	return &v1.CustomerReply{
		Data: &v1.Customer{
			Id:   int64(customer.ID),
			Name: customer.Name,
		},
		Message: "Getting customer successfully",
	}, nil
}

func (s *CustomerService) RegisterCustomer(ctx context.Context, in *v1.RegisterCustomerRequest) (*v1.CustomerReply, error) {
	customer, err := s.usecase.RegisterCustomer(in.Name)
	if err != nil {
		return nil, err
	}
	return &v1.CustomerReply{
		Data: &v1.Customer{
			Id:   int64(customer.ID),
			Name: customer.Name,
		},
		Message: "Register customer successfully",
	}, nil
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, in *v1.Customer) (*v1.CustomerReply, error) {
	customer, err := s.usecase.UpdateCustomer(&biz.Customer{
		ID:   int(in.GetId()),
		Name: in.Name,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CustomerReply{
		Data: &v1.Customer{
			Id:   int64(customer.ID),
			Name: customer.Name,
		},
		Message: "Update customer successfully",
	}, nil
}

func (s *CustomerService) RemoveCustomer(ctx context.Context, in *v1.FindCustomerRequest) (*v1.RemoveCustomerReply, error) {
	err := s.usecase.DeleteCustomer(int(in.Id))
	if err != nil {
		return nil, err
	}
	return &v1.RemoveCustomerReply{
		Message: "Remove customer successfully",
	}, nil
}
