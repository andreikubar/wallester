package test

import "wallester/db"

type TestRepository struct {
}

func (r *TestRepository) AddNewCustomer(customer *db.Customer) uint {
	return 10
}

func (r *TestRepository) UpdateCustomer(id uint, updateValues *db.Customer, keys []string) *db.Customer {
	return &db.Customer{}
}

func (r *TestRepository) GetCustomer(id uint) *db.Customer {
	return &db.Customer{}
}

func (r *TestRepository) FindCustomers(firstName string, lastName string, offset int, sort string) []db.Customer {
	return []db.Customer{{}}
}
