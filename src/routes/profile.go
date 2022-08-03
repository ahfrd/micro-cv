package routes

import (
	controllers "micro-cv/src/controllers/profile"
	"micro-cv/src/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func ProfileRoutes(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerGetProfile := http.HandlerFunc(controllers.ControllerStructure{}.GetProfile)
	r.Handle(uri(base, "/profile/{id}"), ControllerGetProfile).Methods("GET")
	ControllerInsertProfile := http.HandlerFunc(controllers.ControllerStructure{}.ControllerCreateProfile)
	r.Handle(uri(base, "/profile"), ControllerInsertProfile).Methods("POST")
	ControllerUpdateProfile := http.HandlerFunc(controllers.ControllerStructure{}.ControllerUpdateProfile)
	r.Handle(uri(base, "/profile/{id}"), ControllerUpdateProfile).Methods("PUT")

}
