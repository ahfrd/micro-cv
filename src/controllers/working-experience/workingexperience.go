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
func (o ControllerStructure) GetWorkingExperience(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)
	res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	listProfile, db, err := o.SelectDetail(id)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Reply(w)
		return
	}
	if listProfile.WorkingExperience == "" {
		res.Body.Code = constant.NotFoundErrorCode
		res.Body.Msg = constant.NotFoundErrorCodeDesc
		res.Reply(w)
		return
	}
	resData := ControllerWorkingExperienceRes{}
	resData.WorkingExperience = listProfile.WorkingExperience
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return
}
func (o ControllerStructure) ControllerUpdateWorkingExperience(w http.ResponseWriter, req *http.Request) {
	var request ControllerWorkingExperienceReq
	slug := mux.Vars(req)
	res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)

		return
	}
	// return
	workingExperience := request.WorkingExperience
	_, db, err := o.UpdateWorkingExperienceByProfileCode(id, workingExperience)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Body.Data = id
		log.Fatalln(err)

		return
	}
	resData := ControllerWorkingExperienceRes{}
	resData.WorkingExperience = workingExperience
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)

	return
}
