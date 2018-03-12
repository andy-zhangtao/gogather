package tools

import (
	"github.com/kataras/go-mailer"
	"errors"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/3/12.
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

// SendEmail 发送邮件到指定邮箱 content 邮件内容 addr 对方邮箱
func (this *Email) SendEmail() error {

	if this.Port == 0 {
		this.Port = 587
	}

	if this.Host == "" || this.Username == "" || this.Password == "" {
		return errors.New("Host Or UserName Or PassWord is NULL")
	}

	cfg := mailer.Config{
		Host:     this.Host,
		Username: this.Username,
		Password: this.Password,
		Port:     this.Port,
	}

	mailService := mailer.New(cfg)

	return mailService.Send(this.Header, this.Content, this.Dest...)
}
