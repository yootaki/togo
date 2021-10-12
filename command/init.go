package command

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

//create database and tables.
func CmdInit(ctx *cli.Context) {

	isDatabaseExists := exists(dbPath())
	isForceMode := ctx.String("force")

	if isForceMode == "false" && isDatabaseExists {
		fmt.Println("Database is already existed")
		return
	}

	if isForceMode == "true" {
		os.Remove(dbPath())
	}

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
		CREATE TABLE todos (
			id integer NOT NULL PRIMARY KEY,
			title text NOT NULL,
			is_done integer NOT NULL,
			created_at integer NOT NULL,
			updated_at integer NOT NULL
		);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
