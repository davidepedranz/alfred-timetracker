package handler

import (
	"encoding/json"
	"fmt"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	aw "github.com/deanishe/awgo"
)

func DoAuthorize(wf *aw.Workflow, _ []string) (string, error) {
	clientID := alfred.GetClientID(wf)
	clientSecret := alfred.GetClientSecret(wf)

	config := calendar.NewConfig(clientID, clientSecret)

	token, err := calendar.GetToken(config)
	if err != nil {
		return "", fmt.Errorf("cannot get an access token 😢 (%w)", err)
	}

	b, err := json.Marshal(token)
	if err != nil {
		return "", fmt.Errorf("cannot serialize the token to JSON 😢 (%w)", err)
	}

	if err := alfred.SetToken(wf, string(b)); err != nil {
		return "", fmt.Errorf("cannot store the token in the keychain 😢 (%w)", err)
	}

	return "Token stored successfully 😎", nil
}
