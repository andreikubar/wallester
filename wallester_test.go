package main

import (
	"testing"
	"wallester/db"
	"wallester/dto"
	"wallester/service"
	"wallester/test"

	"github.com/go-playground/validator/v10"
)

func TestAdd(t *testing.T) {
	testRepo := &test.TestRepository{}
	service := service.New(testRepo)
	_, errors := service.AddNewCustomer(&db.Customer{})
	if errors == nil {
		t.Error("Validation error expected")
	}
}

func TestUpdate(t *testing.T) {
	testRepo := &test.TestRepository{}
	service := service.New(testRepo)
	_, errors := service.UpdateCustomer(1, &dto.CustomerUpdateDto{EMail: "xx@xx"}, []string{})
	errFound := false
	for i := 0; i < len(errors); i++ {
		fieldError := errors[i].(validator.FieldError)
		if fieldError.Field() == "EMail" {
			errFound = true
		}
	}
	if !errFound {
		t.Error("Expected validation error for field EMail")
	}
}
