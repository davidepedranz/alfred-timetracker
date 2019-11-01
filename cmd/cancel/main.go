package main

import (
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/timetracker"
	"github.com/deanishe/awgo"
)

func main() {
	wf := aw.New()
	wf.Run(func() { run(wf) })
}

func run(wf *aw.Workflow) {
	args := wf.Args()
	if len(args) != 1 {
		timetracker.PrintError("Please provide some input 👀", nil)
		return
	}

	tasks, err := timetracker.LoadOngoingTasks(wf)
	if err != nil {
		timetracker.PrintError("Cannot load the ongoing tasks, please try again later 🙏", err)
		return
	}

	index := search(tasks, args[0])
	if index == -1 {
		timetracker.PrintError("Cannot find the provided task, maybe it was already stopped? 🤨", err)
		return
	}

	remaining := append(tasks[:index], tasks[index+1:]...)
	if err := timetracker.StoreOngoingTasks(wf, remaining); err != nil {
		timetracker.PrintError("Cannot store the left tasks, please try again later 🙏", err)
		return
	}

	fmt.Print("Task canceled ❌")
}

func search(tasks []timetracker.Task, id string) int {
	for index, task := range tasks {
		if task.ID == id {
			return index
		}
	}
	return -1
}
