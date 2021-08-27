package conf

import "os"

type db_configuration struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
}

var (
	DB_conf = db_configuration{
		Host:     "localhost",
		Port:     5432,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBname:   os.Getenv("POSTGRES_DB"),
	}
)
