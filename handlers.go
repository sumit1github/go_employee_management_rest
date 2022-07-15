package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	message string `json:message`
	status  int    `json:status`
	data    string `json:nested`
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}
func EmployeeList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee
	json.NewDecoder(r.Body).Decode(&employees)
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}
func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	json.NewDecoder(r.Body).Decode(&employee)
	Database.First(&employee, mux.Vars(r)["eid"])

	res := &Response{
		message: "no error",
		status:  200,
		data:    "employee",
	}
	resp, err := json.Marshal(res)
	if err != nil {
		fmt.Print("error")
	}
	json.NewEncoder(w).Encode(resp)
}

func UpdateEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee

	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)
	json.NewEncoder(w).Encode(employee)
}

func DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	Database.Delete(&employee)
	json.NewEncoder(w).Encode("employee is deleted")
}
