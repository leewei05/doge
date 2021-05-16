package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArgsIsValid(t *testing.T) {
	args := []string{}

	err := IsValidArgs(args)
	assert.Error(t, err)

	args = append(args, "task#123")
	err = IsValidArgs(args)
	assert.NoError(t, err)

	args = append(args, "task#456")
	err = IsValidArgs(args)
	assert.Error(t, err)
}

// TODO: test current data with mock
func TestCurrentDate(t *testing.T) {
	//mockClock := clock.NewMock()

	//expect := mockClock.Now().Format("2006-01-02")
	//assert.Equal(t, expect, currentDate())
}
