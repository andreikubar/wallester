package db

import (
	"time"
)

type Customer struct {
	Id uint `csv:"id"`
    FirstName string `csv:"first_name"`
	LastName string `csv:"last_name"`
	BirthDate time.Time `csv:"birth_date"`
	Gender Gender `csv:"gender"`
	EMail string `csv:"e_mail"`
	Address string `csv:"address"`
}

type Gender string

const (
	MALE Gender = "M"
	FEMALE = "F"
)

func (g Gender) String() string {
	switch g {
	case MALE:
		return "Male"
	case FEMALE:
		return "Female"
	}
	return "unknown"
}