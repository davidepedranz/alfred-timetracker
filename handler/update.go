package handler

import (
	"fmt"
	"log"

	aw "github.com/deanishe/awgo"
)

// nolint:unparam
func DoUpdate(wf *aw.Workflow, _ []string) (string, error) {
	log.Println("Checking for updates...")

	if err := wf.CheckForUpdate(); err != nil {
		return "", fmt.Errorf("unknown error during the update (%w)", err)
	}

	if wf.UpdateAvailable() {
		wf.Feedback.Clear()
		wf.
			NewItem("New version found 😎").
			Subtitle("Please press Enter to install...").
			Arg("install").
			Valid(true)
	} else {
		wf.
			NewItem("Congratulations 🎉").
			Subtitle("Your workflow is already up-to-date!").
			Valid(true)
	}

	wf.SendFeedback()

	return "", nil
}
