package canvas

import (
	"fmt"
)

type SlackCanvas struct {
}

func Slack() *SlackCanvas {
	fmt.Println("Providing in slack")
	return &SlackCanvas{}
}

func (s *SlackCanvas) Draw(art, title, artist string) {
	fmt.Println(art)
}
