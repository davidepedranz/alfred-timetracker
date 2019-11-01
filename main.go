package main

import (
	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/handler"
	"github.com/deanishe/awgo"
)

func main() {
	wf := aw.New()
	wf.Run(func() { run(wf) })
}

func run(wf *aw.Workflow) {
	args := wf.Args()
	if len(args) == 0 {
		// TODO: change error message, consider using native awgo method
		alfred.PrintError("Please provide some input 👀", nil)
		return
	}

	handlers := map[string]func(*aw.Workflow, []string){
		"authorize": handler.DoAuthorize,
		"cancel":    handler.DoCancel,
		"filter":    handler.DoFilter,
		"setup":     handler.DoSetup,
		"start":     handler.DoStart,
		"stop":      handler.DoStop,
	}

	cmd := args[0]
	if h, found := handlers[cmd]; !found {
		alfred.PrintError("Command not recognized", nil)
		return
	} else {
		h(wf, args[1:])
	}
}
