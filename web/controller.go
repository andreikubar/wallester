package web

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"wallester/db"
	"wallester/dto"
	"wallester/service"
	"wallester/util"
)

type Controller struct {
	service service.IService
	port    uint
}

type pageData struct {
	Customers []db.Customer
	Filter    filter
}

type filter struct {
	FirstName string
	LastName  string
}

func New(service service.IService, port uint) Controller {
	return Controller{
		service: service,
		port:    port}
}

func (this *Controller) Listen() {
	http.Handle("/", http.RedirectHandler("/customer/search", 302))
	http.HandleFunc("/customer/search", this.search)
	http.HandleFunc("/customer/show", this.show)
	http.HandleFunc("/customer/new", this.add)
	http.HandleFunc("/customer/edit", this.edit)
	http.ListenAndServe(fmt.Sprintf(":%d", this.port), nil)
}

func (this *Controller) show(writer http.ResponseWriter, request *http.Request) {
	id := getCustomerIdFromQuery(request)
	customer := this.service.GetCustomer(id)
	templ := template.Must(template.ParseFiles("web/templates/show_customer.html"))
	err := templ.Execute(writer, customer)
	util.CheckError(err)
}

func (this *Controller) search(writer http.ResponseWriter, request *http.Request) {
	firstName := request.URL.Query().Get("FirstName")
	lastName := request.URL.Query().Get("LastName")
	page, _ := strconv.Atoi(request.URL.Query().Get("Page"))
	sort := request.URL.Query().Get("Sort")
	customers := this.service.FindCustomers(firstName, lastName, page, sort)
	tmpl, err := template.ParseFiles("web/templates/list_customers.html")
	util.CheckError(err)
	pageData := pageData{
		Customers: customers,
		Filter:    filter{FirstName: firstName, LastName: lastName},
	}
	err = tmpl.Execute(writer, pageData)
	util.CheckError(err)
}

func (this *Controller) add(writer http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "GET" {
		templ := template.Must(template.ParseFiles("web/templates/add_customer.html"))
		err = templ.Execute(writer, nil)
		util.CheckError(err)
	} else if request.Method == "POST" {
		customer := getCustomerData(request)
		custId, errors := this.service.AddNewCustomer(customer)
		if errors != nil {
			displayValidationErrors(writer, errors)
			return
		}
		templ := template.Must(template.ParseFiles("web/templates/add_customer_success.html"))
		err = templ.Execute(writer, custId)
		util.CheckError(err)
	}
}

func (this *Controller) edit(writer http.ResponseWriter, request *http.Request) {
	var err error
	if request.Method == "GET" {
		id := getCustomerIdFromQuery(request)
		customer := this.service.GetCustomer(id)
		templ := template.Must(template.ParseFiles("web/templates/edit_customer.html"))
		err = templ.Execute(writer, customer)
		util.CheckError(err)
	} else if request.Method == "POST" {
		updateValues := getCustomerDataForUpdate(request)
		keys := getFormKeys(request)
		id64, _ := strconv.ParseUint(request.Form.Get("Id"), 10, 32)
		customer, errors := this.service.UpdateCustomer(uint(id64), updateValues, keys)
		if errors != nil {
			displayValidationErrors(writer, errors)
			return
		}
		templ := template.Must(template.ParseFiles("web/templates/show_customer.html"))
		err = templ.Execute(writer, customer)
		util.CheckError(err)
	}
}

func getCustomerData(request *http.Request) *db.Customer {
	firstName, lastName, birthDate, gender, eMail, address := getFormValues(request)
	return &db.Customer{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		Gender:    gender,
		EMail:     eMail,
		Address:   address,
	}
}

func getCustomerDataForUpdate(request *http.Request) *dto.CustomerUpdateDto {
	firstName, lastName, birthDate, gender, eMail, address := getFormValues(request)
	return &dto.CustomerUpdateDto{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		Gender:    gender,
		EMail:     eMail,
		Address:   address,
	}
}

func getFormValues(request *http.Request) (firstName string, lastName string,
	birthDate time.Time, gender string, eMail string, address string) {
	err := request.ParseForm()
	util.CheckError(err)
	firstName = request.Form.Get("FirstName")
	lastName = request.Form.Get("LastName")
	birthDate, _ = time.Parse("2006-01-02", request.Form.Get("BirthDate"))
	gender = request.Form.Get("Gender")
	eMail = request.Form.Get("EMail")
	address = request.Form.Get("Address")
	return
}

func getFormKeys(request *http.Request) []string {
	var keys []string
	for k := range request.Form {
		if k != "Id" {
			keys = append(keys, k)
		}
	}
	return keys
}

func getCustomerIdFromQuery(request *http.Request) uint {
	var id uint64
	id, _ = strconv.ParseUint(request.URL.Query().Get("Id"), 10, 32)
	return uint(id)
}

func displayValidationErrors(writer http.ResponseWriter, errors []error) {
	templ := template.Must(template.ParseFiles("web/templates/display_errors.html"))
	err := templ.Execute(writer, errors)
	util.CheckError(err)
}
