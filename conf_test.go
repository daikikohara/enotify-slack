package main_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/daikikohara/enotify-slack"
	"github.com/stretchr/testify/assert"
)

var (
	validConf = "keyword: keyword1,keyword2\nnickname: person1,person2\nplace:\n  - Tokyo\n  - Kanagawa\nerror_to_slack: true\nprovider:\n  connpass:\n    url:   http://connpass.com/api/v1/event/?order=3&count=50&\n    color: \"#D00000\"\n    interval: 300\nslack:\n  pretext: slackpretext\n  url:      https://test.slack.com/services/hooks/incoming-webhook?token=tokenname\n  channel: \"#notify-channel\"\n  short:  false\n  interval: 3\nboltdb:\n  dbfile: it-event.db\n  bucketname: it-event\nlogfile: ./enotify-slack.log\ntimezone: Asia/Tokyo\n"
	// slice for invalid conf contents
	invalidConf = []string{
		// invalid slack interval data
		"keyword: keyword1,keyword2\nnickname: person1,person2\nplace:\n  - Tokyo\n  - Kanagawa\nerror_to_slack: true\nprovider:\n  connpass:\n    url:   http://connpass.com/api/v1/event/?order=3&count=50&\n    color: #D00000\n    interval: 0\nslack:\n  pretext: slackpretext\n  url:      https://test.slack.com/services/hooks/incoming-webhook?token=tokenname\n  channel: #notify-channel\n  short:  false\n  interval: 3\nboltdb:\n  dbfile: it-event.db\n  bucketname: it-event\nlogfile: ./enotify-slack.log\n",
		// invalid provider interval data
		"keyword: keyword1,keyword2\nnickname: person1,person2\nplace:\n  - Tokyo\n  - Kanagawa\nerror_to_slack: true\nprovider:\n  connpass:\n    url:   http://connpass.com/api/v1/event/?order=3&count=50&\n    color: #D00000\n    interval: 30\nslack:\n  pretext: slackpretext\n  url:      https://test.slack.com/services/hooks/incoming-webhook?token=tokenname\n  channel: #notify-channel\n  short:  false\n  interval: 0\nboltdb:\n  dbfile: it-event.db\n  bucketname: it-event\nlogfile: ./enotify-slack.log\n",
	}
)

func TestNewConfig(t *testing.T) {
	// initialize conf file
	file, err := ioutil.TempFile(os.TempDir(), "enotify-slack-config.yml")
	if err != nil {
		t.Error(err.Error())
	}
	defer os.Remove(file.Name())
	ioutil.WriteFile(file.Name(), []byte(validConf), 0644)
	conf := NewConfig(file.Name())

	// assert contents of valid conf
	assert := assert.New(t)
	assert.Equal("./enotify-slack.log", conf.Logfile, "conf parse error: logfile")
	assert.Equal("it-event", conf.Boltdb.Bucketname, "conf parse error: boltdb bucket name")
	assert.Equal("it-event.db", conf.Boltdb.Dbfile, "conf parse error: boltdb dbfile")
	assert.Equal(uint(0x3), conf.Slack.Interval, "conf parse error: slack interval")
	assert.False(conf.Slack.Short, "conf parse error: slack short")
	assert.Equal("#notify-channel", conf.Slack.Channel, "conf parse error: slack channel")
	assert.Equal("https://test.slack.com/services/hooks/incoming-webhook?token=tokenname", conf.Slack.Url, "conf parse error: slack url")
	assert.Equal("slackpretext", conf.Slack.Pretext, "conf parse error: slack pretext")
	assert.Equal(uint(0x12c), conf.Provider["connpass"].Interval, "conf parse error: provider interval")
	assert.Equal("#D00000", conf.Provider["connpass"].Color, "conf parse error: provider color")
	assert.Equal("http://connpass.com/api/v1/event/?order=3&count=50&", conf.Provider["connpass"].Url, "conf parse error: provider url")
	assert.Equal("Tokyo", conf.Place[0], "conf parse error: place[0]")
	assert.Equal("Kanagawa", conf.Place[1], "conf parse error: place[1]")
	assert.Equal("keyword1,keyword2", conf.Keyword, "conf parse error: keyword")
	assert.Equal("person1,person2", conf.Nickname, "conf parse error: nickname")
	assert.True(conf.ErrorToSlack, "conf parse error: error_to_slack")

	// function to test invalid config
	fn := func(f string) {
		defer func() {
			if r := recover(); r == nil {
				t.Fail()
			}
		}()
		NewConfig(f)
	}
	for _, s := range invalidConf {
		ioutil.WriteFile(file.Name(), []byte(s), 0644)
		fn(file.Name())
	}
}
