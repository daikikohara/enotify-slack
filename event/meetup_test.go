package event_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
)

func TestGetMeetup(t *testing.T) {
	SetTimezone("Asia/Tokyo")
	// success
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"results":[{"utc_offset":-25200000,"venue":{"country":"us","city":"place123","address_1":"address123","name":"Hacker Dojo (event room)","lon":-122.107498,"id":123,"state":"CA","lat":37.411819,"repinned":false},"rsvp_limit":140,"headcount":0,"visibility":"public","waitlist_count":75,"created":1430850935000,"maybe_rsvp_count":0,"description":"summary123","event_url":"http://example.connpass.com/event/123/","yes_rsvp_count":140,"name":"example#1","id":"123","time":1443607200000,"updated":1431026798000,"group":{"join_mode":"open","created":1413333080000,"name":"Docker Networking","group_lon":-122.1500015258789,"id":17629502,"urlname":"Docker-Networking","group_lat":37.439998626708984,"who":"Members"},"status":"upcoming"}]}`)
	}))
	defer tsSuccess.Close()
	meetup := new(Meetup)
	events, err := meetup.Get(tsSuccess.URL+"?", keyword, nickname, place)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(eventok, events[0]) {
		t.Error("meetup event is not ok!")
	}

	// different timezone
	SetTimezone("America/Los_Angeles")
	events, err = meetup.Get(tsSuccess.URL+"?", keyword, nickname, place)
	if err != nil {
		t.Error(err.Error())
	}
	e := eventok
	e.Started_at = "2015-09-30 02:00"
	if !reflect.DeepEqual(e, events[0]) {
		// daylight saving
		e.Started_at = "2015-09-30 03:00"
		if !reflect.DeepEqual(e, events[0]) {
			t.Error("meetup event is not ok!")
		}
	}

	// invalid url
	events, err = meetup.Get("", keyword, nickname, place)
	if err == nil {
		t.Error(err.Error())
	}
}
