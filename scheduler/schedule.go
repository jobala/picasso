package scheduler

import (
	"context"
	"time"
)

type Action func()

func Do(action Action, dur time.Duration, ctx context.Context) {
	timer := time.NewTicker(dur)
	go func() {
		for {
			<-timer.C
			action()
		}
	}()

	<-ctx.Done()
}
