package zoutput

import "fmt"

func ProgressBar(c chan int, size int) {
	for {
		select {
		case v, ok := <-c:
			if !ok {
				fmt.Print("\033[G\033[1K")
				fmt.Printf("%0.2f%%", 100.00)
				return
			} else {
				fmt.Print("\033[G\033[1K")
				fmt.Printf("%0.2f%%", (float64(v)/float64(size))*100)
			}
		}
	}
}
