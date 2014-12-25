package main

import . "github.com/daikikohara/enotify-slack/event"

// Payload represents a payload sent to Slack.
// The values are sent to Slack via incoming-webhook.
// See - https://my.slack.com/services/new/incoming-webhook
type Payload struct {
	Channel      string       `json:"channel"`
	Username     string       `json:"username"`
	Text         string       `json:"text"`
	Icon_emoji   string       `json:"icon_emoji"`
	Unfurl_links bool         `json:"unfurl_links"`
	Attachments  []Attachment `json:"attachments"`
}

// Attachment is an attachment to Payload.
// The format is defined in Slack Api document.
// See - https://api.slack.com/docs/attachments
type Attachment struct {
	Fallback string  `json:"fallback"`
	Pretext  string  `json:"pretext"`
	Color    string  `json:"color"`
	Fields   []Field `json:"fields"`
}

// Field is a field to Attachment.
// Like Attachment, the format is defined in Slack Api document.
// see - https://api.slack.com/docs/attachments
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// ConstructSlackMessage constructs a message sent to Slack.
// Each message contains a detail of an event gotten by GetEvent function.
func ConstructSlackMessage(provider string, event *Event, config *Config) *Payload {
	payload := Payload{}
	fields := []Field{
		Field{
			Title: "Title",
			Value: event.Title,
			Short: config.Slack.Short,
		},
		Field{
			Title: "Time",
			Value: event.Started_at,
			Short: config.Slack.Short,
		},
		Field{
			Title: "Place",
			Value: event.Place,
			Short: config.Slack.Short,
		},
		Field{
			Title: "Url",
			Value: event.Url,
			Short: config.Slack.Short,
		},
		Field{
			Title: "Summary",
			Value: event.Summary,
			Short: config.Slack.Short,
		},
	}

	attachment := Attachment{
		Fallback: "New event arrived from " + provider,
		Pretext:  config.Slack.Pretext,
		Color:    config.Provider[provider].Color,
		Fields:   fields,
	}

	payload.Channel = config.Slack.Channel
	payload.Username = provider + "-bot"
	payload.Icon_emoji = ":" + provider + ":"
	payload.Unfurl_links = true
	payload.Text = ""
	payload.Attachments = []Attachment{attachment}
	return &payload
}

// ConstructSlackError constructs an error message sent to Slack.
func ConstructSlackError(msg string, channel string) *Payload {
	fields := []Field{
		Field{
			Title: "Detail",
			Value: msg,
		},
	}

	attachment := Attachment{
		Fallback: "Error occured on enotify-slack",
		Pretext:  "Error occured on enotify-slack",
		Color:    "#FF0000",
		Fields:   fields,
	}

	payload := Payload{}
	payload.Channel = channel
	payload.Username = "notify-error"
	payload.Icon_emoji = ":persevere:"
	payload.Unfurl_links = true
	payload.Attachments = []Attachment{attachment}
	payload.Text = ""

	return &payload
}
