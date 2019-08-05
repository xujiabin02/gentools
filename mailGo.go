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

func Mail(sub, content, mailList, serverHost string, serverPort int, userName, passWord string) error {
	//host := "smtp.exmail.qq.com"
	//port := 465
	host := serverHost
	port := serverPort
	toEmail := mailList
	header := make(map[string]string)
	header["From"] = "ai_watch" + "<" + userName + ">"
	header["To"] = mailList
	header["Subject"] = sub
	header["Content-Type"] = "text/html; charset=UTF-8"
	body := content
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth(
		"",
		userName,
		passWord,
		host,
	)
	err := sendMailUsingTLS(
		fmt.Sprintf("%s:%d", host, port),
		auth,
		userName,
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
