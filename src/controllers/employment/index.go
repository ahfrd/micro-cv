package controllers

import (
	"micro-cv/src/database"
	"micro-cv/src/helpers"

	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a
type ControllerEmployerReq struct {
	JobTitle    string `json:"jobTitle" validate:"required"`
	Employer    string `json:"employer" validate:"required"`
	StartDate   string `json:"startDate" validate:"required"`
	EndDate     string `json:"endDate" validate:"required"`
	City        string `json:"city" validate:"required"`
	Description string `json:"description"`
}
type ControllerEmployerRes struct {
	JobTitle    string `json:"jobTitle"`
	Employer    string `json:"employer"`
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
	database.TblEmploy
}
