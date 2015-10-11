package event

import (
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
)

// alphabet is a regex for a string which contains only alphabets and spaces.
var alphabet = regexp.MustCompile(`(?i)^[a-z\s]+$`)

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
			Venue_id string
		}
	}
}

func (self *Eventbrite) Get(baseurl, keyword, nickname string, places []string) ([]Event, error) {
	var events []Event
	t := time.Now().Format("2006-01-02T15:04:05Z")
	for _, param := range strings.Split(keyword, ",") {
		for _, city := range places {
			c := city
			if !alphabet.MatchString(city) {
				c = cityMap[city]
				if c == "" {
					continue
				}
			}
			query := "q=" + param + "&sort_by=date" + "&start_date.range_start=" + t + "&venue.city=" + url.QueryEscape(c)
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
					Summary:     trim(e.Description.Text),
					Place:       city,
					Description: e.Description.Text,
				}
				events = append(events, event)
			}
			time.Sleep(time.Duration(1 * time.Second))
		}
	}
	return events, nil
}
