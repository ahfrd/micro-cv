package controllers

import (
	"micro-cv/src/database"
	"micro-cv/src/helpers"

	_ "github.com/go-sql-driver/mysql"
)

type PhotoReq struct {
	Base64Image string `json:"base64img"`
}
type ControllerPhotoRes struct {
	ProfileCode int    `json:"profileCode"`
	PhotoUrl    string `json:"photoUrl"`
}
type StructDeletePhotoRes struct {
	ProfileCode int `json:"profileCode"`
}
type ControllerStructurPhoto struct {
	helpers.ELK
	helpers.Response
	database.TblCvBPJS
}
