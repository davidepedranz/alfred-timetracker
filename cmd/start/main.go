package main

import (
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/timetracker"
	"github.com/deanishe/awgo"
	"github.com/google/uuid"
	"time"
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

	task := timetracker.Task{ID: uuid.New().String(), Description: args[0], Start: time.Now()}
	tasks, _ := timetracker.LoadOngoingTasks(wf)
	tasks = append(tasks, task)

	if err := timetracker.StoreOngoingTasks(wf, tasks); err != nil {
		timetracker.PrintError("Cannot store the new task, please try again later 🙏", err)
		return
	}

	fmt.Print("Task started, remember to stop it 😉")
}
