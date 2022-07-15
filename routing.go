package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/create_employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee_list", EmployeeList).Methods("GET")
	r.HandleFunc("/get_employee/{eid}", GetEmployeeById).Methods("GET")
	r.HandleFunc("/update_employee/{eid}", UpdateEmployeeById).Methods("PUT")
	r.HandleFunc("/delete_employee/{eid}", DeleteEmployeeById).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
