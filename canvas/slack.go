package canvas

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
	// upload to slack
	return nil
}
