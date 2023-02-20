package alfred

import aw "github.com/deanishe/awgo"

const (
	clientID     = "client_id"
	clientSecret = "client_secret"
	calendarID   = "calendar_id"
)

func GetClientID(wf *aw.Workflow) string {
	return wf.Config.GetString(clientID)
}

func GetClientSecret(wf *aw.Workflow) string {
	return wf.Config.GetString(clientSecret)
}

func GetCalendarID(wf *aw.Workflow) string {
	return wf.Config.Get(calendarID)
}

func SetCalendarID(wf *aw.Workflow, id string) error {
	return wf.Config.Set(calendarID, id, false).Do()
}
