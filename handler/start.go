package handler

import (
	"fmt"
	"time"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	aw "github.com/deanishe/awgo"
	"github.com/google/uuid"
)

// TODO: remove job from history when doing "cancel"

func DoStart(wf *aw.Workflow, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	description := args[0]

	task := alfred.Task{ID: uuid.New().String(), Description: description, Start: time.Now()}
	if err := updateTasks(wf, task); err != nil {
		return "", err
	}

	// nolint:errcheck // we keep track of recent tasks on a best-effort basis
	_ = updatePastTasks(wf, description)

	return "Task started, remember to stop it 😉", nil
}

func updateTasks(wf *aw.Workflow, task alfred.Task) error {
	// nolint:errcheck // we ignore errors because the file will be overridden in the next step in case of errors
	tasks, _ := alfred.LoadOngoingTasks(wf)
	tasks = append(tasks, task)

	if err := alfred.StoreOngoingTasks(wf, tasks); err != nil {
		return fmt.Errorf("cannot store the new task, please try again later 🙏 (%w)", err)
	}

	return nil
}

func updatePastTasks(wf *aw.Workflow, task string) error {
	pastTasks, _ := alfred.LoadPastTasks(wf)
	pastTasks = append(pastTasks, alfred.PastTask{Description: task})

	// TODO: deduplicate, max size, max age

	return alfred.StorePastTasks(wf, pastTasks)
}
