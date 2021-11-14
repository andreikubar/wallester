package db

import (
	"fmt"
	"wallester/util"
)

type PgRepository struct {
}

func (r *PgRepository) AddNewCustomer(customer *Customer) uint {
	result := DbConn.Debug().Create(customer)
	util.CheckError(result.Error)
	fmt.Printf("Added new customer with ID %d\n", customer.Id)
	return customer.Id
}

func (r *PgRepository) UpdateCustomer(id uint, updateValues *Customer) *Customer {
	var customer Customer = Customer{Id: id}
	result := DbConn.Debug().Model(&customer).Updates(updateValues)
	util.CheckError(result.Error)
	customerUpdated := r.GetCustomer(id)
	return customerUpdated
}

func (r *PgRepository) GetCustomer(id uint) *Customer {
	var customer Customer
	DbConn.Take(&customer, id)
	return &customer
}

func (r *PgRepository) FindCustomers(firstName string, lastName string, offset int, sort string) []Customer {
	var customers []Customer
	var pageSize int = 10
	DbConn.Debug().Limit(pageSize).Offset(pageSize * offset).Order(sort).
		Where(&Customer{FirstName: firstName, LastName: lastName}).Find(&customers)
	return customers
}
