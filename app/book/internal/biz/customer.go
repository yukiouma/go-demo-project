package biz

import "context"

type Customer struct {
	Id   int
	Name string
}

type CustomerClient interface {
	FindCustomer(ctx context.Context, id int) (*Customer, error)
}
