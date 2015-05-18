package event

import (
	"net/url"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Atnd implements Api
// Atnd represents event data returned by ATND API.
type Atnd struct {
	result struct {
		Results_returned int
		Results_start    int
		Events           []struct {
			Event struct {
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
}

func (self *Atnd) Get(baseurl, keyword, nickname string) ([]Event, error) {
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
			for _, atnd := range self.result.Events {
				event := Event{
					Id:          atnd.Event.Event_id,
					Title:       atnd.Event.Title,
					Started_at:  format(atnd.Event.Started_at),
					Url:         atnd.Event.Event_url,
					Summary:     atnd.Event.Catch,
					Place:       atnd.Event.Address + "\n" + atnd.Event.Place,
					Description: atnd.Event.Description,
				}
				events = append(events, event)
			}
			time.Sleep(time.Duration(1 * time.Second))
		}
	}
	return events, nil
}
