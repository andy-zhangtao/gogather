# znet
--
    import "github.com/andy-zhangtao/gogather/znet"


## Usage

#### func  GetFreePort

```go
func GetFreePort() (int, error)
```
GetFreePort 获取当前空闲的端口

##### Example

```go

    port, err := GetFreePort()

```

#### func  LocallIP

```go
func LocallIP() (string, error)
```
LocallIP 获取本地IP地址
