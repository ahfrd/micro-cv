package database

import (
	"database/sql"
	"fmt"
	"micro-cv/src/helpers"

	log "github.com/sirupsen/logrus"
)

// TblBwLimitGroup is a
type Tblskill struct {
	Database
	helpers.ELK
}

type TblDataskill struct {
	Id    int    `json:"id"`
	Skill string `json:"skill"`
	Level string `json:"level"`
}

// SelectRowDataPackage merupakan fungsi untuk get data data_package by id_operator

func (o Tblskill) SelectDetailskill(idReq int) ([]TblDataskill, *sql.DB, error) {
	var result []TblDataskill
	var obj TblDataskill
	db, err := o.ConnectDB()
	if err != nil {
		defer db.Close()
		return result, db, fmt.Errorf("%s", err)
	}

	var query string = fmt.Sprintf(`
	SELECT id,skill,level
	from tbl_skill
	where profileCode = "%d"`, idReq)
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		return result, db, fmt.Errorf("failed Select Row SQL for assesment_privy : %v", err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(
			&obj.Id, &obj.Skill, &obj.Level)

		if err != nil {
			return result, db, fmt.Errorf("failed Select Row SQL for assesment_privy : %v", err)
		}

		result = append(result, obj)
	}
	defer rows.Close()

	return result, db, nil
}
func (o Tblskill) Insertskill(skill string, level string, profileCode int) (int64, *sql.DB, error) {
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

	queryInsert := "INSERT INTO tbl_skill(skill,level,profileCode) values (?,?,?)"
	prepare, err = db.Prepare(queryInsert)
	if err != nil {
		return 0, db, fmt.Errorf("failed to insert profile SQL : %v", err)
	}
	res, err = prepare.Exec(skill, level, profileCode)

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

func (o Tblskill) SkillDelete(id int) (int64, *sql.DB, error) {
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

	queryUpdate := "DELETE FROM tbl_skill WHERE id = ?;"
	prepare, err = db.Prepare(queryUpdate)
	if err != nil {
		log.Warnf("failed to delete tbl skill SQL : %v", err)
		return 0, db, fmt.Errorf("failed to delete tbl skill SQL : %v", err)
	}
	res, err = prepare.Exec(id)

	if err != nil {
		log.Warnf("failed to delete limit on tbl skill SQL : %v", err)
		return 0, db, fmt.Errorf("failed to delete limit on tbl skill SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Warnf("failed to populate status delete : %v", err)
		return 0, db, fmt.Errorf("failed to populate status delete : %v", err)
	}
	return count, db, err
}
