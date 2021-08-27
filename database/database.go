package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"shortener/conf"
)

const table_name string = `addresses`

var db *sql.DB = nil

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func RunScript(db *sql.DB, filePath string) (interface{}, error) {
	c, err := ioutil.ReadFile(filePath)
	CheckError(err)

	res, err := db.Exec(string(c))
	CheckError(err)

	return res, err
}

func ConnectDB(host, user, password, dbname string, port int) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	return db, err
}

func GetConn() *sql.DB {
	var err error
	if db == nil {
		db, err = ConnectDB(conf.DB_conf.Host, conf.DB_conf.User, conf.DB_conf.Password, conf.DB_conf.DBname, conf.DB_conf.Port)
		CheckError(err)
	}

	return db
}

func InsertUrls(db *sql.DB, long_url string, short_url string) error {
	insertDynStmt := fmt.Sprintf(`INSERT INTO %s (long_url, short_url) VALUES ('%s', '%s');`, table_name, long_url, short_url)
	_, err := db.Exec(insertDynStmt)
	return err
}

func SearchByUrl(db *sql.DB, nameField, valueField string) (string, error) {
	var (
		targetValue, nameTargetField string
	)

	switch nameField {
	case "long_url":
		nameTargetField = "short_url"
	case "short_url":
		nameTargetField = "long_url"
	}

	q := fmt.Sprintf(`SELECT %s FROM %s WHERE %s = '%s';`, nameTargetField, table_name, nameField, valueField)
	err := db.QueryRow(q).Scan(&targetValue)

	if err != nil || targetValue == "" {
		err = errors.New("Not found")
	}

	return targetValue, err
}
