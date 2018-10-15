# znet
--
    import "github.com/andy-zhangtao/gogather/znet"


## Usage

#### func  ConvertToHex

```go
func ConvertToHex(ip string) (ipHex []byte, err error)
```
ConvertToHex 将IP地址转换为十六进制数据 例如将 127.0.0.1 转换为 7F 00 00 01

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
