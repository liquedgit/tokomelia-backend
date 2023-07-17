package service

import (
	"github.com/liquedgit/tokoMeLia/helper"
	"net/smtp"
)

func SendMail(msg string, receiverMail string) {
	smtpServer := helper.GoDotEnvVariables("SMTP_HOST")
	smtpPort := helper.GoDotEnvVariables("SMTP_PORT")
	senderMail := helper.GoDotEnvVariables("SENDER_MAIL")
	senderPassword := helper.GoDotEnvVariables("SENDER_MAIL_PASSWORD")
	receiver := []string{receiverMail}

	//body := []byte(msg)
	body := "From: " + senderMail + "\n" +
		"To: " + receiverMail + "\n" +
		"Subject: " + "Verify Email" + "\n\n" +
		msg
	auth := smtp.PlainAuth("", senderMail, senderPassword, smtpServer)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, senderMail, receiver, []byte(body))
	if err != nil {
		panic(err)
	}
	return
}
