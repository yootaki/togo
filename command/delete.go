package command

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/urfave/cli"
)

//delete list all tasks.
func CmdDelete(ctx *cli.Context) {

	var (
		id string
		s  string
		m  string
	)

	isAllMode := ctx.String("all")

	if len(ctx.Args()) == 1 {
		id = ctx.Args()[0]
	}

	if isAllMode == "true" {
		s = "DELETE FROM todos"
		m = "Deleted all tasks"
	} else {
		s = "DELETE FROM todos WHERE id = " + id
		m = "Deleted task id = " + id
	}

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}
