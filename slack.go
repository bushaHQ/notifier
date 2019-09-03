package notifier

import (
	"fmt"
	"runtime"

	"github.com/nlopes/slack"
)

// Slack a slack notifier
type Slack struct {
	product string
	channel string
	api     *slack.Client
}

// NewSlack returns a Notifier that notifies via slack.
func NewSlack(token, channelID, product string) Notifier {
	return Slack{
		api:     slack.New(token),
		product: product,   // product (service) sending the message
		channel: channelID, //default channel to send to
	}
}

// Error sends notification on errors that occur
// msg is the message to send. optChan is used to send to a
// channel different from the default channel
func (s Slack) Error(msg string, optChan ...string) {
	_, file, line, _ := runtime.Caller(1)
	text := fmt.Sprintf("*Error on `%s`*", s.product)
	channel := s.channel
	if len(optChan) > 0 {
		channel = optChan[0]
	}
	attachment := slack.Attachment{
		// Text:    "",
		Color: "#ff0000",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "message",
				Value: msg,
			},
			slack.AttachmentField{
				Title: "file",
				Value: fmt.Sprintf("%s:%d", file, line),
			},
		},
	}

	s.api.PostMessage(channel, slack.MsgOptionText(text, false), slack.MsgOptionAttachments(attachment))
}

// Info send notification about important info.
// just sends a message that give info.
// msg is the message to send. optChan is used to send to a
// channel different from the default channel.
func (s Slack) Info(msg string, optChan ...string) {
	text := fmt.Sprintf("*Important info `%s`*", s.product)
	channel := s.channel
	if len(optChan) > 0 {
		channel = optChan[0]
	}
	attachment := slack.Attachment{
		// Text:    "",
		Color: "#0000ff",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "message",
				Value: msg,
			},
		},
	}

	s.api.PostMessage(channel, slack.MsgOptionText(text, false), slack.MsgOptionAttachments(attachment))
}
