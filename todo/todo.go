package todo

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	// ID is a unique id for todo
	ID uuid.UUID
	// Row indicates how the todo will be shown
	Row int
	// Status indicates the working status for todo
	Status int
	// Name is the name of a todo
	Name string
	// Description is the context of a todo
	Description string
	// CreateTime is the creation time of a todo
	CreateTime time.Time
	// LastUpdateTime is the last update time of a todo
	LastUpdateTime time.Time
}

const (
	// STATUS_ONGOING means the todo is working in progress
	STATUS_ONGOING = 1
	// STATUS_DONE means the todo is finished
	STATUS_DONE = 2
	// STATUS_CLOSE means the todo is closed
	STATUS_CLOSE = 3
)
