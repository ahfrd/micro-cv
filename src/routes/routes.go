package routes

import (
	"github.com/gorilla/mux"
)

// Route is as
func Route() *mux.Router {

	r := mux.NewRouter()
	PhotoRoutes(r, "/api")
	ProfileRoutes(r, "/api")
	WorkingExperienceRoutes(r, "/api")
	EmploymentRoute(r, "/api")
	EducationRoute(r, "/api")
	SkillRoute(r, "/api")

	return r
}
