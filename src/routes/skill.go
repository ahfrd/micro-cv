package routes

import (
	controllers "micro-cv/src/controllers/skill"
	"micro-cv/src/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func SkillRoute(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerCreate := http.HandlerFunc(controllers.ControllerStructure{}.ControllerCreateskill)
	r.Handle(uri(base, "/skill/{id}"), ControllerCreate).Methods("POST")
	ControllerGet := http.HandlerFunc(controllers.ControllerStructure{}.Getskill)
	r.Handle(uri(base, "/skill/{id}"), ControllerGet).Methods("GET")
	ControllerDelete := http.HandlerFunc(controllers.ControllerStructure{}.ControllerDeleteEmployment)
	r.Handle(uri(base, "/skill/{profileCode}"), ControllerDelete).Methods("DELETE")
}
