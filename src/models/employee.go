package models

import (
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	"github.com/kalyani1217/CRUDWithMYSQL/config"
)

var (
	db *gorm.DB
)

type Employee struct {
	Name       string `json:"name"`
	Id         int    `json:"id"`
	Age        int    `json:"age"`
	Department string `json:"department"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Employee{})
}

func (b *Employee) CreateEmployee() *Employee {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllEmployees() []Employee {
	var Employees []Employee
	db.Find(&Employees)
	return Employees
}

func GetEmployeeById(Id int64) (*Employee, *gorm.DB) {
	var getEmployee Employee
	db := db.Where("id=?", Id).Find(&getEmployee)
	return &getEmployee, db
}

func DeleteEmployee(ID int64) Employee {
	var employ Employee
	db.Where("id=?", ID).Delete(employ)
	return employ
}
