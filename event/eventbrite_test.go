package event_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
)

func TestGetEventbrite(t *testing.T) {
	SetTimezone("Asia/Tokyo")
	// success
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"events": [{"name":{"text":"example#1","html":"DockerConSFHackathon"},"description":{"text":"summary123","html":""},"id":"123","url":"http://example.connpass.com/event/123/","start":{"timezone":"America/Los_Angeles","local":"2015-06-20T09:30:00","utc":"2015-09-30T10:00:00Z"},"venue":{"address":{"address_1":"address123","address_2":null,"city":"place123","region":"CA","postal_code":"94103","country":"US","latitude":37.7850596,"longitude":-122.40419989999998},"resource_uri":"https://www.eventbriteapi.com/v3/venues/9687712/","id":"9687712","name":"SanFranciscoMarriottMarquis","latitude":"37.7850596","longitude":"-122.40419989999998"}}]}`)
	}))
	defer tsSuccess.Close()
	eventbrite := new(Eventbrite)
	events, err := eventbrite.Get(tsSuccess.URL+"?", keyword, nickname, place)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(eventbriteok, events[0]) {
		t.Error("eventbrite event is not ok!")
	}

	// different timezone
	SetTimezone("America/Los_Angeles")
	events, err = eventbrite.Get(tsSuccess.URL+"?", keyword, nickname, place)
	if err != nil {
		t.Error(err.Error())
	}
	e := eventbriteok
	e.Started_at = "2015-09-30 02:00"
	if !reflect.DeepEqual(e, events[0]) {
		// daylight saving
		e.Started_at = "2015-09-30 03:00"
		if !reflect.DeepEqual(e, events[0]) {
			t.Error("eventbrite event is not ok!")
		}
	}

	// invalid url
	events, err = eventbrite.Get("", keyword, nickname, place)
	if err == nil {
		t.Error(err.Error())
	}
}
