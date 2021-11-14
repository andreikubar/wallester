package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"wallester/db"
	"wallester/util"
)

func main() {
	DbInit()
}

func DbInit() {
	if hasCutomers := db.DbConn.Migrator().HasTable(&db.Customer{}); hasCutomers == true {
		fmt.Println("Customers table already exists, not doing anything")
		return
	}

	file, err := os.Open("customers.csv")
	util.CheckError(err)
	defer file.Close()
	var customers []db.Customer

	err = gocsv.Unmarshal(file, &customers)
	util.CheckError(err)
	err = db.DbConn.AutoMigrate(&db.Customer{})
	util.CheckError(err)
	result := db.DbConn.Create(customers)
	util.CheckError(result.Error)
	fmt.Println("Loaded customers data")
}
