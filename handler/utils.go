package handler

import (
	"github.com/davidepedranz/alfred-timetracker/alfred"
)

func search(tasks []alfred.Task, id string) int {
	for index, task := range tasks {
		if task.ID == id {
			return index
		}
	}

	return -1
}
