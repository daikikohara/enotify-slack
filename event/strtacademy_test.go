package event_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
)

func TestGetStracademy(t *testing.T) {
	SetTimezone("Asia/Tokyo")
	// success
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"events": [{"title":"example#1","event_id":123,"start_at":"2015-09-30T19:00:00+09:00","end_at":"2014-12-17T12:00:00.000Z","venue":"place123","address":"address123","details":"summary123","url":"http://example.connpass.com/event/123/"},{"title":"example#1","event_id":123,"start_at":"2015-09-30T19:00:00+09:00","end_at":"2014-12-17T12:00:00.000Z","venue":"place123","address":"address123","details":"summary123toolongtoolongtoolongtoolongtoolongtoolongsummary123toolongtoolongtoolongtoolongtoolongtoolong","url":"http://example.connpass.com/event/123/"},{"title":"example#1","event_id":123,"start_at":"2015-09-30T19:00:00+09:00","end_at":"2014-12-17T12:00:00.000Z","venue":"place123","address":"address123","details":"toolong","url":"http://example.connpass.com/event/123/"}]}`)
	}))
	defer tsSuccess.Close()
	strt := new(Strtacademy)
	events, err := strt.Get(tsSuccess.URL+"?", "summary", nickname)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(eventok, events[0]) {
		t.Error("Strtacademy return value is unexpected. expected and actual are")
		t.Error(eventok)
		t.Error(events[0])
	}
	if !reflect.DeepEqual(eventoklong, events[1]) {
		t.Error("Strtacademy return value is unexpected. expected and actual are")
		t.Error(eventoklong)
		t.Error(events[1])
	}
	if len(events) != 20 {
		// first 2 elements of tsSuccess. and strt.Get() loops over 10 pages, so 20 in total.
		fmt.Println("length does not match!")
		t.Fail()
	}

	// invalid url
	events, err = strt.Get("", keyword, nickname)
	if err == nil {
		t.Error(err.Error())
	}
}
