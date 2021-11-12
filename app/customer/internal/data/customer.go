package data

import (
	"errors"
	"frame/app/customer/internal/biz"
	"sync"
	"time"
)

var errNotFound = errors.New("error: Not found")
var errExist = errors.New("error: record existed")
var errNotExist = errors.New("error: record not existed")
var errDialFailed = errors.New("error: dail failed")

type row struct {
	id   int
	name string
}

type fakeDB map[int]*row

type customerRepo struct {
	db fakeDB
	sync.RWMutex
}

var _ biz.CustomerRepo = new(customerRepo)

// simulate connect to db
func (d *fakeDB) Dial(host, port, user, password string) error {
	time.Sleep(time.Millisecond)
	if host == "localhost" && port == "3306" && user == "root" && password == "root" {
		return nil
	}
	return errDialFailed
}

func (r *customerRepo) SaveOneCustomer(customer *biz.Customer) (*biz.Customer, error) {
	r.Lock()
	defer r.Unlock()
	if customer.ID == 0 {
		id := len(r.db) + 1
		customer.ID = id
		return r.createCustomer(customer)
	}
	return r.updateCustomer(customer)
}

func (r *customerRepo) FindCustomer(id int) (*biz.Customer, error) {
	r.RLock()
	defer r.RUnlock()
	customer, ok := r.db[id]
	if !ok {
		return nil, errNotExist
	}
	return &biz.Customer{
		ID:   customer.id,
		Name: customer.name,
	}, nil
}

func (r *customerRepo) DeleteCustomer(id int) error {
	r.Lock()
	defer r.Unlock()
	if _, err := r.FindCustomer(id); err != nil {
		return errNotFound
	}
	delete(r.db, id)
	return nil
}

func (r *customerRepo) createCustomer(customer *biz.Customer) (*biz.Customer, error) {
	if _, ok := r.db[customer.ID]; ok {
		return nil, errExist
	}
	r.db[customer.ID] = &row{
		id:   customer.ID,
		name: customer.Name,
	}
	return customer, nil
}

func (r *customerRepo) updateCustomer(customer *biz.Customer) (*biz.Customer, error) {
	if _, ok := r.db[customer.ID]; !ok {
		return nil, errNotExist
	}
	r.db[customer.ID] = &row{
		id:   customer.ID,
		name: customer.Name,
	}
	return customer, nil
}
