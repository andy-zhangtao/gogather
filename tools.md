# tools
--
    import "github.com/andy-zhangtao/gogather/tools"


## Usage

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
