package event_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
)

func TestGetDoorkeeper(t *testing.T) {
	// success
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[{"event":{"title":"example#1","id":123,"starts_at":"2015-09-30T10:00:00.000Z","ends_at":"2014-12-17T12:00:00.000Z","venue_name":"place123","address":"address123","lat":"34.6985045","long":"135.4902115","ticket_limit":15,"published_at":"2014-12-07T13:52:23.341Z","updated_at":"2014-12-07T13:52:23.343Z","description":"notcontain123","public_url":"http://example.connpass.com/event/123/","participants":0,"waitlisted":0,"group":{"id":234,"name":"","country_code":"JP","logo":"hogehoge.jpg","description":"hogehoge","public_url":"http://hogehoge/"}}},{"event":{"title":"example#1","id":123,"starts_at":"2015-09-30T10:00:00.000Z","ends_at":"2014-12-17T12:00:00.000Z","venue_name":"place123","address":"address123","lat":"34.6985045","long":"135.4902115","ticket_limit":15,"published_at":"2014-12-07T13:52:23.341Z","updated_at":"2014-12-07T13:52:23.343Z","description":"summary123","public_url":"http://example.connpass.com/event/123/","participants":0,"waitlisted":0,"group":{"id":234,"name":"","country_code":"JP","logo":"hogehoge.jpg","description":"hogehoge","public_url":"http://hogehoge/"}}}]`)
	}))
	defer tsSuccess.Close()
	dk := new(Doorkeeper)
	events, err := dk.Get(tsSuccess.URL+"?", "summary", nickname)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(eventok, events[0]) {
		t.Error("doorkeeper return value is unexpected. expected and actual are")
		t.Error(eventok)
		t.Error(events[0])
	}

	// invalid url
	events, err = dk.Get("", keyword, nickname)
	if err == nil {
		t.Error(err.Error())
	}

	// invalid time format
	tsInvalidTime := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[{"event":{"title":"example#1","id":123,"starts_at":"20150930T190000.0000900","ends_at":"2014-12-17T12:00:00.000Z","venue_name":"place123","address":"address123","lat":"34.6985045","long":"135.4902115","ticket_limit":15,"published_at":"2014-12-07T13:52:23.341Z","updated_at":"2014-12-07T13:52:23.343Z","description":"summary123","public_url":"http://example.connpass.com/event/123/","participants":0,"waitlisted":0,"group":{"id":234,"name":"","country_code":"JP","logo":"hogehoge.jpg","description":"hogehoge","public_url":"http://hogehoge/"}}}]`)
	}))
	defer tsInvalidTime.Close()
	events, err = dk.Get(tsInvalidTime.URL+"?", "summary", nickname)
	if err != nil {
		t.FailNow()
	}
	if !reflect.DeepEqual(eventInvalidTime, events[0]) {
		t.FailNow()
	}
}
