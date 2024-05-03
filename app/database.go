package app

import (
	"database/sql"
	"pilotkode/belajar-golang-restfull-api/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:rendyganteng@tcp(localhost:3306)/belajar_golang_resful_api")
	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
