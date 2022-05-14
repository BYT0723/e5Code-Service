package mailx

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

const (
	Admin = "twang9739@163.com"

	VerifyTitle = "e5code verify"
)

const (
	VerifyTemplate = `
    <p>e5code 用户邮箱验证</p>
    <p>验证码: %s </p>
    <p>激活码10分钟内有效，请尽快登陆激活</p>
  `
)

func NewDialer() *gomail.Dialer {
	dialer := gomail.NewDialer("smtp.163.com", 25, "twang9739@163.com", "ZGGGDWYPVRGIZKYE")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return dialer
}

func NewMessage(from, to, title, body string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)
	return m
}

func GenBody(template string, arg ...interface{}) string {
	return fmt.Sprintf(template, arg...)
}
