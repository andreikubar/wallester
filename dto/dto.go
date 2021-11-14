package dto

import (
	"time"
	"wallester/db"

	"github.com/go-playground/validator/v10"
)

type CustomerUpdateDto struct {
	FirstName string `validate:"max=100"`
	LastName  string `validate:"max=100"`
	BirthDate time.Time
	Gender    string `validate:"omitempty,oneof=M F"`
	EMail     string `validate:"omitempty,email"`
	Address   string `validate:"max=200"`
}

func (u *CustomerUpdateDto) ToCustomer() *db.Customer {
	return &db.Customer{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		BirthDate: u.BirthDate,
		Gender:    u.Gender,
		EMail:     u.EMail,
		Address:   u.Address,
	}
}

func (c *CustomerUpdateDto) Validate() []error {
	validate := validator.New()
	if err := validate.Struct(c); err != nil {
		var errors []error
		for i := 0; i < len(err.(validator.ValidationErrors)); i++ {
			errors = append(errors, err.(validator.ValidationErrors)[i])
		}
		return errors
	}
	return nil
}
