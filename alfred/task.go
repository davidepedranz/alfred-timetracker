package alfred

import (
	"fmt"
	"time"

	aw "github.com/deanishe/awgo"
)

const ongoingTasks = "ongoing_tasks.json"

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
}

func LoadOngoingTasks(wf *aw.Workflow) ([]Task, error) {
	// fallback load function doing nothing
	nop := func() (interface{}, error) {
		return []Task{}, nil
	}

	var tasks []Task
	if err := wf.Data.LoadOrStoreJSON(ongoingTasks,0, nop, &tasks); err != nil {
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
