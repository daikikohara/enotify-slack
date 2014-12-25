package event

import (
	"strings"
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
		if !self.contains(event.Title, event.Summary, keyword) {
			continue
		}
		if len(event.Summary) > 100 {
			event.Summary = event.Summary[:100] + "..."
		}
		events = append(events, event)
	}
	return events, nil
}

// contains checks if Doorkeeper event contains keywords specified in the conf file
// Doorkeeper api does not take keywords as the parameter, so whether or not keywords are included has to be checked after getting the response.
func (self *Doorkeeper) contains(title, description, keyword string) bool {
	for _, q := range strings.Split(keyword, ",") {
		if strings.Contains(strings.ToLower(title), strings.ToLower(q)) || strings.Contains(strings.ToLower(description), strings.ToLower(q)) {
			return true
		}
	}
	return false
}
