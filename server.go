package main

import (
	"github.com/albrow/learning/peeps-negroni/controllers"
	"github.com/albrow/learning/peeps-negroni/middleware/recovery"
	"github.com/albrow/learning/peeps-negroni/models"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func main() {
	models.Init()

	router := mux.NewRouter()
	persons := controllers.Persons{}
	router.HandleFunc("/persons", persons.Create).Methods("POST")

	n := negroni.New(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(recovery.JSONRecovery))
	n.UseHandler(router)

	n.Run(":3000")
}