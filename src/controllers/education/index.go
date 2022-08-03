package controllers

import (
	"micro-cv/src/database"
	"micro-cv/src/helpers"

	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a
type ControllerEducationReq struct {
	School      string `json:"school" validate:"required"`
	Degree      string `json:"degree" validate:"required"`
	StartDate   string `json:"startDate" validate:"required"`
	EndDate     string `json:"endDate" validate:"required"`
	City        string `json:"city" validate:"required"`
	Description string `json:"description"`
}
type ControllerEducationRes struct {
	School      string `json:"school"`
	Degree      string `json:"degree"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	City        string `json:"city"`
	Description string `json:"description"`
}
type ControllerCreateResponse struct {
	ProfileCode int `json:"profileCode"`
	Id          int `json:"id"`
}
type ControllerResponseProfileCode struct {
	ProfileCode int `json:"profileCode"`
}
type ControllerStructure struct {
	helpers.ELK
	helpers.Response
	database.TblEducation
}
