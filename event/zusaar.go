package event

import (
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Zusaar implements Api
// Zusaar represents event data returned by Zusaar Api.
type Zusaar struct {
	result struct {
		Results_returned int
		Event            []struct {
			Event_id   string
			Title      string
			Catch      string
			Event_url  string
			Started_at string
			Address    string
			Place      string
		}
	}
}

func (self *Zusaar) Get(baseurl, keyword, nickname string) ([]Event, error) {
	var events []Event
	for _, param := range []string{"keyword_or=" + url.QueryEscape(keyword), "owner_nickname=" + nickname, "nickname=" + nickname} {
		for t := time.Now().Local(); t.Before(time.Now().Local().AddDate(0, 3, 0)); t = t.AddDate(0, 1, 0) {
			query := param + "&ym=" + t.Format("200601")
			result, err := GetJson(baseurl + query)
			if err != nil {
				return nil, err
			}
			if err = mapstructure.WeakDecode(result, &self.result); err != nil {
				return nil, err
			}
			for _, e := range self.result.Event {
				event := Event{
					Id:         e.Event_id,
					Title:      e.Title,
					Started_at: format(e.Started_at),
					Url:        e.Event_url,
					Summary:    e.Catch,
					Place:      e.Address + "\n" + e.Place,
				}
				events = append(events, event)
			}
			time.Sleep(time.Duration(1 * time.Second))
		}
	}
	return events, nil
}
