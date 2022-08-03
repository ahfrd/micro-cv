package controllers

import (
	"micro-cv/src/database"
	"micro-cv/src/helpers"

	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a
type ControllerWorkingExperienceReq struct {
	WorkingExperience string `json:"workingExperience" validate:"required"`
}
type ControllerWorkingExperienceRes struct {
	WorkingExperience string `json:"workingExperience"`
}
type ControllerStructure struct {
	helpers.ELK
	helpers.Response
	database.TblCvBPJS
}
