package canvas

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

// Slack creates an instance of the slack canvas
func Slack(channelId string) *SlackCanvas {
	return &SlackCanvas{
		channel: channelId,
		client:  slack.New(os.Getenv("SLACK_TOKEN")),
	}
}

// Draw "draws" the art by uploading them to slack
func (s *SlackCanvas) Draw(art, title, artist string) error {
	img, err := http.Get(art)

	if err != nil {
		return errors.New("Failed to get the image")
	}

	f, err := s.client.UploadFile(slack.FileUploadParameters{
		Filename: fmt.Sprintf("%s, %s", artist, title),
		Filetype: "image/jpeg",
		Channels: []string{s.channel},
		Reader:   img.Body,
	})

	if err != nil {
		return fmt.Errorf("Failed to upload image, %v", err)
	}

	fmt.Println("Image uploaded successfully")
	fmt.Println(f.Title)

	return nil
}

type SlackCanvas struct {
	channel string
	client  *slack.Client
}

var SLACK_TOKEN = os.Getenv("SLACK_TOKEN")
