package zoutput

import (
	"testing"
	"time"
)

func Test_ScrollBar(t *testing.T) {
	c := make(chan int)
	go ScrollBar(c, DOT)
	time.Sleep(5 * time.Second)
	close(c)


	c = make(chan int)
	go ScrollBar(c, STAR)
	time.Sleep(5 * time.Second)
	close(c)

	c = make(chan int)
	go ScrollBar(c, EQUAL)
	time.Sleep(5 * time.Second)
	close(c)

	c = make(chan int)
	go ScrollBar(c, ARROWS)
	time.Sleep(5 * time.Second)
	close(c)

	c = make(chan int)
	go ScrollBar(c, RANDOM)
	time.Sleep(5 * time.Second)
	close(c)
}
