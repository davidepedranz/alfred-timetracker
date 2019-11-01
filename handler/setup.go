package handler

import (
	"context"
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	"github.com/deanishe/awgo"
)

func DoSetup(wf *aw.Workflow, _ []string) {
	token, err := alfred.ReadToken(wf)
	if err != nil {
		alfred.PrintError("Please authorize TimeTracker with `tt authorize` first 👀", err)
		return
	}

	clientID := wf.Config.Get(alfred.ClientID)
	client, err := calendar.NewClient(calendar.NewConfig(clientID), token, context.Background())
	if err != nil {
		alfred.PrintError("Something wrong happened, please try again later 🙏", err)
		return
	}

	id, err := client.CreateCalendar()
	if err != nil {
		alfred.PrintError("Could not create the calendar, please try again later 🙏", err)
		return
	}

	if err := wf.Config.Set(alfred.CalendarID, *id, false).Do(); err != nil {
		alfred.PrintError("Cannot save the configuration in Alfred, please try again later 🙏", err)
		return
	}

	fmt.Print("Calendar created successfully 📅")
}
