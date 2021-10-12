package command

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

//update task.
func CmdUpdate(ctx *cli.Context) {

	if len(ctx.Args()) < 2 {
		fmt.Println("[ERROR] Must set task id and title")
		return
	}

	id := ctx.Args()[0]
	title := strings.Join(ctx.Args()[1:], " ")

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	now := strconv.FormatInt(time.Now().Unix(), 10)

	_, err = db.Exec("UPDATE todos SET title = '" + title + "', updated_at = " + now + " WHERE id = " + id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated task id = " + id)
}
