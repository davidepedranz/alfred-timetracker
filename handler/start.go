package handler

import (
	"fmt"
	"time"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	aw "github.com/deanishe/awgo"
	"github.com/google/uuid"
)

func DoStart(wf *aw.Workflow, args []string) {
	if len(args) != 1 {
		alfred.PrintError("Please provide some input 👀", nil)
		return
	}

	task := alfred.Task{ID: uuid.New().String(), Description: args[0], Start: time.Now()}

	// nolint:errcheck
	tasks, _ := alfred.LoadOngoingTasks(wf)
	tasks = append(tasks, task)

	if err := alfred.StoreOngoingTasks(wf, tasks); err != nil {
		alfred.PrintError("Cannot store the new task, please try again later 🙏", err)
		return
	}

	fmt.Print("Task started, remember to stop it 😉")
}
