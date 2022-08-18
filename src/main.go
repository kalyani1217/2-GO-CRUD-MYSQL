package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kalyani1217/CRUDWithMYSQL/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterEmployeeRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
