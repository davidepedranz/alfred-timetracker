package alfred

import (
	"encoding/json"
	"fmt"
	"time"

	aw "github.com/deanishe/awgo"
	"golang.org/x/oauth2"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Start       time.Time `json:"start"`
}

func LoadOngoingTasks(wf *aw.Workflow) ([]Task, error) {
	var tasks []Task
	if err := wf.Data.LoadJSON(OngoingTasks, &tasks); err != nil {
		return nil, fmt.Errorf("error loading the ongoing tasks: %w", err)
	}

	return tasks, nil
}

func StoreOngoingTasks(wf *aw.Workflow, tasks []Task) error {
	if err := wf.Data.StoreJSON(OngoingTasks, tasks); err != nil {
		return fmt.Errorf("error storing the ongoing tasks: %w", err)
	}

	return nil
}

func ReadToken(wf *aw.Workflow) (*oauth2.Token, error) {
	var err error

	raw, err := wf.Keychain.Get(Token)
	if err != nil {
		return nil, fmt.Errorf("token not found in the keychain")
	}

	token := new(oauth2.Token)
	if err := json.Unmarshal([]byte(raw), token); err != nil {
		return nil, fmt.Errorf("cannot parse the token in the keychain")
	}

	return token, nil
}

func RemoveToken(wf *aw.Workflow) error {
	return wf.Keychain.Delete(Token)
}
