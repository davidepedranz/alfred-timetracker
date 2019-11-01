package main

import (
	"encoding/json"
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/timetracker"
	"github.com/deanishe/awgo"
)

func main() {
	wf := aw.New()
	wf.Run(func() { run(wf) })
}

func run(wf *aw.Workflow) {
	config := timetracker.NewConfig(wf.Config.GetString(timetracker.ClientID))

	token, err := timetracker.GetAccessToken(config)
	if err != nil {
		timetracker.PrintError("Cannot get an access token 😢", nil)
		return
	}

	b, _ := json.Marshal(token)
	if err := wf.Keychain.Set("token", string(b)); err != nil {
		timetracker.PrintError("Cannot store the token in the keychain 😢", nil)
		return
	}

	fmt.Print("Token stored successfully 😎")
}
