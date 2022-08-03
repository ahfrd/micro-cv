package database

import (
	"database/sql"
	"fmt"
	"micro-cv/src/helpers"

	log "github.com/sirupsen/logrus"
)

// TblBwLimitGroup is a
type TblEducation struct {
	Database
	helpers.ELK
}

type TblDataEducation struct {
	Id          int    `json:"id"`
	School      string `json:"school"`
	Degree      string `json:"degree"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	City        string `json:"city"`
	Description string `json:"description"`
}

// SelectRowDataPackage merupakan fungsi untuk get data data_package by id_operator

func (o TblEducation) SelectDetailEducation(idReq int) ([]TblDataEducation, *sql.DB, error) {
	var result []TblDataEducation
	var obj TblDataEducation
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return result, db, fmt.Errorf("%s", err)
	}

	var query string = fmt.Sprintf(`
	SELECT id,school, degree, 
	startDate,endDate, city, description
	from education
	where profileCode = "%d"`, idReq)
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return result, db, fmt.Errorf("failed Select Row SQL for assesment_privy : %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&obj.Id, &obj.School, &obj.Degree, &obj.StartDate,
			&obj.EndDate, &obj.City,
			&obj.Description)

		if err != nil {
			return result, db, fmt.Errorf("failed Select Row SQL for assesment_privy : %v", err)
		}

		result = append(result, obj)
	}
	defer rows.Close()

	return result, db, nil
}
func (o TblEducation) InsertEducation(school string, degree string, profileCode int, startDate string, endDate string, city string, description string) (int64, *sql.DB, error) {
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

	queryInsert := "INSERT INTO education(school,degree,profileCode,startDate,endDate,city,description) values (?,?,?,?,?,?,?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert profile SQL : %v", err)
	}
	res, err = prepare.Exec(school, degree, profileCode, startDate, endDate, city, description)

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

func (o TblEducation) EducationDelete(id int) (int64, *sql.DB, error) {
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

	queryUpdate := "DELETE FROM education WHERE id = ?;"
	prepare, err = db.Prepare(queryUpdate)
	if err != nil {
		log.Warnf("failed to delete tbl education SQL : %v", err)
		return 0, db, fmt.Errorf("failed to delete tbl education SQL : %v", err)
	}
	res, err = prepare.Exec(id)

	if err != nil {
		log.Warnf("failed to delete limit on tbl education SQL : %v", err)
		return 0, db, fmt.Errorf("failed to delete limit on tbl education SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Warnf("failed to populate status delete : %v", err)
		return 0, db, fmt.Errorf("failed to populate status delete : %v", err)
	}
	return count, db, err
}
