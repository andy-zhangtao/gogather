package zoutput

import (
	"fmt"
	"time"
)

const (
	DOT = iota
	EQUAL
	STAR
	RANDOM
	ARROWS
)

// ScrollBar 滚动条
// 指示当前程序正在运行, 不表示进度信息
// 当前外部逻辑运行完毕后，close(c)之后, ScrollBar会自行退出
/*
##### Example

```go
	c := make(chan int)
	go ScrollBar(c, DOT)
	time.Sleep(5 * time.Second)
	close(c)
```
*/
func ScrollBar(c chan int, kind int) {
	for {
		select {
		case _, ok := <-c:
			if !ok {
				fmt.Println()
				return
			}
		default:
			output(kind)
			time.Sleep(1 * time.Second)
		}
	}
}

func output(kind int) {
	switch kind {
	case DOT:
		fmt.Print(". ")
	case EQUAL:
		fmt.Print("= ")
	case STAR:
		fmt.Print("* ")
	case ARROWS:
		fmt.Print("> ")
	case RANDOM:
		fmt.Print("@ ")
	}
}
