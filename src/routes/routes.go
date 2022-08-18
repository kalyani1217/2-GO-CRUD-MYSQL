package routes

import (
	"github.com/gorilla/mux"
	"github.com/kalyani1217/CRUDWithMYSQL/controllers"
)

var RegisterEmployeeRoutes = func(router *mux.Router) {
	router.HandleFunc("/employees/", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees/", controllers.GetEmployee).Methods("GET")
	router.HandleFunc("/employee/{id}", controllers.GetEmployeeById).Methods("GET")
	router.HandleFunc("/employee/{id}", controllers.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/employee/{id}", controllers.DeleteEmployee).Methods("DELETE")
}
