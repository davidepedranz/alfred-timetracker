package handler

import (
	"context"
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	"github.com/deanishe/awgo"
	"time"
)

func DoStop(wf *aw.Workflow, args []string) {
	if len(args) != 1 {
		alfred.PrintError("Please provide some input 👀", nil)
		return
	}

	token, err := alfred.ReadToken(wf)
	if err != nil {
		alfred.PrintError("Please authorize TimeTracker with `tt authorize` first 👀", err)
		return
	}

	calendarID := wf.Config.Get(alfred.CalendarID)
	if calendarID == "" {
		alfred.PrintError("Please setup your tracking calendar with `tt setup` first 👀", err)
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

	clientID := wf.Config.Get(alfred.ClientID)
	client, err := calendar.NewClient(calendar.NewConfig(clientID), token, context.Background())
	if err != nil {
		alfred.PrintError("Something wrong happened, please try again later 🙏", err)
		return
	}

	task := tasks[index]
	now := time.Now()
	if err := client.InsertEvent(calendarID, task.Description, &task.Start, &now); err != nil {
		alfred.PrintError("Something wrong happened, please try again later 🙏", err)
		return
	}

	fmt.Print("Stored in your calendar 📅")
}
