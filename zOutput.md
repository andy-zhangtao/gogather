# zoutput
--
    import "github.com/andy-zhangtao/gogather/zoutput"


## Usage

```go
const (
	DOT = iota
	EQUAL
	STAR
	RANDOM
	ARROWS
)
```

#### func  ScrollBar

```go
func ScrollBar(c chan int, kind int)
```
ScrollBar 滚动条 指示当前程序正在运行, 不表示进度信息 当前外部逻辑运行完毕后，close(c)之后, ScrollBar会自行退出

##### Example

```go

    c := make(chan int)
    go ScrollBar(c, DOT)
    time.Sleep(5 * time.Second)
    close(c)

```
