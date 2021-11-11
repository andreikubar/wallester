package db

import (
	"fmt"
	"wallester/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test123"
	dbname   = "wallester_test"
)

var DbConn *gorm.DB

func init() {
	dbConn, err := gorm.Open(postgres.Open(
        fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Paris",
        host, user, password, dbname, port)))
	DbConn = dbConn
	util.CheckError(err)
}