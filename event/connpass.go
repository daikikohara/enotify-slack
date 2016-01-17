package event

import (
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Connpass implements Api
// Connpass represents event data returned by Connpass.
type Connpass struct {
	result struct {
		Results_returned int
		Events           []struct {
			Event_id    string
			Title       string
			Catch       string
			Event_url   string
			Started_at  string
			Address     string
			Place       string
			Description string
		}
	}
}

func (self *Connpass) Get(baseurl, keyword, nickname string, places []string) ([]Event, error) {
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
			for _, e := range self.result.Events {
				event := Event{
					Id:          e.Event_id,
					Title:       e.Title,
					Started_at:  format(e.Started_at),
					Url:         e.Event_url,
					Summary:     e.Catch,
					Place:       e.Address + "\n" + e.Place,
					Description: parse(e.Description),
				}
				if event.Summary == "" {
					event.Summary = trim(event.Description)
				}
				events = append(events, event)
			}
			time.Sleep(time.Duration(1 * time.Second))
		}
	}
	return events, nil
}
