package mail

import (
	"Gin/until"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net/smtp"
)

func SendEmail(title, content string, receiver, cc []string) (err error) {

	defer func() {
		if err := recover(); err != nil {
			until.LogObj().WithFields(logrus.Fields{
				"func": "Run",
			}).Error("获取参数：", err)
			fmt.Println("发送邮件失败")
			return
		}

	}()

	SMTPHost := viper.GetString("mail.host")
	SMTPPort := viper.GetString("mail.port")
	SMTPUsername := viper.GetString("mail.user")
	SMTPPassword := viper.GetString("mail.password")

	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHost)
	e := &email.Email{
		From:    "发送者名字<" + SMTPUsername + ">",
		To:      receiver,
		Cc:      cc,
		Subject: title,

		HTML: []byte(content),
	}

	err = e.Send(SMTPHost+":"+SMTPPort, auth)

	if err != nil {
		log.Fatal(err)
	}
	return err
}
