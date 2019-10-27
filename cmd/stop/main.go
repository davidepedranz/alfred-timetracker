package main

import (
	"context"
	"fmt"
	"github.com/deanishe/awgo"
	"time"
	"timetracker/timetracker"
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

	token, err := timetracker.ReadToken(wf)
	if err != nil {
		timetracker.PrintError("Please authorize TimeTracker with `tt authorize` first 👀", err)
		return
	}

	calendarID := wf.Config.Get(timetracker.CalendarID)
	if calendarID == "" {
		timetracker.PrintError("Please setup your tracking calendar with `tt setup` first 👀", err)
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

	clientID := wf.Config.Get(timetracker.ClientID)
	client, err := timetracker.NewClient(timetracker.NewConfig(clientID), token, context.Background())
	if err != nil {
		timetracker.PrintError("Something wrong happened, please try again later 🙏", err)
		return
	}

	task := tasks[index]
	now := time.Now()
	if err := client.InsertEvent(calendarID, task.Description, &task.Start, &now); err != nil {
		timetracker.PrintError("Something wrong happened, please try again later 🙏", err)
		return
	}

	fmt.Print("Stored in your calendar 📅")
}

func search(tasks []timetracker.Task, id string) int {
	for index, task := range tasks {
		if task.ID == id {
			return index
		}
	}
	return -1
}
