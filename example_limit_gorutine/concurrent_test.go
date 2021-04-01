package example_limit_gorutine_test

import (
	"github.com/hoyeonUM/golang-example/example_limit_gorutine"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCrawler(t *testing.T) {
	t.Run("given concurrent count 10 and total count 100 when increase number 10 and wait time is 300ms then over execute time ge 1500ms", func(t *testing.T) {
		startTime := time.Now()
		example_limit_gorutine.Run(2, 100)
		endTime := time.Now()
		assert.GreaterOrEqual(t, endTime.Sub(startTime).Milliseconds(), int64(1500))
	})
}
