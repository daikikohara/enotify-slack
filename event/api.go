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
	Get(baseurl, keyword, nickname string) ([]Event, error)
}

const timeFormat = "2006-01-02 15:04"

// GetApi is a factory function.
// GetApi returns an implementation of Api which gets actual events provided by each event provider.
func GetApi(provider string) Api {
	switch provider {
	case "doorkeeper":
		return new(Doorkeeper)
	case "partake":
		return new(Partake)
	case "atnd":
		return new(Atnd)
	case "connpass":
		return new(Connpass)
	case "zusaar":
		return new(Zusaar)
	default:
		log.Fatalln("Invalid api name:" + provider + "\ncheck conf file.")
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

// format formats date format for the common format used for Slack message.
func format(date string) string {
	t, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		return date
	}
	return t.In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format(timeFormat)
}
