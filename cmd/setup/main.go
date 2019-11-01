package main

import (
	"context"
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/timetracker"
	"github.com/deanishe/awgo"
)

func main() {
	wf := aw.New()
	wf.Run(func() { run(wf) })
}

func run(wf *aw.Workflow) {
	token, err := timetracker.ReadToken(wf)
	if err != nil {
		timetracker.PrintError("Please authorize TimeTracker with `tt authorize` first 👀", err)
		return
	}

	clientID := wf.Config.Get(timetracker.ClientID)
	client, err := timetracker.NewClient(timetracker.NewConfig(clientID), token, context.Background())
	if err != nil {
		timetracker.PrintError("Something wrong happened, please try again later 🙏", err)
		return
	}

	id, err := client.CreateCalendar()
	if err != nil {
		timetracker.PrintError("Could not create the calendar, please try again later 🙏", err)
		return
	}

	if err := wf.Config.Set(timetracker.CalendarID, *id, false).Do(); err != nil {
		timetracker.PrintError("Cannot save the configuration in Alfred, please try again later 🙏", err)
		return
	}

	fmt.Print("Calendar created successfully 📅")
}
