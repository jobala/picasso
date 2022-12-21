package main

import (
	"github.com/jobala/picasso/canvas"
	"github.com/jobala/picasso/painter"
)

func main() {
	slack := canvas.Slack()
	painter := painter.NewPainter()

	painter.PaintOn(slack)
}
