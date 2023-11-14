package helper

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/sid04naik/send-email-go/config"
)


type EmailHelper struct {
	Config *config.Config

}

func (h *EmailHelper) SendEmail(fromEmail string, toEmail []string, subject, message string, signal chan struct{}) {
	messageBody := getMessageBody(subject, message, toEmail)
	// Set up authentication information.
 	smtp.SendMail(h.getAddress(),h.getAuth(), fromEmail, toEmail, messageBody)
	close(signal)
}

func (h *EmailHelper) getAuth() smtp.Auth {
	return smtp.PlainAuth("", h.Config.EmailConfig.AUTH.USER, h.Config.EmailConfig.AUTH.PASSWORD, h.Config.EmailConfig.HOST)

}

func (h *EmailHelper) getAddress() string {
	return fmt.Sprintf("%s:%d", h.Config.EmailConfig.HOST, h.Config.EmailConfig.PORT)
}

func getMessageBody(subject, message string, to []string) []byte {
	lineBreak := "\r\n";
	toEmail := fmt.Sprintf("To:%s%s",strings.Join(to, ","), lineBreak)   
	subject = fmt.Sprintf("Subject: %s%s", subject, lineBreak)
	messageBody := fmt.Sprintf("%s%s%s", toEmail, subject, message)	
	return []byte(messageBody)

}