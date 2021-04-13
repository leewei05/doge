package common

import (
	"time"

	"github.com/pkg/errors"
)

func IsValidArgs(args []string) error {
	if len(args) > 1 {
		return errors.New("can only add one task at a time")
	}

	return nil
}

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}
