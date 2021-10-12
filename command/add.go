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

//add new task
func CmdAdd(ctx *cli.Context) {
	if len(ctx.Args()) == 0 {
		fmt.Println("[ERROR] Must set task title")
		return
	}

	title := strings.Join(ctx.Args(), " ")
	fmt.Println("Added new task :", title)

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		fmt.Println("open err")
		log.Fatal(err)
	}
	defer db.Close()

	now := strconv.FormatInt(time.Now().Unix(), 10)

	_, err = db.Exec("INSERT INTO todos(title, is_done, created_at, updated_at) values('" + title + "', 0, " + now + ", " + now + ")")
	if err != nil {
		fmt.Println("exec err")
		log.Fatal(err)
	}
}
