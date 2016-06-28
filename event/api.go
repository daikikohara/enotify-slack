package event

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Api is an interface to get a slice of Event
type Api interface {
	// Get gets a slice of Event using API depending on api providers.
	Get(baseurl, keyword, nickname string, places []string) ([]Event, error)
}

// timeFormat holds time format.
// Currently it only supports "2006-01-02 15:04" style format.
const timeFormat = "2006-01-02 15:04"

// timezone holds timezone such as "Asia/Tokyo"
// This can be set via configuration file.
var timezone *time.Location

// SetTimezone sets timezone specified in the configuration file.
func SetTimezone(t string) {
	l, err := time.LoadLocation(t)
	if err != nil {
		panic(err.Error())
	}
	timezone = l
}

// GetApi is a factory function.
// GetApi returns an implementation of Api which gets actual events provided by each event provider.
func GetApi(provider string) Api {
	switch provider {
	case "doorkeeper":
		return new(Doorkeeper)
	case "atnd":
		return new(Atnd)
	case "connpass":
		return new(Connpass)
	case "zusaar":
		return new(Zusaar)
	case "strtacademy":
		return new(Strtacademy)
	case "meetup":
		return new(Meetup)
	case "eventbrite":
		return new(Eventbrite)
	default:
		log.Panic("Invalid api name:" + provider + "\ncheck conf file.")
	}
	return nil
}

// GetJson sends get request to the url passed by the argument and returns json-formatted event data.
func GetJson(url string) (interface{}, error) {
	client := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
