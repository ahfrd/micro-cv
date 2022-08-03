package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"micro-cv/constant"
	"micro-cv/src/helpers"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ControllerNamaFungsi is a
func (o ControllerStructure) GetProfile(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)
	res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	listProfile, db, err := o.SelectDetail(id)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		log.Fatalln(err)
		return
	}
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = listProfile
	res.Reply(w)
	return
}
func (o ControllerStructure) ControllerCreateProfile(w http.ResponseWriter, req *http.Request) {
	var request ControllerProfileReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wantedJobTitle := request.WantedJobtitle
	firstname := request.FirstName
	lastname := request.LastName
	email := request.Email
	phone := request.Phone
	country := request.Country
	city := request.City
	address := request.Address
	postalcode := request.PostalCode
	drivinglicense := request.DrivingLicense
	nationality := request.Nationality
	placeOfBirth := request.PlaceOfBirth
	dateOfBirth := request.DateOfBirth
	createInsertWithLastInsertId, db, err := o.InsertProfile(wantedJobTitle, firstname, lastname, email, phone, country, city, address, postalcode, drivinglicense, nationality, placeOfBirth, dateOfBirth)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		log.Fatalln(err)
		return
	}
	resData := ControllerResponseProfileCode{}
	resData.ProfileCode = int(createInsertWithLastInsertId)
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return
}
func (o ControllerStructure) ControllerUpdateProfile(w http.ResponseWriter, req *http.Request) {
	var request ControllerProfileReq
	slug := mux.Vars(req)
	res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)

		return
	}
	wantedJobTitle := request.WantedJobtitle
	firstname := request.FirstName
	lastname := request.LastName
	email := request.Email
	phone := request.Phone
	country := request.Country
	city := request.City
	address := request.Address
	postalcode := request.PostalCode
	drivinglicense := request.DrivingLicense
	nationality := request.Nationality
	placeOfBirth := request.PlaceOfBirth
	dateOfBirth := request.DateOfBirth
	// return
	_, db, err := o.UpdateProfileById(id, wantedJobTitle, firstname, lastname, email, phone, country, city, address, postalcode, drivinglicense, nationality, placeOfBirth, dateOfBirth)

	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Body.Data = id
		log.Fatalln(err)

		return
	}
	resData := ControllerResponseProfileCode{}
	resData.ProfileCode = id
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)

	return
}
