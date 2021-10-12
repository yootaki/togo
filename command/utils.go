package command

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

//エラー処理が甘い気がする
func dbPath() string {
	home, err := homedir.Dir()
	if err != nil {
	}
	return home + "/togo.db"
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return true
}
