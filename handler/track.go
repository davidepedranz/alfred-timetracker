package handler

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/davidepedranz/alfred-timetracker/alfred"
	"github.com/davidepedranz/alfred-timetracker/calendar"
	aw "github.com/deanishe/awgo"
)

const durationRegex = `([1-9]\d*)\s*([smh])`

type event struct {
	description string
	start       time.Time
	end         time.Time
}

func DoTrack(wf *aw.Workflow, args []string) (string, error) {
	log.Println(args)

	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	e := parseEvent(args[0], time.Now())
	if e == nil {
		return "", fmt.Errorf("cannot recognize the input, please checkout the documentation for examples 🙏")
	}

	token, err := alfred.GetToken(wf)
	if err != nil {
		return "", fmt.Errorf("please authorize with `tt authorize` first 👀 (%w)", err)
	}

	calendarID := alfred.GetCalendarID(wf)
	if calendarID == "" {
		return "", fmt.Errorf("please setup your tracking calendar with `tt setup` first 👀")
	}

	clientID := alfred.GetClientID(wf)
	clientSecret := alfred.GetClientSecret(wf)

	client, err := calendar.NewClient(context.Background(), calendar.NewConfig(clientID, clientSecret), token)

	if err != nil {
		return "", fmt.Errorf("something wrong happened, please try again later 🙏 (%w)", err)
	}

	if err := client.InsertEvent(calendarID, e.description, &e.start, &e.end); err != nil {
		return "", fmt.Errorf("something wrong happened, please try again later 🙏 (%w)", err)
	}

	// nolint:goconst // we are forced to return a constant value to maintain the Handler signature
	return "Stored in your calendar 📅", nil
}

func parseEvent(raw string, now time.Time) *event {
	parsers := []func(string, time.Time) *event{parseDurationMessage, parseMessageDuration}

	for _, parser := range parsers {
		event := parser(raw, now)
		if event != nil {
			return event
		}
	}

	return nil
}

func parseDurationMessage(raw string, now time.Time) *event {
	pattern := regexp.MustCompile(`^\s*` + durationRegex + `\s+(\S.*)$`)
	match := pattern.FindStringSubmatch(raw)
	if match == nil {
		return nil
	}

	duration := int64(mustParseInt(match[1]))
	unit := match[2]

	return &event{
		description: strings.TrimSpace(match[3]),
		start:       now.Add(-time.Duration(duration) * toDuration(unit)),
		end:         now,
	}
}

func parseMessageDuration(raw string, now time.Time) *event {
	pattern := regexp.MustCompile(`^\s*(\S.*)\s+` + durationRegex + `$`)
	match := pattern.FindStringSubmatch(raw)
	if match == nil {
		return nil
	}

	duration := int64(mustParseInt(match[2]))
	unit := match[3]

	return &event{
		description: strings.TrimSpace(match[1]),
		start:       now.Add(-time.Duration(duration) * toDuration(unit)),
		end:         now,
	}
}

func toDuration(unit string) time.Duration {
	units := map[string]time.Duration{
		"s": time.Second,
		"m": time.Minute,
		"h": time.Hour,
	}
	return units[unit]
}

func mustParseInt(raw string) int {
	if raw == "" {
		return 0
	}
	result, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return result
}
