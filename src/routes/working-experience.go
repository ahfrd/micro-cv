package routes

import (
	controllers "micro-cv/src/controllers/working-experience"
	"micro-cv/src/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func WorkingExperienceRoutes(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerUpdate := http.HandlerFunc(controllers.ControllerStructure{}.ControllerUpdateWorkingExperience)
	r.Handle(uri(base, "/working-experience/{id}"), ControllerUpdate).Methods("PUT")
	ControllerGet := http.HandlerFunc(controllers.ControllerStructure{}.GetWorkingExperience)
	r.Handle(uri(base, "/working-experience/{id}"), ControllerGet).Methods("GET")

}
