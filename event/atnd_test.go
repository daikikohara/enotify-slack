package event_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
)

func TestGetAtnd(t *testing.T) {
	SetTimezone("Asia/Tokyo")
	// success
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"events":[{"event":{"event_id":123,"title":"example#1","catch":"summary123","description":"summary123","event_url":"http://example.connpass.com/event/123/","started_at":"2015-09-30T19:00:00.000+09:00","ended_at":null,"url":"http://not-this-one.com","limit":2147483647,"address":"address123","place":"place123","lat":"0.0","lon":"0.0","owner_id":123,"owner_nickname":"nickname1","owner_twitter_id":"nickname1","accepted":19,"waiting":0,"updated_at":"2018-09-24T13:51:01.000+09:00"}}]}`)
	}))
	defer tsSuccess.Close()
	atnd := new(Atnd)
	events, err := atnd.Get(tsSuccess.URL+"?", keyword, nickname)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(eventok, events[0]) {
		t.Error("atnd return value is unexpected. expected and actual are")
		t.Error(eventok)
		t.Error(events[0])
		t.FailNow()
	}

	// invalid url
	events, err = atnd.Get("", keyword, nickname)
	if err == nil {
		t.Error(err.Error())
	}

	// invalid time format
	tsInvalidTime := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"events":[{"event":{"event_id":123,"title":"example#1","catch":"summary123","description":"summary123","event_url":"http://example.connpass.com/event/123/","started_at":"20150930T190000.0000900","ended_at":null,"url":"http://not-this-one.com","limit":2147483647,"address":"address123","place":"place123","lat":"0.0","lon":"0.0","owner_id":123,"owner_nickname":"nickname1","owner_twitter_id":"nickname1","accepted":19,"waiting":0,"updated_at":"2018-09-24T13:51:01.000+09:00"}}]}`)
	}))
	defer tsInvalidTime.Close()
	events, err = atnd.Get(tsInvalidTime.URL+"?", keyword, nickname)
	if err != nil {
		t.FailNow()
	}
	if !reflect.DeepEqual(eventInvalidTime, events[0]) {
		t.Error("atnd return value is unexpected. expected and actual are")
		t.Error(eventInvalidTime)
		t.Error(events[0])
		t.FailNow()
	}
}
