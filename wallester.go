package main

import (
	"fmt"
	"wallester/web"
)

func main() {
	fmt.Println("Starting web service on 8080")
	web.Listen()
}