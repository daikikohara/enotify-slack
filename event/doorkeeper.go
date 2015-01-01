package event

import (
	"time"

	"github.com/mitchellh/mapstructure"
)

// Doorkeeper implements Api
// Doorkeeper represents event data returned by DoorKeeper API.
type Doorkeeper struct {
	result []struct {
		Event struct {
			Id          string
			Title       string
			Public_url  string
			Address     string
			Venue_name  string
			Starts_at   string
			Description string
		}
	}
}

func (self *Doorkeeper) Get(baseurl, keyword, nickname string) ([]Event, error) {
	t := time.Now().Format("2006-01-02")
	param := "since=" + t
	result, err := GetJson(baseurl + param)
	if err != nil {
		return nil, err
	}
	if err = mapstructure.WeakDecode(result, &self.result); err != nil {
		return nil, err
	}
	var events []Event
	for _, e := range self.result {
		event := Event{
			Id:         e.Event.Id,
			Title:      e.Event.Title,
			Started_at: format(e.Event.Starts_at),
			Url:        e.Event.Public_url,
			Summary:    e.Event.Description,
			Place:      e.Event.Address + "\n" + e.Event.Venue_name,
		}
		if !contains(event.Title, event.Summary, keyword) {
			continue
		}
		event.Summary = trim(event.Summary)
		events = append(events, event)
	}
	return events, nil
}
