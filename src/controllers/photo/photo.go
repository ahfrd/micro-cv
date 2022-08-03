package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"micro-cv/constant"
	"micro-cv/src/helpers"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// ControllerNamaFungsi is a
func (o ControllerStructurPhoto) ControllerDownloadImage(w http.ResponseWriter, req *http.Request) {

	slug := mux.Vars(req)
	res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	fmt.Println(id)
	selectPhoto, db, err := o.SelectDetail(id)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf(err.Error())
		res.Reply(w)
		return
	}
	if selectPhoto.PhotoUrl == "" {
		res.Body.Code = constant.NotFoundErrorCode
		res.Body.Msg = "Data Not Found"
		res.Reply(w)
		return
	}
	// get data photo url from table and change "file.png" to retrive data from db
	// pathFile := "app/upload/photo/"
	bytes, err := ioutil.ReadFile(selectPhoto.PhotoUrl)
	if err != nil {
		log.Fatal(err)
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/gif":
		base64Encoding += "data:image/gif;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += helpers.ToBase64(bytes)
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = base64Encoding
	res.Reply(w)
	return
}
func (o ControllerStructurPhoto) ControllerUploadPhoto(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)
	var bodyReq PhotoReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)

		return
	}
	// res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	slugId := strconv.Itoa(id)
	//data := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAApgAAAKYB3X3/OAAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAANCSURBVEiJtZZPbBtFFMZ/M7ubXdtdb1xSFyeilBapySVU8h8OoFaooFSqiihIVIpQBKci6KEg9Q6H9kovIHoCIVQJJCKE1ENFjnAgcaSGC6rEnxBwA04Tx43t2FnvDAfjkNibxgHxnWb2e/u992bee7tCa00YFsffekFY+nUzFtjW0LrvjRXrCDIAaPLlW0nHL0SsZtVoaF98mLrx3pdhOqLtYPHChahZcYYO7KvPFxvRl5XPp1sN3adWiD1ZAqD6XYK1b/dvE5IWryTt2udLFedwc1+9kLp+vbbpoDh+6TklxBeAi9TL0taeWpdmZzQDry0AcO+jQ12RyohqqoYoo8RDwJrU+qXkjWtfi8Xxt58BdQuwQs9qC/afLwCw8tnQbqYAPsgxE1S6F3EAIXux2oQFKm0ihMsOF71dHYx+f3NND68ghCu1YIoePPQN1pGRABkJ6Bus96CutRZMydTl+TvuiRW1m3n0eDl0vRPcEysqdXn+jsQPsrHMquGeXEaY4Yk4wxWcY5V/9scqOMOVUFthatyTy8QyqwZ+kDURKoMWxNKr2EeqVKcTNOajqKoBgOE28U4tdQl5p5bwCw7BWquaZSzAPlwjlithJtp3pTImSqQRrb2Z8PHGigD4RZuNX6JYj6wj7O4TFLbCO/Mn/m8R+h6rYSUb3ekokRY6f/YukArN979jcW+V/S8g0eT/N3VN3kTqWbQ428m9/8k0P/1aIhF36PccEl6EhOcAUCrXKZXXWS3XKd2vc/TRBG9O5ELC17MmWubD2nKhUKZa26Ba2+D3P+4/MNCFwg59oWVeYhkzgN/JDR8deKBoD7Y+ljEjGZ0sosXVTvbc6RHirr2reNy1OXd6pJsQ+gqjk8VWFYmHrwBzW/n+uMPFiRwHB2I7ih8ciHFxIkd/3Omk5tCDV1t+2nNu5sxxpDFNx+huNhVT3/zMDz8usXC3ddaHBj1GHj/As08fwTS7Kt1HBTmyN29vdwAw+/wbwLVOJ3uAD1wi/dUH7Qei66PfyuRj4Ik9is+hglfbkbfR3cnZm7chlUWLdwmprtCohX4HUtlOcQjLYCu+fzGJH2QRKvP3UNz8bWk1qMxjGTOMThZ3kvgLI5AzFfo379UAAAAASUVORK5CYII="
	data := bodyReq.Base64Image
	datax := strings.Split(data, ",")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(datax[1]))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println(bounds, formatString)
	pathFile := "app/upload/photo/"
	//Encode from image format to writer
	pngFilename := pathFile + slugId + "." + formatString
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}
	switch formatString {
	case "png":
		err = png.Encode(f, m)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Png file", pngFilename, "created")
	case "jpeg":
		err = jpeg.Encode(f, m, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Jpeg file", pngFilename, "created")
	case "gif":
		err = gif.Encode(f, m, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Jpeg file", pngFilename, "created")
	}
	_, db, err := o.UpdatePhotoUrlByProfileCode(id, pngFilename)
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Body.Data = id
		log.Fatalln(err)

		return
	}
	resData := ControllerPhotoRes{}
	resData.PhotoUrl = pngFilename
	resData.ProfileCode = id
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)
	return

}
func (o ControllerStructurPhoto) ControllerDeleteProfile(w http.ResponseWriter, req *http.Request) {
	slug := mux.Vars(req)
	res := helpers.Response{}
	id, _ := strconv.Atoi(slug["id"])
	selectPhoto, db, errs := o.SelectDetail(id)
	db.Close()
	if errs != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf(errs.Error())
		res.Reply(w)
		return
	}
	if selectPhoto.PhotoUrl == "" {
		res.Body.Code = constant.NotFoundErrorCode
		res.Body.Msg = "Data Not Found"
		res.Reply(w)
		return
	}

	path := selectPhoto.PhotoUrl
	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
		return
	}
	_, db, err = o.UpdatePhotoUrlByProfileCode(id, "")
	db.Close()
	if err != nil {
		res.Body.Code = constant.RCDatabaseError
		res.Body.Msg = fmt.Sprintf("%v", err.Error())
		res.Body.Data = id
		log.Fatalln(err)

		return
	}
	resData := StructDeletePhotoRes{}
	resData.ProfileCode = id
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = resData
	res.Reply(w)

	return
}
