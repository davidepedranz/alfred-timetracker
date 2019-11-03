package handler

import (
	"context"
	"fmt"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	aw "github.com/deanishe/awgo"
)

func DoSetup(wf *aw.Workflow, _ []string) (string, error) {
	token, err := alfred.ReadToken(wf)
	if err != nil {
		return "", fmt.Errorf("please authorize TimeTracker with `tt authorize` first 👀 (%w)", err)
	}

	clientID := wf.Config.Get(alfred.ClientID)
	client, err := calendar.NewClient(context.Background(), calendar.NewConfig(clientID), token)

	if err != nil {
		return "", fmt.Errorf("something wrong happened, please try again later 🙏 (%w)", err)
	}

	id, err := client.CreateCalendar()
	if err != nil {
		return "", fmt.Errorf("could not create the calendar, please try again later 🙏 (%w)", err)
	}

	if err := wf.Config.Set(alfred.CalendarID, *id, false).Do(); err != nil {
		return "", fmt.Errorf("cannot save the configuration in Alfred, please try again later 🙏 (%w)", err)
	}

	return "Calendar created successfully 📅", nil
}
