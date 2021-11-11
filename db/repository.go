package db

import (
	"fmt"
)

func addNewCustomer() {
	fmt.Printf("adding new customer with name %s", "name")

}

func FindCustomers(firstName string, lastName string, offset int) []Customer {
	var customers []Customer
	var pageSize int = 10
	DbConn.Debug().Limit(pageSize).Offset(pageSize * offset).
		Where(&Customer{FirstName: firstName, LastName: lastName}).Find(&customers)
	return customers
}
