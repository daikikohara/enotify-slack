package event_test

import (
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
	"github.com/stretchr/testify/assert"
)

var validTestTable = []struct {
	event  Event
	place  []string
	result bool
}{
	{ //OK
		Event{
			Id:         "id",
			Title:      "title",
			Summary:    "summary",
			Url:        "url",
			Started_at: "9999-12-31 23:59",
			Place:      "Tokyo",
		},
		[]string{"Tokyo", "place"},
		true,
	},
	{ //invalid place
		Event{
			Id:         "id",
			Title:      "title",
			Summary:    "summary",
			Url:        "url",
			Started_at: "9999-12-31 23:59",
			Place:      "Kanagawa",
		},
		[]string{"Tokyo", "place"},
		false,
	},
	{ // invalid date
		Event{
			Id:         "id",
			Title:      "title",
			Summary:    "summary",
			Url:        "url",
			Started_at: "2013-12-31 23:59",
			Place:      "Tokyo",
		},
		[]string{"Tokyo", "place"},
		false,
	},
	{ // date format is invalid
		Event{
			Id:         "id",
			Title:      "title",
			Summary:    "summary",
			Url:        "url",
			Started_at: "12-31-9999 23:59",
			Place:      "Tokyo",
		},
		[]string{"Tokyo", "place"},
		false,
	},
}

func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	for _, testcase := range validTestTable {
		assert.Equal(testcase.result, testcase.event.IsValid(testcase.place), "Event.IsValid test failed")
	}
}
