package db

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Customer struct {
	Id        uint      `gorm:"primary_key"`
	FirstName string    `csv:"first_name" validate:"required,max=100"`
	LastName  string    `csv:"last_name" validate:"required,max=100"`
	BirthDate time.Time `csv:"birth_date" validate:"required"`
	Gender    string    `csv:"gender" validate:"required,oneof=M F"`
	EMail     string    `csv:"e_mail" validate:"required,email"`
	Address   string    `csv:"address" validate:"max=200"`
}

func (c *Customer) Validate() []error {
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

func (c *Customer) GenderStr() string {
	switch c.Gender {
	case "M":
		return "Male"
	case "F":
		return "Female"
	default:
		return "unknown"
	}
}
