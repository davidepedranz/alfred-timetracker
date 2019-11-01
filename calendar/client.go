package calendar

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"time"
)

type client struct {
	service *calendar.Service
}

func NewClient(config *oauth2.Config, token *oauth2.Token, ctx context.Context) (*client, error) {
	service, err := calendar.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return nil, fmt.Errorf("cannot instantiate the calendar service: %w", err)
	}
	return &client{service}, nil
}

func (c *client) CreateCalendar() (*string, error) {
	cal := &calendar.Calendar{
		Summary:     "Tracking",
		Description: "Calendar for time-tracking managed by the Alfred TimeTracker workflow.",
	}
	call := c.service.Calendars.Insert(cal)
	created, err := call.Do()
	if err != nil {
		return nil, err
	}
	id := created.Id
	return &id, nil
}

func (c *client) InsertEvent(calendarID string, summary string, start *time.Time, end *time.Time) error {
	call := c.service.Events.Insert(calendarID, &calendar.Event{
		Summary:               summary,
		Start:                 &calendar.EventDateTime{DateTime: start.Format(time.RFC3339)},
		End:                   &calendar.EventDateTime{DateTime: end.Format(time.RFC3339)},
		GuestsCanInviteOthers: &[]bool{false}[0],
		Transparency:          "transparent",
		Visibility:            "private",
	})
	call.ConferenceDataVersion(1)
	event, err := call.Do()
	_ = event
	return err
}