package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	aw "github.com/deanishe/awgo"
)

func DoStop(wf *aw.Workflow, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	token, err := alfred.GetToken(wf)
	if err != nil {
		return "", fmt.Errorf("please authorize with `tt authorize` first 👀 (%w)", err)
	}

	calendarID := alfred.GetCalendarID(wf)
	if calendarID == "" {
		return "", fmt.Errorf("please setup your tracking calendar with `tt setup` first 👀")
	}

	tasks, err := alfred.LoadOngoingTasks(wf)
	if err != nil {
		return "", fmt.Errorf("cannot load the ongoing tasks, please try again later 🙏 (%w)", err)
	}

	index := search(tasks, args[0])
	if index == -1 {
		return "", fmt.Errorf("cannot find the provided task, maybe it was already stopped? 🤨")
	}

	// the next command will remove the task, so we save it here
	task := tasks[index]

	remaining := append(tasks[:index], tasks[index+1:]...)
	if err := alfred.StoreOngoingTasks(wf, remaining); err != nil {
		return "", fmt.Errorf("cannot store the left tasks, please try again later 🙏 (%w)", err)
	}

	clientID := alfred.GetClientID(wf)
	clientSecret := alfred.GetClientSecret(wf)

	client, err := calendar.NewClient(context.Background(), calendar.NewConfig(clientID, clientSecret), token)

	if err != nil {
		return "", fmt.Errorf("something wrong happened, please try again later 🙏 (%w)", err)
	}

	now := time.Now()
	if err := client.InsertEvent(calendarID, task.Description, &task.Start, &now); err != nil {
		return "", fmt.Errorf("something wrong happened, please try again later 🙏 (%w)", err)
	}

	// nolint:goconst // we are forced to return a constant value to maintain the Handler signature
	return "Stored in your calendar 📅", nil
}
