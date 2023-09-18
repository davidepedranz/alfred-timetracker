package alfred

import (
	"fmt"
	"time"

	aw "github.com/deanishe/awgo"
)

const ongoingTasks = "ongoing_tasks.json"
const pastTasks = "past_tasks.json"

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
}

type PastTask struct {
	Description string `json:"description"`
}

func LoadOngoingTasks(wf *aw.Workflow) ([]Task, error) {
	// fallback load function doing nothing
	nop := func() (interface{}, error) {
		return []Task{}, nil
	}

	var tasks []Task
	if err := wf.Data.LoadOrStoreJSON(ongoingTasks, 0, nop, &tasks); err != nil {
		return nil, fmt.Errorf("error loading the ongoing tasks: %w", err)
	}

	return tasks, nil
}

func StoreOngoingTasks(wf *aw.Workflow, tasks []Task) error {
	if err := wf.Data.StoreJSON(ongoingTasks, tasks); err != nil {
		return fmt.Errorf("error storing the ongoing tasks: %w", err)
	}

	return nil
}

func LoadPastTasks(wf *aw.Workflow) ([]PastTask, error) {
	// fallback load function doing nothing
	nop := func() (interface{}, error) {
		return []PastTask{}, nil
	}

	var tasks []PastTask
	if err := wf.Data.LoadOrStoreJSON(pastTasks, 0, nop, &tasks); err != nil {
		return nil, fmt.Errorf("error loading the past tasks: %w", err)
	}

	return tasks, nil
}

func StorePastTasks(wf *aw.Workflow, tasks []PastTask) error {
	if err := wf.Data.StoreJSON(pastTasks, tasks); err != nil {
		return fmt.Errorf("error storing the past tasks: %w", err)
	}

	return nil
}
