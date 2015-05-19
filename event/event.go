package event

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"

	"strings"
	"time"
)

// Event represents details of an event.
// Event is sent to Slack in this format.
// Atcual event type depends on event api providers, so provider-specific event type is defined in the each file and converted to this Event.
type Event struct {
	Id          string
	Title       string
	Summary     string
	Url         string
	Started_at  string
	Place       string
	Description string
}

// IsValid checks if the event is valid according to the place and date.
func (event *Event) IsValid(place []string, taboo string) bool {
	if !event.isValidPlace(place) {
		return false
	}
	if !event.isValidDate() {
		return false
	}
	if !event.isValidDesc(taboo) {
		return false
	}
	return true
}

// isValidPlace checks if the place of the event is valid.
// Valid palces are specified in the configuration file(conf.yml by default).
func (event *Event) isValidPlace(place []string) bool {
	for _, p := range place {
		if strings.Contains(event.Place, p) {
			return true
		}
	}
	return false
}

// isValidDate checks if the date of the event is valid.
func (event *Event) isValidDate() bool {
	t, err := time.Parse("2006-01-02 15:04", event.Started_at)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	if t.After(time.Now()) {
		return true
	}
	return false
}

// isValidDesc checks if the tile, summary or description is valid.
func (event *Event) isValidDesc(taboo string) bool {
	for _, t := range strings.Split(taboo, ",") {
		if strings.Contains(strings.ToLower(event.Title), strings.ToLower(t)) {
			return false
		}
		if strings.Contains(strings.ToLower(event.Summary), strings.ToLower(t)) {
			return false
		}
		if strings.Contains(strings.ToLower(event.Description), strings.ToLower(t)) {
			return false
		}
	}
	return true
}

// contains checks if keyword is contained in either title or description.
func (event *Event) contains(keyword string) bool {
	for _, q := range strings.Split(keyword, ",") {
		if strings.Contains(strings.ToLower(event.Title), strings.ToLower(q)) ||
			strings.Contains(strings.ToLower(event.Summary), strings.ToLower(q)) ||
			strings.Contains(strings.ToLower(event.Description), strings.ToLower(q)) {
			return true
		}
	}
	return false
}

// trim returns first 100 characters of string passed by the argument.
func trim(s string) string {
	if len([]rune(s)) > 100 {
		return string([]rune(s)[:100]) + "..."
	}
	return s
}

// format formats date format for the common format used for Slack message.
func format(date string) string {
	t, err := time.Parse(time.RFC3339Nano, date)
	if err != nil {
		return date
	}
	return t.In(timezone).Format(timeFormat)
}

// formatEpoch formats time in milliseconds in epoch to the common format used for Slack message.
func formatEpoch(i int64) string {
	tm := time.Unix(i/1000, 0)
	return tm.In(timezone).Format(timeFormat)
}

// parse parses html to change it into plain text.
func parse(h string) string {
	doc, err := html.Parse(strings.NewReader(h))
	if err != nil {
		// just return original html in case of error
		return h
	}

	var buffer bytes.Buffer
	var f func(*html.Node, *bytes.Buffer)
	f = func(n *html.Node, buffer *bytes.Buffer) {
		if n.Type == html.TextNode {
			// this call might fail but ignore.
			buffer.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, buffer)
		}
	}
	f(doc, &buffer)
	return buffer.String()
}
