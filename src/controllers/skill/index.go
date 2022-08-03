package controllers

import (
	"micro-cv/src/database"
	"micro-cv/src/helpers"

	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a
type ControllerskillReq struct {
	Skill string `json:"skill" validate:"required"`
	Level string `json:"level" validate:"required"`
}
type ControllerskillRes struct {
	Skill string `json:"skill"`
	Level string `json:"level"`
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
	database.Tblskill
}
