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
func (o ControllerStructure) GetEducation(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)
	res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	listProfile, db, err := o.SelectDetailEducation(id)
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
func (o ControllerStructure) ControllerCreateEducation(w http.ResponseWriter, req *http.Request) {
	var request ControllerEducationReq
	slug := mux.Vars(req)
	id, _ := strconv.Atoi(slug["id"])
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	school := request.School
	degree := request.Degree
	startDate := request.StartDate
	endDate := request.EndDate
	city := request.City
	Description := request.Description
	createEducation, db, err := o.InsertEducation(school, degree, id, startDate, endDate, city, Description)
	// createInsertWithLastInsertId, db, err := o.InsertProfile(wantedJobTitle, firstname, lastname, email, phone, country, city, address, postalcode, drivinglicense, nationality, placeOfBirth, dateOfBirth)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		log.Fatalln(err)
		return
	}
	resData := ControllerCreateResponse{}
	resData.Id = int(createEducation)
	resData.ProfileCode = id
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return
}
func (o ControllerStructure) ControllerDeleteEmployment(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)
	res := helpers.Response{}
	profileCode, _ := strconv.Atoi(slug["profileCode"])

	keys, ok := req.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	id, _ := strconv.Atoi(keys[0])
	_, db, err := o.EducationDelete(id)
	// // return
	// _, db, err := o.UpdateProfileById(id, wantedJobTitle, firstname, lastname, email, phone, country, city, address, postalcode, drivinglicense, nationality, placeOfBirth, dateOfBirth)

	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Body.Data = id
		log.Fatalln(err)

		return
	}
	resData := ControllerResponseProfileCode{}
	resData.ProfileCode = profileCode
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)

	return
}
