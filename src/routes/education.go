package routes

import (
	controllers "micro-cv/src/controllers/education"
	"micro-cv/src/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func EducationRoute(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerCreate := http.HandlerFunc(controllers.ControllerStructure{}.ControllerCreateEducation)
	r.Handle(uri(base, "/education/{id}"), ControllerCreate).Methods("POST")
	ControllerGet := http.HandlerFunc(controllers.ControllerStructure{}.GetEducation)
	r.Handle(uri(base, "/education/{id}"), ControllerGet).Methods("GET")
	ControllerDelete := http.HandlerFunc(controllers.ControllerStructure{}.ControllerDeleteEmployment)
	r.Handle(uri(base, "/education/{profileCode}"), ControllerDelete).Methods("DELETE")
}
