package event

import (
	"fmt"

	"strings"
	"time"
)

// Event represents details of an event.
// Event is sent to Slack in this format.
// Atcual event type depends on event api providers, so provider-specific event type is defined in the each file and converted to this Event.
type Event struct {
	Id          string
	Title       string
	Summary     string
	Url         string
	Started_at  string
	Place       string
	Description string
}

// IsValid checks if the event is valid according to the place and date.
func (event *Event) IsValid(place []string, taboo string) bool {
	if !event.isValidPlace(place) {
		return false
	}
	if !event.isValidDate() {
		return false
	}
	if !event.isValidDesc(taboo) {
		return false
	}
	return true
}

// isValidPlace checks if the place of the event is valid.
// Valid palces are specified in the configuration file(conf.yml by default).
func (event *Event) isValidPlace(place []string) bool {
	for _, p := range place {
		if strings.Contains(event.Place, p) {
			return true
		}
	}
	return false
}

// isValidDate checks if the date of the event is valid.
func (event *Event) isValidDate() bool {
	t, err := time.Parse("2006-01-02 15:04", event.Started_at)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if t.After(time.Now()) {
		return true
	}
	return false
}

// isValidDesc checks if the tile, summary or description is valid.
func (event *Event) isValidDesc(taboo string) bool {
	for _, t := range strings.Split(taboo, ",") {
		if strings.Contains(event.Title, t) {
			return false
		}
		if strings.Contains(event.Summary, t) {
			return false
		}
		if strings.Contains(event.Description, t) {
			return false
		}
	}
	return true
}
