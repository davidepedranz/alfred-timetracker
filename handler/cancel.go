package handler

import (
	"fmt"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	aw "github.com/deanishe/awgo"
)

func DoCancel(wf *aw.Workflow, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	tasks, err := alfred.LoadOngoingTasks(wf)
	if err != nil {
		return "", fmt.Errorf("cannot load the ongoing tasks, please try again later 🙏 (%w)", err)
	}

	index := search(tasks, args[0])
	if index == -1 {
		return "", fmt.Errorf("cannot find the provided task, maybe it was already stopped? 🤨")
	}

	remaining := append(tasks[:index], tasks[index+1:]...)
	if err := alfred.StoreOngoingTasks(wf, remaining); err != nil {
		return "", fmt.Errorf("cannot store the left tasks, please try again later 🙏 (%w)", err)
	}

	return "Task canceled ❌", nil
}
