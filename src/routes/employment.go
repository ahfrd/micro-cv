package routes

import (
	controllers "micro-cv/src/controllers/employment"
	"micro-cv/src/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func EmploymentRoute(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerCreate := http.HandlerFunc(controllers.ControllerStructure{}.ControllerCreateEmployment)
	r.Handle(uri(base, "/employment/{id}"), ControllerCreate).Methods("POST")
	ControllerGet := http.HandlerFunc(controllers.ControllerStructure{}.GetEmployment)
	r.Handle(uri(base, "/employment/{id}"), ControllerGet).Methods("GET")
	ControllerDelete := http.HandlerFunc(controllers.ControllerStructure{}.ControllerDeleteEmployment)
	r.Handle(uri(base, "/employment/{profileCode}"), ControllerDelete).Methods("DELETE")
}
