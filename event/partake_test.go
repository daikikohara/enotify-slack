package event_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
)

var (
	eventpartake = Event{
		Id:         "123",
		Title:      "example#1",
		Summary:    "summary123",
		Url:        "http://partake.in/events/123",
		Started_at: "2015-09-30 19:00",
		Place:      "address123\nplace123",
	}
)

func TestGetPartake(t *testing.T) {
	// success
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"events":[{"id":"123","title":"example#1","summary":"summary123","category":"computer","beginDate":"2015-09-30 19:00","beginDateText":"2014-12-21 13:00","beginDateTime":1419134400000,"endDate":"2014-12-21 17:00","endDateText":"2014-12-21 17:00","endDateTime":1419148800000,"eventDuration":"","url":"","place":"place123","address":"address123","description":"description123","hashTag":"#hashtag","ownerId":"abc","foreImageId":null,"backImageId":null,"passcode":null,"draft":false,"editorIds":[],"relatedEventIds":[],"createdAt":"2014-12-05 15:09","createdAtTime":1417759760446,"modifiedAt":"2014-12-05 15:50","modifiedAtTime":1417762235891,"revision":36}],"result":"ok"}`)
	}))
	defer tsSuccess.Close()
	partake := new(Partake)
	events, err := partake.Get(tsSuccess.URL+"?", keyword, nickname)
	if err != nil {
		t.Error(err.Error())
	}
	if !reflect.DeepEqual(eventpartake, events[0]) {
		t.FailNow()
	}

	// invalid url
	events, err = partake.Get("", keyword, nickname)
	if err == nil {
		t.Error(err.Error())
	}

}
