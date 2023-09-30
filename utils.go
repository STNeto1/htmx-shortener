package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io"
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
		Define("created_at", "TIMESTAMP DEFAULT CURRENT_TIMESTAMP").
		Build()
	schemas = append(schemas, linksTableSql)

	return schemas
}

func createLinkHash(link string) (string, error) {
	salt := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	hasher := sha256.New()

	_, err = hasher.Write([]byte(link))
	if err != nil {
		return "", err
	}

	_, err = hasher.Write(salt)
	if err != nil {
		return "", err
	}

	hashBytes := hasher.Sum(nil)

	// We only need the first 12 characters of the hash
	return fmt.Sprintf("%x", hashBytes)[:12], nil
}

func createFullUrl(link string) string {
	prefix := os.Getenv("BASE_URL")

	return fmt.Sprintf("%s/s/%s", prefix, link)
}

func CreateShortenedLink(link string) (string, error) {
	hashedLink, err := createLinkHash(link)
	if err != nil {
		return "", err
	}

	insertSql, args := sqlbuilder.
		NewInsertBuilder().
		InsertInto("links").
		Cols("link", "shortened_link").
		Values(link, hashedLink).
		Build()
	_, err = DBConn.Exec(insertSql, args...)
	if err != nil {
		return "", err
	}

	return createFullUrl(hashedLink), nil
}

func GetLinkByShortenedLink(shortenedLink string) (string, error) {
	var link string

	sb := sqlbuilder.NewSelectBuilder().From("links")

	selectSql, args := sb.
		Select("link").
		Where(
			sb.Equal("shortened_link", shortenedLink),
		).
		Build()
	err := DBConn.QueryRow(selectSql, args...).Scan(&link)
	if err != nil {
		return "", err
	}

	return link, nil
}
