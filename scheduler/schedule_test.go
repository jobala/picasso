package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScheduler_ActionIsPerformedNTimes(t *testing.T) {
	var count = 0
	action := func() {
		count += 1
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	Do(action, 1*time.Millisecond, ctx)
	assert.Equal(t, 5, count)
}
