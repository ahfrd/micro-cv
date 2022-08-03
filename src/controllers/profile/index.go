package controllers

import (
	"micro-cv/src/database"
	"micro-cv/src/helpers"

	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a
type ControllerResponseProfileCode struct {
	ProfileCode int `json:"profileCode"`
}
type ControllerGetProfileRes struct {
	ProfileCode    int    `json:"profileCode"`
	WantedJobtitle string `json:"wantedJobTitle"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Address        string `json:"address"`
	PostalCode     int    `json:"postalCode"`
	DrivingLicense string `json:"drivingLicense"`
	Nationality    string `json:"nationality"`
	PlaceOfBirth   string `json:"placeOfBrith"`
	DateOfBirth    string `json:"dateOfBirth"`
	PhotoUrl       string `json:"photoUrl"`
}

type ControllerProfileReq struct {
	WantedJobtitle string `json:"wantedJobTitle" validate:"required"`
	FirstName      string `json:"firstName" validate:"required"`
	LastName       string `json:"lastName" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	Phone          string `json:"phone" validate:"required"`
	Country        string `json:"country" validate:"required"`
	City           string `json:"city" validate:"required"`
	Address        string `json:"address" validate:"required"`
	PostalCode     int    `json:"postalCode" validate:"required"`
	DrivingLicense string `json:"drivingLicense"`
	Nationality    string `json:"nationality" validate:"required"`
	PlaceOfBirth   string `json:"placeOfBirth" validate:"required"`
	DateOfBirth    string `json:"dateOfBirth" validate:"required"`
}

type ControllerStructure struct {
	helpers.ELK
	helpers.Response
	database.TblCvBPJS
}
