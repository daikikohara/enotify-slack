package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"
)

// Config represents configurations defined in a yaml file.
type Config struct {
	Keyword  string
	Nickname string
	Taboo    string
	Place    []string
	Provider map[string]*struct {
		Url      string
		Color    string
		Interval uint32
		Token    string
	}
	Slack struct {
		Pretext  string
		Url      string
		Channel  string
		Short    bool
		Interval uint32
	}
	Boltdb struct {
		Dbfile     string
		Bucketname string
	}
	Logfile      string
	Timezone     string
	ErrorToSlack bool `yaml:"error_to_slack"`
}

// defConf holds default values for Config
var defConf = Config{
	Provider: map[string]*struct {
		Url      string
		Color    string
		Interval uint32
		Token    string
	}{
		"atnd": {
			Url:      "https://api.atnd.org/events/?count=50&format=json&",
			Color:    "#000000",
			Interval: 3600,
		},
		"connpass": {
			Url:      "http://connpass.com/api/v1/event/?order=3&count=50&",
			Color:    "#D00000",
			Interval: 3600,
		},
		"doorkeeper": {
			Url:      "http://api.doorkeeper.jp/events/?sort=published_at&locale=ja&",
			Color:    "#00D0A0",
			Interval: 3600,
		},
		"eventbrite": {
			Url:      "https://www.eventbriteapi.com/v3/events/search/?token=",
			Color:    "#FFaa33",
			Interval: 3600,
		},
		"meetup": {
			Url:      "https://api.meetup.com/2/open_events?key=",
			Color:    "#F00000",
			Interval: 3600,
		},
		"partake": {
			Url:      "http://partake.in/api/event/search?sortOrder=createdAt&maxNum=50&query=",
			Color:    "#AAAAAA",
			Interval: 3600,
		},
		"strtacademy": {
			Url:      "http://www.street-academy.com/api/v1/events?page=",
			Color:    "#009900",
			Interval: 3600,
		},
		"zusaar": {
			Url:      "http://www.zusaar.com/api/event/?count=50&",
			Color:    "#0000FF",
			Interval: 3600,
		},
	},
	Slack: struct {
		Pretext  string
		Url      string
		Channel  string
		Short    bool
		Interval uint32
	}{
		Pretext:  "New Event Arrived!",
		Url:      "",
		Channel:  "#event-notify",
		Short:    false,
		Interval: 3,
	},
	Boltdb: struct {
		Dbfile     string
		Bucketname string
	}{
		Dbfile:     "enotify-slack.db",
		Bucketname: "enotify-slack",
	},
	Logfile:      "enotify-slack.log",
	Timezone:     "America/Los_Angeles",
	ErrorToSlack: false,
}

// NewConfig constructs Config using a yaml file passed as the argument.
func NewConfig(yml string) *Config {
	buf, err := ioutil.ReadFile(yml)
	if err != nil {
		log.Panicln(err)
	}
	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		log.Panicln(err)
	}
	if !config.isValid() {
		log.Panicln("conf file contains invalid value.")
	}
	return &config
}

// isValid checks if Config is valid.
// It currently only checks the interval to avoid overloading api providers.
func (config *Config) isValid() bool {
	if config.Slack.Interval < 1 {
		config.Slack.Interval = defConf.Slack.Interval
	}
	if config.Slack.Url == "" {
		log.Println("Slack webhook URL is empty.")
		return false
	}
	if config.Slack.Pretext == "" {
		config.Slack.Pretext = defConf.Slack.Pretext
	}
	if config.Slack.Channel == "" {
		log.Println("Slack channel name is empty.")
		return false
	}
	if config.Boltdb.Dbfile == "" {
		config.Boltdb.Dbfile = defConf.Boltdb.Dbfile
	}
	if config.Boltdb.Bucketname == "" {
		config.Boltdb.Bucketname = defConf.Boltdb.Bucketname
	}
	if config.Logfile == "" {
		config.Logfile = defConf.Logfile
	}
	if config.Timezone == "" {
		config.Timezone = defConf.Timezone
	}
	for name, provider := range config.Provider {
		if provider == nil {
			provider = &struct {
				Url      string
				Color    string
				Interval uint32
				Token    string
			}{
				Url:      defConf.Provider[name].Url,
				Color:    defConf.Provider[name].Color,
				Interval: defConf.Provider[name].Interval,
				Token:    "",
			}
		}
		if provider.Interval == 0 {
			provider.Interval = defConf.Provider[name].Interval
		} else if provider.Interval < 60 {
			log.Println("event check interval must be 60 or bigger. Do NOT overload providers.")
			return false
		}
		if provider.Url == "" {
			provider.Url = defConf.Provider[name].Url
		}
		if provider.Color == "" {
			provider.Color = defConf.Provider[name].Color
		}
		if name == "meetup" || name == "eventbrite" {
			if provider.Token == "" {
				log.Println("token is empty for" + name)
				return false
			}
			provider.Url = provider.Url + provider.Token + "&"
		}
		config.Provider[name] = provider
	}
	return true
}
