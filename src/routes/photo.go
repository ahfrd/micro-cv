package routes

import (
	controllers "micro-cv/src/controllers/photo"
	"micro-cv/src/helpers"
	"net/http"

	"github.com/gorilla/mux"
)

// HealthCheck is a
func PhotoRoutes(r *mux.Router, base string) {
	var uri = helpers.ConRoute
	ControllerUploadPhoto := http.HandlerFunc(controllers.ControllerStructurPhoto{}.ControllerUploadPhoto)
	r.Handle(uri(base, "/photo/{id}"), ControllerUploadPhoto).Methods("PUT")
	ControllerDownloadImage := http.HandlerFunc(controllers.ControllerStructurPhoto{}.ControllerDownloadImage)
	r.Handle(uri(base, "/photo/{id}"), ControllerDownloadImage).Methods("GET")
	ControllerDeleteImage := http.HandlerFunc(controllers.ControllerStructurPhoto{}.ControllerDeleteProfile)
	r.Handle(uri(base, "/photo/{id}"), ControllerDeleteImage).Methods("DELETE")

}
