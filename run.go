package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/url"
	"unicode/utf8"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	shortUrlSize = 10
)

type db_configuration struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
}

var (
	db *sql.DB
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	b_db_conf, err := ioutil.ReadFile("database/conf.json")
	CheckError(err)

	var db_conf db_configuration
	err = json.Unmarshal(b_db_conf, &db_conf)
	CheckError(err)

	db, err = ConnectDB(db_conf.Host, db_conf.Port, db_conf.User, db_conf.Password, db_conf.DBname)
	CheckError(err)
	RunScript(db, "database/create.sql")

	//longU := "https://vk.com/tftf/okji"
	//shortU, err := Create(longU)
	//CheckError(err)
	//fmt.Println(shortU)
	//targetUrl, err := SearchByUrl(db, "long_url", longU)
	//CheckError(err)
	//fmt.Printf("%s:\t\t%s\n", longU, targetUrl)
}

func isValidUrl(longURL string) bool {
	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		return false
	}

	u, err := url.Parse(longURL)
	if err != nil || u.Host == "" {
		return false
	}
	return true
}

func Create(longURL string) (string, error) {
	var err error
	b := make([]byte, shortUrlSize)
	if isValidUrl(longURL) {
		for i := range b {
			b[i] = letterBytes[rand.Intn(utf8.RuneCountInString(letterBytes))]
		}
		err = InsertUrls(db, longURL, string(b))
	} else {
		err = errors.New("Invalid URL")
	}
	return string(b), err
}

func Get(shortURL string) (string, error) {
	return SearchByUrl(db, "short_url", shortURL)
}
