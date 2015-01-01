package event_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
	"github.com/stretchr/testify/assert"
)

const (
	keyword  = "key1"
	nickname = "nickname1"
)

var (
	eventok = Event{
		Id:         "123",
		Title:      "example#1",
		Summary:    "summary123",
		Url:        "http://example.connpass.com/event/123/",
		Started_at: "2015-09-30 19:00",
		Place:      "address123\nplace123",
	}
	eventoklong = Event{
		Id:         "123",
		Title:      "example#1",
		Summary:    "summary123toolongtoolongtoolongtoolongtoolongtoolo...",
		Url:        "http://example.connpass.com/event/123/",
		Started_at: "2015-09-30 19:00",
		Place:      "address123\nplace123",
	}
	eventInvalidTime = Event{
		Id:         "123",
		Title:      "example#1",
		Summary:    "summary123",
		Url:        "http://example.connpass.com/event/123/",
		Started_at: "20150930T190000.0000900",
		Place:      "address123\nplace123",
	}
)

var funcNameTable = []struct {
	providerName string
	functionName string
}{
	{"doorkeeper", "Doorkeeper"},
	{"atnd", "Atnd"},
	{"partake", "Partake"},
	{"connpass", "Connpass"},
	{"zusaar", "Zusaar"},
	{"strtacademy", "Strtacademy"},
}

func TestGetApi(t *testing.T) {
	assert := assert.New(t)
	for _, testcase := range funcNameTable {
		assert.Contains(reflect.TypeOf(GetApi(testcase.providerName)).Elem().Name(), testcase.functionName, "GetApi test failed for "+testcase.providerName)
	}
	//assert.Empty(GetApi("other"))
}

func TestGetJson(t *testing.T) {
	// case1: fail to get request
	if _, err := GetJson(""); err == nil {
		t.Error(err.Error())
	}

	// case2: fail to unmarshall json response
	tsFail := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, ``)
	}))
	defer tsFail.Close()
	if _, err := GetJson(tsFail.URL); err == nil {
		t.Error(err.Error())
	}

	// case3: success
	tsSuccess := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"event":"123"}`)
	}))
	defer tsSuccess.Close()
	result, err := GetJson(tsSuccess.URL)
	if err != nil {
		t.Error(err.Error())
	}
	if result.(map[string]interface{})["event"] != "123" {
		t.Fail()
	}
}
