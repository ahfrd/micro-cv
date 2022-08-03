package database

import (
	"database/sql"
	"fmt"
	"micro-cv/src/helpers"

	log "github.com/sirupsen/logrus"
)

// TblBwLimitGroup is a
type TblEmploy struct {
	Database
	helpers.ELK
}

type TblDataEmployer struct {
	Id          int    `json:"id"`
	JobTitle    string `json:"jobTitle"`
	Employer    string `json:"employer"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	City        string `json:"city"`
	Description string `json:"description"`
}

// SelectRowDataPackage merupakan fungsi untuk get data data_package by id_operator

func (o TblEmploy) SelectDetailEmployment(idReq int) ([]TblDataEmployer, *sql.DB, error) {
	var result []TblDataEmployer
	var obj TblDataEmployer
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return result, db, fmt.Errorf("%s", err)
	}

	var query string = fmt.Sprintf(`
	SELECT id,jobTitle, employer, 
	startDate,endDate, city, description
	from employment
	where profileCode = "%d"`, idReq)
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return result, db, fmt.Errorf("failed Select Row SQL for assesment_privy : %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&obj.Id, &obj.JobTitle, &obj.Employer, &obj.StartDate,
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
func (o TblEmploy) InsertEmployment(jobTitle string, employer string, profileCode int, startDate string, endDate string, city string, description string) (int64, *sql.DB, error) {
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

	queryInsert := "INSERT INTO employment(jobTitle,employer,profileCode,startDate,endDate,city,description) values (?,?,?,?,?,?,?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert profile SQL : %v", err)
	}
	res, err = prepare.Exec(jobTitle, employer, profileCode, startDate, endDate, city, description)

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

func (o TblEmploy) EmploymentDelete(id int) (int64, *sql.DB, error) {
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

	queryUpdate := "DELETE FROM employment WHERE id = ?;"
	prepare, err = db.Prepare(queryUpdate)
	if err != nil {
		log.Warnf("failed to delete TblEmploy SQL : %v", err)
		return 0, db, fmt.Errorf("failed to delete TblEmploy SQL : %v", err)
	}
	res, err = prepare.Exec(id)

	if err != nil {
		log.Warnf("failed to delete limit on TblEmploy SQL : %v", err)
		return 0, db, fmt.Errorf("failed to delete limit on TblEmploy SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Warnf("failed to populate status delete : %v", err)
		return 0, db, fmt.Errorf("failed to populate status delete : %v", err)
	}
	return count, db, err
}
