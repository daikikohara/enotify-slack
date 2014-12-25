package main_test

import (
	"testing"

	. "github.com/daikikohara/enotify-slack"
	"github.com/stretchr/testify/assert"
)

func TestConstructSlackMessage(t *testing.T) {

	//payload := ConstructSlackMessage("error message", "#test-channel")
	//	t.FailNow()

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
