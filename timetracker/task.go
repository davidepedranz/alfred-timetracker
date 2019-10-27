package timetracker

import (
	"time"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
}
