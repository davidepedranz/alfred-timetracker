package handler

import (
	"fmt"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	aw "github.com/deanishe/awgo"
)

func DoDeauthorize(wf *aw.Workflow, _ []string) (string, error) {
	token, err := alfred.GetToken(wf)
	if err != nil {
		return "", fmt.Errorf("already deauthorized 👀 (%w)", err)
	}

	if err := calendar.RevokeToken(token); err != nil {
		return "", fmt.Errorf("error during deauthorization, please try again later 🙏 (%w)", err)
	}

	// nolint:errcheck // removing the token from the keychain is best effort
	_ = alfred.RemoveToken(wf)

	return "TimeTracker deauthorized successfully 😎", nil
}
