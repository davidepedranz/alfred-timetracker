package handler

import (
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/alfred"
	aw "github.com/deanishe/awgo"
	"strings"
)

// TODO: commands to:
//  - forget a past task (forget)
//  - clear recent tasks (clear)

// nolint:unparam // we are forced to return a constant value to maintain the Handler signature
func DoFilterPast(wf *aw.Workflow, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	tasks, err := alfred.LoadPastTasks(wf)
	if err != nil {
		return "", fmt.Errorf("cannot load the past tasks, please try again later 🙏 (%w)", err)
	}

	description := strings.TrimSpace(args[0])
	if description == "" {
		wf.
			NewItem("<please provide a description>").
			Subtitle("Start a new task").
			Arg(description).
			Valid(false)
	} else {
		wf.
			NewItem(description).
			Subtitle("Start a new task").
			Arg(description).
			Valid(true)
	}

	for _, task := range tasks {
		if strings.HasPrefix(strings.ToLower(task.Description), strings.ToLower(description)) {
			wf.
				NewItem(task.Description).
				Subtitle("Continue this task").
				Arg(task.Description).
				Valid(true)
		}
	}

	wf.SendFeedback()

	return "", nil
}
