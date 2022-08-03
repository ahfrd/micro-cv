package database

import (
	"database/sql"
	"fmt"
	"micro-cv/src/helpers"
)

// TblBwLimitGroup is a
type TblCvBPJS struct {
	Database
	helpers.ELK
}

type TblDataProfile struct {
	ProfileCode       int    `json:"profileCode"`
	WantedJobtitle    string `json:"wantedJobTitle"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Country           string `json:"country"`
	City              string `json:"city"`
	Address           string `json:"address"`
	PostalCode        int    `json:"postalCode"`
	DrivingLicense    string `json:"drivingLicense"`
	Nationality       string `json:"nationality"`
	PlaceOfBirth      string `json:"placeOfBrith"`
	DateOfBirth       string `json:"dateOfBirth"`
	PhotoUrl          string `json:"photoUrl"`
	WorkingExperience string `json:"workingExperience"`
}

// SelectRowDataPackage merupakan fungsi untuk get data data_package by id_operator

func (o TblCvBPJS) SelectDetail(idReq int) (TblDataProfile, *sql.DB, error) {
	var result TblDataProfile
	var profileCode sql.NullInt64
	var wantedJobTitle sql.NullString
	var firstName sql.NullString
	var lastName sql.NullString
	var email sql.NullString
	var phone sql.NullString
	var country sql.NullString
	var city sql.NullString
	var address sql.NullString
	var postalCode sql.NullInt64
	var drivingLicense sql.NullString
	var nationality sql.NullString
	var placeOfBirth sql.NullString
	var dateOfBirth sql.NullString
	var photoUrl sql.NullString
	var workingexperience sql.NullString
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return result, db, fmt.Errorf("%s", err)
	}

	var query string = fmt.Sprintf(`
	SELECT profileCode, wantedJobTitle, 
	firstName,lastName, email, 
	phone, country, city,address,postalCode,
	drivingLicense,nationality,placeOfBirth,dateOfBirth,
	photoUrl,workingExperience
	from profile
	where profileCode = "%d"`, idReq)
	defer db.Close()
	err = db.QueryRow(query).Scan(
		&profileCode,
		&wantedJobTitle,
		&firstName,
		&lastName,
		&email,
		&phone,
		&country,
		&city,
		&address,
		&postalCode,
		&drivingLicense,
		&nationality,
		&placeOfBirth,
		&dateOfBirth,
		&photoUrl,
		&workingexperience,
	)
	result.ProfileCode = int(profileCode.Int64)
	result.WantedJobtitle = wantedJobTitle.String
	result.FirstName = firstName.String
	result.LastName = lastName.String
	result.Email = email.String
	result.Phone = phone.String
	result.Country = country.String
	result.City = city.String
	result.Address = address.String
	result.PostalCode = int(postalCode.Int64)
	result.DrivingLicense = drivingLicense.String
	result.Nationality = nationality.String
	result.PlaceOfBirth = placeOfBirth.String
	result.DateOfBirth = dateOfBirth.String
	result.PhotoUrl = photoUrl.String
	result.WorkingExperience = workingexperience.String
	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return result, db, fmt.Errorf("failed Select SQL for tbl_user : %v", err)
	}
	return result, db, nil
}
func (o TblCvBPJS) InsertProfile(wantedJobTitle string, firstname string, lastname string, email string, phone string, country string, city string, address string, postalcode int, drivinglicense string, nationality string, placeOfBirth string, dateOfBirth string) (int64, *sql.DB, error) {
	var err error
	var res sql.Result
	var prepare *sql.Stmt
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return 0, nil, fmt.Errorf("%s", err)
	}
	defer db.Close()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	queryInsert := "INSERT INTO profile (wantedJobTitle, firstName,lastName, email, phone, country, city,address,postalCode,drivingLicense,nationality,placeOfBirth,dateOfBirth) values (?,?,?,?,?,?,?,?,?,?,?,?,?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert profile SQL : %v", err)
	}
	res, err = prepare.Exec(wantedJobTitle, firstname, lastname, email, phone, country, city, address, postalcode, drivinglicense, nationality, placeOfBirth, dateOfBirth)

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert error_general on profile SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return count, db, fmt.Errorf("failed to populate status inserted : %v", err)
	}
	LastinsertId, err := res.LastInsertId()
	return LastinsertId, db, err
}
func (o TblCvBPJS) UpdateProfileById(profileCode int, wantedJobTitle string, firstname string, lastname string, email string, phone string, country string, city string, address string, postalcode int, drivinglicense string, nationality string, placeOfBirth string, dateOfBirth string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE profile set wantedJobTitle = ?, firstName = ?,lastName = ?, email = ?, phone = ?, country = ?, city = ?,address = ?,postalCode = ?,drivingLicense = ?,nationality = ?,placeOfBirth = ?,dateOfBirth = ? where profileCode = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(wantedJobTitle, firstname, lastname, email, phone, country, city, address, postalcode, drivinglicense, nationality, placeOfBirth, dateOfBirth, profileCode)
	defer db.Close()
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	counter, err := res.RowsAffected()
	if err != nil {
		return 0, db, fmt.Errorf("failed to populate status updated : %v", err)
	}
	return counter, db, err
}
func (o TblCvBPJS) UpdatePhotoUrlByProfileCode(profileCode int, photoUrl string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE profile set photoUrl = ? where profileCode = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(photoUrl, profileCode)
	defer db.Close()
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	counter, err := res.RowsAffected()
	if err != nil {
		return 0, db, fmt.Errorf("failed to populate status updated : %v", err)
	}
	return counter, db, err
}
func (o TblCvBPJS) UpdateWorkingExperienceByProfileCode(profileCode int, workingExperience string) (int64, *sql.DB, error) {
	db, err := o.ConnectDB()
	if err != nil {
		return 0, nil, fmt.Errorf("%s", err)
	}

	var queryUpdate string
	queryUpdate = `UPDATE profile set workingExperience = ? where profileCode = ?`
	prepare, err := db.Prepare(queryUpdate)
	res, err := prepare.Exec(workingExperience, profileCode)
	defer db.Close()
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	if err != nil {
		return 0, db, fmt.Errorf("failed to insert tbl profile SQL : %v", err)
	}

	counter, err := res.RowsAffected()
	if err != nil {
		return 0, db, fmt.Errorf("failed to populate status updated : %v", err)
	}
	return counter, db, err
}
