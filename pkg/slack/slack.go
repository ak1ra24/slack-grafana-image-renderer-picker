package slack

import (
	"github.com/slack-go/slack"
	"github.com/ak1ra24/slack-grafana-image-renderer-picker/pkg/grafana"
)

type Slack struct {
	token   string
	channel string
}

func NewSlack(token, channel string) *Slack {
	s := &Slack{}
	s.token = token
	s.channel = channel

	return s
}

func (s *Slack) PostImage(graph *grafana.Graph) error {

	api := slack.New(s.token)

	params := slack.FileUploadParameters{
		Reader: graph.Graph,
		Filename: graph.URL,
		Channels: []string{s.channel},
	}

	_, err := api.UploadFile(params)
	if err != nil {
		return err
	}

    return  nil
}
