package per

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

//func Mail() {
func Mail(sub, content, mailList, serverHost string, serverPort int) error {
	//host := "smtp.exmail.qq.com"
	//port := 465
	host := serverHost
	port := serverPort
	email := "ai_watching@9fbank.com.cn"
	password := "lv7yFtoHh8yafbBnUup0"
	toEmail := mailList
	header := make(map[string]string)
	header["From"] = "ai_watch" + "<" + email + ">"
	header["To"] = mailList
	header["Subject"] = sub
	header["Content-Type"] = "text/html; charset=UTF-8"
	body := content
	//header["From"] = "ai_watch" + "<" + email + ">"
	//header["To"] = toEmail
	//header["Subject"] = "邮件标题"
	//header["Content-Type"] = "text/html; charset=UTF-8"
	//body := "我是一封电子邮件!golang发出. by:www.361way.com <运维之路>"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth(
		"",
		email,
		password,
		host,
	)
	err := sendMailUsingTLS(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		email,
		strings.Split(toEmail, ";"),
		[]byte(message),
	)
	if err != nil {
		return err
	}
	return nil
}

//return a smtp client
func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}
func someError(addr, err string) error {
	someA := errors.New(err + " ," + addr + " send failed")
	return someA
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {
	//create smtp client
	c, err := dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				//panic(err)
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			someP := someError(addr, err.Error())
			return someP
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

//func Mail(sub, content, mailList, serverHost string, serverPort int) (error, bool) {
//	var (
//	//	sub = `[mail test!]`
//	//	content = `Just test it !!~`
//	//	mailList = `xujiabin@9fbank.com.cn,xujiabin02@outlook.com`
//		me=`ai_watching@9fbank.com.cn`
//		contentType="text/html"
//	)
//	m:=gomail.NewMessage()
//	//mailL:=strings.Split(mailList, ",")
//	m.SetHeader("From", me)
//	m.SetHeader("To", strings.Split(mailList,",")...)
//	m.SetHeader("Subject", sub)
//	m.SetBody(contentType, content)
//	d:=gomail.NewDialer(serverHost, serverPort, "ai_watching@9fbank.com.cn", "lv7yFtoHh8yafbBnUup0")
//	// Send the email to Bob, Cora and Dan.
//	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
//	if err := d.DialAndSend(m); err != nil {
//		//godbg.Dbg(err)
//		panic(err)
//		return err, false
//	}
//	return nil, true
//}
