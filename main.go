package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jobala/picasso/canvas"
	"github.com/jobala/picasso/painter"
	"github.com/jobala/picasso/scheduler"
)

func main() {
	slack := canvas.Slack("C04DM14LK2N")
	painter := painter.New()

	painter.GetInspiration()

	scheduler.Do(func() {
		err := painter.PaintOn(slack)
		if err != nil {
			fmt.Println(fmt.Errorf("Couldn't paint: %v", err))
			os.Exit(1)
		}
	}, 24*time.Hour, context.Background())

}
