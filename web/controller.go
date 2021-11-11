package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"wallester/db"
	"wallester/util"
)

type PageData struct {
	Customers []db.Customer
	Filter Filter
	Page int
}

type Filter struct {
	FirstName string
	LastName string
}

func hello (writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello!\n")
}

func add(writer http.ResponseWriter, request *http.Request) {

}

func edit(writer http.ResponseWriter, request *http.Request) {

}

func search(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("FirstName")
	lastName := request.URL.Query().Get("LastName")
	page, _ := strconv.Atoi(request.URL.Query().Get("Page"))
	customers:=db.FindCustomers(firstName, lastName, page)
	tmpl, err := template.ParseFiles("web/list-customers.html")
	util.CheckError(err)
	pageData := PageData{
		Customers: customers, 
		Filter: Filter{FirstName: firstName, LastName: lastName},
		Page: page,
	}
	err = tmpl.Execute(writer, pageData)
	util.CheckError(err)
}


func Listen() {
	http.HandleFunc("/hello", hello);
	http.HandleFunc("/customer/new", add)
	http.HandleFunc("/customer/edit", edit)
	http.HandleFunc("/customer/search", search)
	http.ListenAndServe(":8080", nil)
}
