package biz

type Customer struct {
	ID   int
	Name string
}

type CustomerRepo interface {
	SaveOneCustomer(customer *Customer) (*Customer, error)
	FindCustomer(id int) (*Customer, error)
	DeleteCustomer(id int) error
}

type CustomerUsecase struct {
	repo CustomerRepo
}

// 顾客注册
func (cu *CustomerUsecase) RegisterCustomer(name string) (*Customer, error) {
	customer := &Customer{
		Name: name,
	}
	return cu.repo.SaveOneCustomer(customer)
}

// 顾客查询
func (cu *CustomerUsecase) FindCustomer(id int) (*Customer, error) {
	return cu.repo.FindCustomer(id)
}

// 顾客更名
func (cu *CustomerUsecase) UpdateCustomer(customer *Customer) (*Customer, error) {
	return cu.repo.SaveOneCustomer(customer)
}

// 顾客注销
func (cu *CustomerUsecase) DeleteCustomer(id int) error {
	return cu.repo.DeleteCustomer(id)
}
