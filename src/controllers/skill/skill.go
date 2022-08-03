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
func (o ControllerStructure) Getskill(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)
	res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	listProfile, db, err := o.SelectDetailskill(id)
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
func (o ControllerStructure) ControllerCreateskill(w http.ResponseWriter, req *http.Request) {
	var request ControllerskillReq
	slug := mux.Vars(req)
	id, _ := strconv.Atoi(slug["id"])
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	skill := request.Skill
	level := request.Level
	createskill, db, err := o.Insertskill(skill, level, id)
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
	resData.Id = int(createskill)
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
	_, db, err := o.SkillDelete(id)
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
