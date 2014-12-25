package event_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
)

func TestGetConnpass(t *testing.T) {
	// success
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"events":[{"event_url": "http://example.connpass.com/event/123/", "event_type": "advertisement", "owner_nickname": "nickname1", "series": {"url": "http://example.connpass.com/", "id": 1, "title": "example"}, "updated_at": "2014-11-30T10:44:20+09:00", "lat": 35.646470900000, "started_at": "2015-09-30T19:00:00+09:00", "hash_tag": "example", "title": "example#1", "event_id": 123, "lon": 139.706581400000, "waiting": 0, "limit": 60, "owner_id": 8, "owner_display_name": "nickname1", "description": "description keyword", "accepted": 0, "ended_at": "2015-09-30T21:00:00+09:00", "place": "place123", "address": "address123", "catch": "summary123"}]}`)
	}))
	defer tsSuccess.Close()
	connpass := new(Connpass)
	events, err := connpass.Get(tsSuccess.URL+"?", keyword, nickname)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(eventok, events[0]) {
		t.Fail()
	}

	// invalid url
	events, err = connpass.Get("", keyword, nickname)
	if err == nil {
		t.Fail()
	}

	// invalid time format
	tsInvalidTime := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"events":[{"event_url": "http://example.connpass.com/event/123/", "event_type": "advertisement", "owner_nickname": "nickname1", "series": {"url": "http://example.connpass.com/", "id": 1, "title": "example"}, "updated_at": "2014-11-30T10:44:20+09:00", "lat": 35.646470900000, "started_at": "20150930T190000.0000900", "hash_tag": "example", "title": "example#1", "event_id": 123, "lon": 139.706581400000, "waiting": 0, "limit": 60, "owner_id": 8, "owner_display_name": "nickname1", "description": "description keyword", "accepted": 0, "ended_at": "2015-09-30T21:00:00+09:00", "place": "place123", "address": "address123", "catch": "summary123"}]}`)
	}))
	defer tsInvalidTime.Close()
	events, err = connpass.Get(tsInvalidTime.URL+"?", keyword, nickname)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(eventInvalidTime, events[0]) {
		t.FailNow()
	}

}
