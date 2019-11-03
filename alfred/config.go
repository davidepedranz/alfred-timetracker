package alfred

import aw "github.com/deanishe/awgo"

const (
	clientID   = "client_id"
	calendarID = "calendar_id"
)

func GetClientID(wf *aw.Workflow) string {
	return wf.Config.GetString(clientID)
}

func GetCalendarID(wf *aw.Workflow) string {
	return wf.Config.Get(calendarID)
}

func SetCalendarID(wf *aw.Workflow, id string) error {
	return wf.Config.Set(calendarID, id, false).Do()
}
