package handler

import (
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/deanishe/awgo"
)

func DoCancel(wf *aw.Workflow, args []string) {
	if len(args) != 1 {
		alfred.PrintError("Please provide some input 👀", nil)
		return
	}

	tasks, err := alfred.LoadOngoingTasks(wf)
	if err != nil {
		alfred.PrintError("Cannot load the ongoing tasks, please try again later 🙏", err)
		return
	}

	index := search(tasks, args[0])
	if index == -1 {
		alfred.PrintError("Cannot find the provided task, maybe it was already stopped? 🤨", err)
		return
	}

	remaining := append(tasks[:index], tasks[index+1:]...)
	if err := alfred.StoreOngoingTasks(wf, remaining); err != nil {
		alfred.PrintError("Cannot store the left tasks, please try again later 🙏", err)
		return
	}

	fmt.Print("Task canceled ❌")
}
