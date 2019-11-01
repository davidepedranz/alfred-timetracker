package handler

import (
	"encoding/json"
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	"github.com/deanishe/awgo"
)

func DoAuthorize(wf *aw.Workflow, _ []string) {
	config := calendar.NewConfig(wf.Config.GetString(alfred.ClientID))

	token, err := calendar.GetAccessToken(config)
	if err != nil {
		alfred.PrintError("Cannot get an access token 😢", nil)
		return
	}

	b, _ := json.Marshal(token)
	if err := wf.Keychain.Set("token", string(b)); err != nil {
		alfred.PrintError("Cannot store the token in the keychain 😢", nil)
		return
	}

	fmt.Print("Token stored successfully 😎")
}
