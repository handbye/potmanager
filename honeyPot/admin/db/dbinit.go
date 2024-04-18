package db

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/pressly/goose/v3"
)

//go:embed sql/001_create_tables.sql

var dbfs embed.FS

func DbInit() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
	}
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(dbfs)

	if err := goose.Up(db, "sql"); err != nil {
		panic(err)
	}
	if err := goose.Version(db, "sql"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("数据库初始化成功!")
	}
}
