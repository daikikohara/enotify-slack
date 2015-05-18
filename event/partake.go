package event

import (
	"net/url"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Partake implements Api
// Partake represents event data returned by Partake API.
type Partake struct {
	result struct {
		Result string
		Events []struct {
			Id          string
			Title       string
			Address     string
			Place       string
			BeginDate   string
			Summary     string
			Description string
		}
	}
}

func (self *Partake) Get(baseurl, keyword, nickname string) ([]Event, error) {
	var events []Event
	for _, param := range strings.Split(keyword, ",") {
		result, err := GetJson(baseurl + url.QueryEscape(param))
		if err != nil {
			return nil, err
		}
		if err = mapstructure.WeakDecode(result, &self.result); err != nil {
			return nil, err
		}
		for _, e := range self.result.Events {
			event := Event{
				Id:          e.Id,
				Title:       e.Title,
				Started_at:  e.BeginDate,
				Url:         "http://partake.in/events/" + e.Id,
				Summary:     e.Summary,
				Place:       e.Address + "\n" + e.Place,
				Description: e.Description,
			}
			events = append(events, event)
		}
		time.Sleep(time.Duration(1 * time.Second))
	}
	return events, nil
}
