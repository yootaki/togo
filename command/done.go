package command

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/urfave/cli"
)

//mark task as finished.
func CmdDone(ctx *cli.Context) {

	if len(ctx.Args()) != 1 {
		return
	}

	id := ctx.Args()[0]
	fmt.Printf("Task %s is done\n", id)

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE todos SET is_done = 1 WHERE id = " + id)
	if err != nil {
		log.Fatal(err)
	}
}
