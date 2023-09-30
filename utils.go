package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/huandu/go-sqlbuilder"
	_ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var DBConn *sql.DB

func CreateDB() {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = "file:./db.sqlite3"
	}

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		log.Panicln("failed to open connection", err)
	}

	for _, schema := range getSchemas() {
		if _, err := db.Exec(schema); err != nil {
			log.Panicln("failed to run schema", err)
		}
	}

	DBConn = db
}

func getSchemas() []string {
	var schemas []string

	linksTableSql, _ := sqlbuilder.
		NewCreateTableBuilder().
		CreateTable("links").IfNotExists().
		Define("id", "INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT ").
		Define("link", "VARCAR(255)").
		Define("shortened_link", "VARCHAR(255)").
		Build()
	schemas = append(schemas, linksTableSql)

	return schemas
}
