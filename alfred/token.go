package alfred

import (
	"encoding/json"
	"fmt"

	aw "github.com/deanishe/awgo"
	"golang.org/x/oauth2"
)

const tokenKey = "token"

func GetToken(wf *aw.Workflow) (*oauth2.Token, error) {
	var err error

	raw, err := wf.Keychain.Get(tokenKey)
	if err != nil {
		return nil, fmt.Errorf("token not found in the keychain")
	}

	token := new(oauth2.Token)
	if err := json.Unmarshal([]byte(raw), token); err != nil {
		return nil, fmt.Errorf("cannot parse the token in the keychain")
	}

	return token, nil
}

func SetToken(wf *aw.Workflow, token string) error {
	return wf.Keychain.Set(tokenKey, token)
}

func RemoveToken(wf *aw.Workflow) error {
	return wf.Keychain.Delete(tokenKey)
}
