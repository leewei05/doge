/*
Copyright © 2021 Lee Wei <lee10202013@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	c "github.com/leewei05/doge/common"
	"github.com/leewei05/doge/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all TODOs from TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		todoFile, err := c.Path2Todo()
		cobra.CheckErr(err)

		if ok := isFileEmpty(todoFile); ok {
			fmt.Println("Empty Todo list")
			return
		}

		file, err := os.Open(todoFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			t := scanner.Text()
			var todo todo.Todo
			json.Unmarshal([]byte(t), &todo)
			fmt.Println(todo)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func isFileEmpty(file string) bool {
	fi, err := os.Stat(file)
	if err != nil {
		cobra.CheckErr(err)
	}

	if fi.Size() == 0 {
		return true
	}

	return false
}
