package handler

import (
	"fmt"
	"time"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	aw "github.com/deanishe/awgo"
)

// nolint:unparam // we are forced to return a constant value to maintain the Handler signature
func DoFilterOngoing(wf *aw.Workflow, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	tasks, err := alfred.LoadOngoingTasks(wf)
	if err != nil {
		return "", fmt.Errorf("cannot load the ongoing tasks, please try again later 🙏 (%w)", err)
	}

	if len(tasks) == 0 {
		icon := &aw.Icon{Value: "./warning.png"}
		wf.
			NewItem("No ongoing task found").
			Subtitle("Try to add some task...").
			Icon(icon)
	}

	for _, task := range tasks {
		duration := time.Since(task.Start)
		wf.
			NewItem(task.Description).
			Subtitle(formatDuration(duration)).
			Arg(task.ID).
			Valid(true)
	}

	wf.SendFeedback()

	return "", nil
}

func formatDuration(d time.Duration) string {
	rounded := d.Round(time.Minute)

	h := rounded / time.Hour
	m := (rounded - h*time.Hour) / time.Minute

	if h > 0 {
		return fmt.Sprintf("%d h and %d m", h, m)
	}

	return fmt.Sprintf("%d m", m)
}
