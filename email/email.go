package feifanEmail

import (
	"crypto/tls"
	"feifanEmailCron/db"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
)

var (
	auth smtp.Auth
	smtpPort int
	smtpHost string
	smtpUserName string
)

func init() {
	smtpUserName = DbsEmail.GetConfig().Conf.Email.Account
	smtpPassword := DbsEmail.GetConfig().Conf.Email.Password
	smtpHost = DbsEmail.GetConfig().Conf.Email.Smtp
	smtpPort = DbsEmail.GetConfig().Conf.Email.Port
	auth = smtp.PlainAuth("", smtpUserName, smtpPassword, smtpHost)
}

//发送邮件
func SendMail(to []string, body string) {
	//规范协议
	from := "非凡专家"
	user := "feifanexp@jfinfo.com"
	subject := "每日订阅动态"
	content_type := "Content-Type: text/html; charset=UTF-8"
	//邮件消息
	msg := []byte(
		"To: " + strings.Join(to, ",") + "\r\n" +
			"From: " + from + "<" + user + ">\r\n" +
			"Subject: " + subject + "\r\n" +
			content_type + "\r\n\r\n" +
			body)
	err := smtp.SendMail("smtp.exmail.qq.com:" + strconv.Itoa(smtpPort), auth, user, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
	}
}

//通过tls层发送邮件（加密邮件，防止内容泄露）
func SendMailTLS(to []string, body string) error {
	//规范协议
	from := "非凡专家"
	user := "feifanexp@jfinfo.com"
	subject := "每日订阅动态"
	content_type := "Content-Type: text/html; charset=UTF-8"
	//邮件消息
	msg := []byte(
		"To: " + strings.Join(to, ",") + "\r\n" +
			"From: " + from + "<" + user + ">\r\n" +
			"Subject: " + subject + "\r\n" +
			content_type + "\r\n\r\n" +
			body)
	//定义 tls 配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify:true,
		ServerName: smtpHost,
	}
	//整合邮件地址带端口 ip:port
	addr := fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	//tls 拨号连接
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		panic(err)
	}
	//创建邮件客户端
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	//验证发送邮件账户权限
	if auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(auth); err != nil{
				panic(err)
			}
		}
	}
	//设置邮件客户端的发送账号
	if err = client.Mail(smtpUserName); err != nil {
		panic(err)
	}
	//设置发送到邮件的地址
	for _, addr := range to {
		if err = client.Rcpt(addr); err != nil{
			panic(err)
		}
	}
	//获得写入邮件内容的容器
	w, err := client.Data()
	if err != nil {
		panic(err)
	}
	//写入邮件
	_, err = w.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	//关闭w
	err = w.Close()
	if err != nil {
		panic(err)
	}
	//退出邮件客户端
	return client.Quit()
}


















