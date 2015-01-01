package main_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	. "github.com/daikikohara/enotify-slack"
	. "github.com/daikikohara/enotify-slack/event"
	"github.com/stretchr/testify/assert"
)

var event = &Event{
	Id:         "id",
	Title:      "title",
	Summary:    "summary",
	Url:        "url",
	Started_at: "9999-12-31 23:59",
	Place:      "Tokyo",
}

var payloadok = &Payload{
	Channel:      "#notify-channel",
	Username:     "connpass-bot",
	Text:         "",
	Icon_emoji:   ":connpass:",
	Unfurl_links: true,
	Attachments: []Attachment{
		Attachment{
			Fallback: "New event arrived from connpass",
			Pretext:  "slackpretext",
			Color:    "#D00000",
			Fields: []Field{
				Field{
					Title: "Title",
					Value: "title",
					Short: false,
				},
				Field{
					Title: "Time",
					Value: "9999-12-31 23:59",
					Short: false,
				},
				Field{
					Title: "Place",
					Value: "Tokyo",
					Short: false,
				},
				Field{
					Title: "Url",
					Value: "url",
					Short: false,
				},
				Field{
					Title: "Summary",
					Value: "summary",
					Short: false,
				},
			},
		},
	},
}

func TestConstructSlackMessage(t *testing.T) {
	// initialize conf file
	file, err := ioutil.TempFile(os.TempDir(), "enotify-slack-config.yml")
	if err != nil {
		t.Error(err.Error())
	}
	defer os.Remove(file.Name())
	ioutil.WriteFile(file.Name(), []byte(validConf), 0644)
	conf := NewConfig(file.Name())
	payload := ConstructSlackMessage("connpass", event, conf)

	if !reflect.DeepEqual(payload, payloadok) {
		t.Log("payload doesn't match")
		t.Log(payload)
		t.Log(payloadok)
		t.Fail()
	}
}

func TestConstructSlackError(t *testing.T) {

	payload := ConstructSlackError("error message", "#test-channel")

	// assert contents of slack payload
	assert := assert.New(t)
	assert.Equal("#test-channel", payload.Channel, "ConstructSlackError parse error: channel")
	assert.Equal(":persevere:", payload.Icon_emoji, "ConstructSlackError parse error: icon_emoji")
	assert.Equal("", payload.Text, "ConstructSlackError parse error: text")
	assert.True(payload.Unfurl_links, "ConstructSlackError parse error: unfurl_links")
	assert.Equal("notify-error", payload.Username, "ConstructSlackError parse error: user name")
	assert.Equal("Detail", payload.Attachments[0].Fields[0].Title, "ConstructSlackError parse error: attachment/field/title")
	assert.Equal("error message", payload.Attachments[0].Fields[0].Value, "ConstructSlackError parse error: attachment/field/value")
}
