package event

import (
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Strtacademy implements Api
// Strtacademy represents event data returned by street academy API.
type Strtacademy struct {
	result struct {
		Events []struct {
			Event_id string
			Title    string
			Details  string
			Url      string
			Start_at string
			Address  string
			Venue    string
		}
	}
}

func (self *Strtacademy) Get(baseurl, keyword, nickname string) ([]Event, error) {
	var events []Event
	for i := 1; i <= 10; i++ {
		result, err := GetJson(baseurl + strconv.Itoa(i))
		if err != nil {
			return nil, err
		}
		if err = mapstructure.WeakDecode(result, &self.result); err != nil {
			return nil, err
		}
		for _, strtacademy := range self.result.Events {
			event := Event{
				Id:         strtacademy.Event_id,
				Title:      strtacademy.Title,
				Started_at: format(strtacademy.Start_at),
				Url:        strtacademy.Url,
				Summary:    strtacademy.Details,
				Place:      strtacademy.Address + "\n" + strtacademy.Venue,
			}
			if !contains(event.Title, event.Summary, keyword) {
				continue
			}
			event.Summary = trim(event.Summary)
			events = append(events, event)
		}
		time.Sleep(time.Duration(1 * time.Second))
	}
	return events, nil
}
