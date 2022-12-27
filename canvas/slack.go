package canvas

import (
	"fmt"
	"net/http"
)

type SlackCanvas struct {
	channel string
}

func Slack() *SlackCanvas {
	return &SlackCanvas{
		channel: "art",
	}
}

const SLACK_TOKEN = ""

func (s *SlackCanvas) Draw(art, title, artist string) error {
	img, _ := http.Get(art)
	req, _ := http.NewRequest(http.MethodPost, "https://slack.com/api/files.upload", img.Body)

	req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", SLACK_TOKEN))

	client := http.Client{}
	client.Do(req)

	return nil
}
