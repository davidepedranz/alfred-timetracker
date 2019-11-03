package handler

import (
	"fmt"
	"time"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	aw "github.com/deanishe/awgo"
	"github.com/google/uuid"
)

func DoStart(wf *aw.Workflow, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	task := alfred.Task{ID: uuid.New().String(), Description: args[0], Start: time.Now()}

	// nolint:errcheck
	tasks, _ := alfred.LoadOngoingTasks(wf)
	tasks = append(tasks, task)

	if err := alfred.StoreOngoingTasks(wf, tasks); err != nil {
		return "", fmt.Errorf("cannot store the new task, please try again later 🙏 (%w)", err)
	}

	return "Task started, remember to stop it 😉", nil
}
