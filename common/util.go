package common

import (
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	todoFile = "todo.txt"
)

func IsValidArgs(args []string) error {
	if len(args) > 1 {
		return errors.New("can only add one task at a time")
	}

	return nil
}

func currentDate() string {
	return time.Now().Format("2006-01-02")
}

func Path2Todo() (s string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	// refactor with string join
	pathTodo := home + "/.doge/" + currentDate() + "/" + todoFile

	return pathTodo, nil
}

func CreateDefaultTodoFile() error {
	todoFile, err := Path2Todo()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(todoFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	return nil
}

func CheckTodoDir() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	dogeDir := home + "/.doge/" + currentDate() + "/"

	if _, err := os.Stat(dogeDir); os.IsNotExist(err) {
		err := os.Mkdir(dogeDir, 0755)
		if err != nil {
			cobra.CheckErr(err)
		}
	}
}
