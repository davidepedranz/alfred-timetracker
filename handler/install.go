package handler

import (
	"fmt"

	aw "github.com/deanishe/awgo"
)

// nolint:unparam
func DoInstall(wf *aw.Workflow, _ []string) (string, error) {
	fmt.Print("Downloading update...")

	if err := wf.InstallUpdate(); err != nil {
		return "", fmt.Errorf("error while downloading the update (%w)", err)
	}

	return "", nil
}
