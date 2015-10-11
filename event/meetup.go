package event

import (
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Meetup implements Api
// Meetup represents event data returned by Meetup.
type Meetup struct {
	result struct {
		Results []struct {
			Id          string
			Name        string
			Event_url   string
			Time        int64
			Description string
			Venue       struct {
				City      string
				Address_1 string
			}
		}
	}
}

func (self *Meetup) Get(baseurl, keyword, nickname string, places []string) ([]Event, error) {
	var events []Event
	for _, param := range strings.Split(keyword, ",") {
		query := "topic=" + param
		result, err := GetJson(baseurl + query)
		if err != nil {
			return nil, err
		}
		if err = mapstructure.WeakDecode(result, &self.result); err != nil {
			return nil, err
		}
		for _, e := range self.result.Results {
			event := Event{
				Id:          e.Id,
				Title:       e.Name,
				Started_at:  formatEpoch(e.Time),
				Url:         e.Event_url,
				Place:       e.Venue.Address_1 + "\n" + e.Venue.City,
				Description: parse(e.Description),
			}
			event.Summary = trim(event.Description)
			events = append(events, event)
		}
		time.Sleep(time.Duration(1 * time.Second))
	}
	return events, nil
}
