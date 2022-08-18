package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kalyani1217/CRUDWithMYSQL/models"
	"github.com/kalyani1217/CRUDWithMYSQL/utils"
)

var NewEmployee models.Employee

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	newEmployees := models.GetAllEmployees()
	res, _ := json.Marshal(newEmployees)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employId := vars["id"]
	ID, err := strconv.ParseInt(employId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	employDetails, _ := models.GetEmployeeById(ID)
	res, _ := json.Marshal(employDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	CreateEmployee := &models.Employee{}
	utils.ParseBody(r, CreateEmployee)
	b := CreateEmployee.CreateEmployee()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employId := vars["id"]
	ID, err := strconv.ParseInt(employId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	employ := models.DeleteEmployee(ID)
	res, _ := json.Marshal(employ)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var updateEmployee = &models.Employee{}
	utils.ParseBody(r, updateEmployee)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetEmployeeById(ID)
	if updateEmployee.Name != "" {
		bookDetails.Name = updateEmployee.Name
	}
	if updateEmployee.Age != 0 {
		bookDetails.Age = updateEmployee.Age
	}
	if updateEmployee.Department != "" {
		bookDetails.Department = updateEmployee.Department
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
