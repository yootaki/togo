package main

import (
	"os"

	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
}
