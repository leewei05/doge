package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileEmpty(t *testing.T) {
	f, err := os.Create("tmp.txt")
	check(err)

	defer f.Close()

	file := isFileEmpty("tmp.txt")
	assert.Equal(t, true, file)
}

func TestFileNotEmpty(t *testing.T) {
	f, err := os.Create("tmp.txt")
	check(err)

	defer f.Close()
	b := "test context"
	_, err = f.WriteString(b)
	check(err)

	file := isFileEmpty("tmp.txt")
	assert.Equal(t, false, file)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
