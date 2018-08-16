# tools
--
    import "github.com/andy-zhangtao/gogather/tools"


## Usage

#### func  CopyDir

```go
func CopyDir(src, dst string) (err error)
```
CopyDir copies a dir from src to dst. src should be a full path. Also dst too.
If src is a file, then will invoke CopyFile. If src is a dir, then will copy all
the files it contains to dst. * ##### Example

```go package main

import (

    "github.com/andy-zhangtao/gogather/tools"
    "fmt"

)

func main() {

    err := tools.CopyDir("/Users/zhangtao/SourceCode/golang/go/src/temp/test", "/tmp/test")
    if err != nil {
    	fmt.Println(err)
    }

}

```

#### func  CopyFile

```go
func CopyFile(src, dst string) (err error)
```
CopyFile copies a file from src to dst. If src and dst files exist, and are the
same, then return success. Otherise, attempt to create a hard link between the
two files. If that fail, copy the file contents from src to dst.

#### func  LineCounter

```go
func LineCounter(r io.Reader) (int, error)
```
LineCounter 统计文件行数 r 通过os.Open获取的文件reader

#### type Email

```go
type Email struct {
	/*Host MailGun主机地址*/
	Host string `json:"host"`
	/*UserName MailGun用户名*/
	Username string `json:"user"`
	/*PassWord MailGun口令*/
	Password string `json:"passwd"`
	/*Port MailGun邮件发送端口*/
	Port int `json:"port"`
	/*Dest 目标邮件地址列表*/
	Dest []string `json:"dest"`
	/*Content 邮件内容*/
	Content string `json:"content"`
	/*Header 邮件主题*/
	Header string `json:"header"`
}
```

Write by zhangtao<ztao8607@gmail.com> . In 2018/3/12.

#### func (*Email) SendEmail

```go
func (this *Email) SendEmail() error
```
SendEmail 发送邮件到指定邮箱 content 邮件内容 addr 对方邮箱
