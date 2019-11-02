package handler

import (
	"fmt"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	aw "github.com/deanishe/awgo"
)

func DoDeauthorize(wf *aw.Workflow, _ []string) {
	token, err := alfred.ReadToken(wf)
	if err != nil {
		alfred.PrintError("TimeTracker already deauthorized 👀", err)
		return
	}

	if err := calendar.RevokeToken(token); err != nil {
		alfred.PrintError("Error while trying to deauthorize TimeTracker, please try again later 🙏", err)
	}

	// nolint:errcheck
	_ = alfred.RemoveToken(wf)

	fmt.Print("TimeTracker deauthorized successfully 😎")
}
