package main

import (
	"fmt"
	"wallester/db"
	"wallester/service"
	"wallester/web"
)

const (
	port = 8080
)

func main() {
	fmt.Printf("Starting web service at http://localhost:%d\n", port)
	repo := &db.PgRepository{}
	service := service.New(repo)
	controller := web.New(service, port)
	controller.Listen()
}
