package event_test

import (
	"testing"

	. "github.com/daikikohara/enotify-slack/event"
	"github.com/stretchr/testify/assert"
)

var validTestTable = []struct {
	event  Event
	place  []string
	taboo  string
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
		"taboo",
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
		"taboo",
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
		"taboo",
		false,
	},
	{ // contains taboo in title
		Event{
			Id:         "id",
			Title:      "title,taboo",
			Summary:    "summary",
			Url:        "url",
			Started_at: "9999-12-31 23:59",
			Place:      "Tokyo",
		},
		[]string{"Tokyo", "place"},
		"taboo",
		false,
	},
	{ // contains taboo in summary
		Event{
			Id:         "id",
			Title:      "title",
			Summary:    "summary,taboo",
			Url:        "url",
			Started_at: "9999-12-31 23:59",
			Place:      "Tokyo",
		},
		[]string{"Tokyo", "place"},
		"taboo",
		false,
	},
	{ // contains taboo in description
		Event{
			Id:          "id",
			Title:       "title",
			Summary:     "summary",
			Url:         "url",
			Started_at:  "9999-12-31 23:59",
			Place:       "Tokyo",
			Description: "description foo",
		},
		[]string{"Tokyo", "place"},
		"taboo,foo",
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
		"taboo",
		false,
	},
}

func TestIsValid(t *testing.T) {
	assert := assert.New(t)
	for _, testcase := range validTestTable {
		assert.Equal(testcase.result, testcase.event.IsValid(testcase.place, testcase.taboo), "Event.IsValid test failed")
	}
}
