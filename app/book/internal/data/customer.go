package data

import (
	"context"
	v1 "frame/api/customer/v1"
	"frame/app/book/internal/biz"

	"google.golang.org/grpc"
)

type CustomerClient struct {
	addr string
}

var _ biz.CustomerClient = new(CustomerClient)

func (c *CustomerClient) FindCustomer(ctx context.Context, id int) (*biz.Customer, error) {
	client, conn, err := c.client()
	defer conn.Close()
	if err != nil {
		return nil, err
	}

	customer, err := client.FindCustomer(ctx, &v1.FindCustomerRequest{
		Id: int64(id),
	})
	if err != nil {
		return nil, err
	}
	return &biz.Customer{
		Id:   int(customer.GetData().Id),
		Name: customer.GetData().Name,
	}, nil
}

func (c *CustomerClient) client() (v1.CustomerServiceClient, *grpc.ClientConn, error) {
	// var opts []grpc.DialOption
	conn, err := grpc.Dial(c.addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	return v1.NewCustomerServiceClient(conn), conn, nil
}
