package service

import (
	"time"
	"wallester/db"
)

type IService interface {
	GetCustomer(id uint) *db.Customer
	FindCustomers(firstName string, lastName string, offset int, sort string) []db.Customer
	AddNewCustomer(customer *db.Customer) (uint, []error)
	UpdateCustomer(id uint, updateValues map[string]string, keys []string) (*db.Customer, []error)
}

type Service struct {
	repo db.IRepository
}

func New(repo db.IRepository) IService {
	return &Service{repo: repo}
}

func (this *Service) GetCustomer(id uint) *db.Customer {
	return this.repo.GetCustomer(id)
}

func (this *Service) FindCustomers(firstName string, lastName string, offset int, sort string) []db.Customer {
	return this.repo.FindCustomers(firstName, lastName, offset, sort)
}

func (this *Service) AddNewCustomer(customer *db.Customer) (uint, []error) {
	if validationErrors := customer.Validate(); validationErrors != nil {
		var errors []error
		for i := 0; i < len(validationErrors); i++ {
			errors = append(errors, validationErrors[i])
		}
		return 0, errors
	}
	custId := this.repo.AddNewCustomer(customer)
	return custId, nil
}

func (this *Service) UpdateCustomer(id uint, updateValues map[string]string, keys []string) (*db.Customer, []error) {
	customer := this.GetCustomer(id)

	if value, ok := updateValues["FirstName"]; ok {
		customer.FirstName = value
	}

	if value, ok := updateValues["LastName"]; ok {
		customer.LastName = value
	}

	if value, ok := updateValues["BirthDate"]; ok {
		customer.BirthDate, _ = time.Parse("2006-01-02", value)
	}

	if value, ok := updateValues["Gender"]; ok {
		customer.Gender = value
	}

	if value, ok := updateValues["EMail"]; ok {
		customer.EMail = value
	}

	if value, ok := updateValues["Address"]; ok {
		customer.Address = value
	}


	if validationErrors := customer.Validate(); validationErrors != nil {
		var errors []error
		for i := 0; i < len(validationErrors); i++ {
			errors = append(errors, validationErrors[i])
		}
		return nil, errors
	}
	return this.repo.UpdateCustomer(id, customer, keys), nil
}
