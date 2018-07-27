package zoutput

import (
	"testing"
	"time"
)

func Test_ProgressBar(t *testing.T) {
	c := make(chan int)

	go ProgressBar(c, 1234)

	progress := 0
	for {
		if progress > 1234 {
			close(c)
			break
		}

		c <- progress
		time.Sleep(1 * time.Second)
		progress += 120
	}
}
