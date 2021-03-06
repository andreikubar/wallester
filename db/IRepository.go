package db

type IRepository interface {
	AddNewCustomer(customer *Customer) uint
	UpdateCustomer(id uint, updateValues *Customer, keys []string) *Customer
	GetCustomer(id uint) *Customer
	FindCustomers(firstName string, lastName string, offset int, sort string) []Customer
}
