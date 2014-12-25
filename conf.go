package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v1"
)

// Config represents configurations defined in a yaml file.
type Config struct {
	Keyword      string
	Nickname     string
	Place        []string
	ErrorToSlack bool `yaml:"error_to_slack"`
	Provider     map[string]struct {
		Url      string
		Color    string
		Interval uint32
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
	Logfile string
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
		log.Println("slack send interval must be 1 or bigger.")
		return false
	}
	for _, provider := range config.Provider {
		if provider.Interval < 60 {
			log.Println("event check interval must be 60 or bigger. Do NOT overload providers.")
			return false
		}
	}
	return true
}
