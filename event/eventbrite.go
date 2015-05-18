package event

import (
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Eventbrite implements Api
// Eventbrite represents event data returned by Eventbrite.
type Eventbrite struct {
	result struct {
		Events []struct {
			Id   string
			Name struct {
				Text string
			}
			Start struct {
				Utc string
			}
			Url         string
			Description struct {
				Text string
			}
			Venue struct {
				Address struct {
					Address_1 string
					City      string
				}
			}
		}
	}
}

func (self *Eventbrite) Get(baseurl, keyword, nickname string) ([]Event, error) {
	var events []Event
	for _, param := range strings.Split(keyword, ",") {
		query := "q=" + param
		result, err := GetJson(baseurl + query)
		if err != nil {
			return nil, err
		}
		if err = mapstructure.WeakDecode(result, &self.result); err != nil {
			return nil, err
		}
		for _, e := range self.result.Events {
			event := Event{
				Id:          e.Id,
				Title:       e.Name.Text,
				Started_at:  format(e.Start.Utc),
				Url:         e.Url,
				Summary:     e.Description.Text,
				Place:       e.Venue.Address.Address_1 + "\n" + e.Venue.Address.City,
				Description: e.Description.Text,
			}
			event.Summary = trim(event.Summary)
			events = append(events, event)
		}
		time.Sleep(time.Duration(1 * time.Second))
	}
	return events, nil
}
