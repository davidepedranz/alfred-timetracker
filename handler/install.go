package handler

import (
	"fmt"
	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/deanishe/awgo"
)

func DoInstall(wf *aw.Workflow, _ []string) {
	fmt.Print("Downloading update...")
	if err := wf.InstallUpdate(); err != nil {
		alfred.PrintError("Error while downloading the update", err)
	}
}
