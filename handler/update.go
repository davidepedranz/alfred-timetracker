package handler

import (
	"log"

	aw "github.com/deanishe/awgo"
)

func DoUpdate(wf *aw.Workflow, _ []string) {
	log.Println("Checking for updates...")

	if err := wf.CheckForUpdate(); err != nil {
		wf.FatalError(err)
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
}
