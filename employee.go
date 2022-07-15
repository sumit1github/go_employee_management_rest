package main

import "gorm.io/gorm"
type Employee struct{
	gorm.Model
	EmpName string `json:"name"`
	EmpSalary float64 `json:"Salary"`
	Email string `json:"email"`
}