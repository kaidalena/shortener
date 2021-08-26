package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
)

type db_configuration struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
}

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

func ConnectDB(host, port, user, password, dbname string) (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	return db, err
}

func GetConn() *sql.DB {
	if db == nil {
		b_db_conf, err := ioutil.ReadFile("database/conf.json")
		CheckError(err)

		var db_conf db_configuration
		err = json.Unmarshal(b_db_conf, &db_conf)
		CheckError(err)

		db, err = ConnectDB(db_conf.Host, db_conf.Port, db_conf.User, db_conf.Password, db_conf.DBname)
		CheckError(err)

		RunScript(db, "database/create.sql")
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
