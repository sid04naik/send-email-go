package main

import (
	"fmt"

	"github.com/sid04naik/send-email-go/config"
	"github.com/sid04naik/send-email-go/helper"
)

var Config *config.Config 

func main() {
	emailConfig := Config.EmailConfig
	toEmailAddress := []string{
		"gotest@mailinator.com",
	}
	var message string = "This is test email"
	var subject string = "This is email Subject"
	fromEmail := emailConfig.AUTH.USER

	var signal = make(chan struct{})

	emailHelper := helper.EmailHelper{
		Config: Config,
	}
	fmt.Println("Send Email Initiated")
	go emailHelper.SendEmail(fromEmail, toEmailAddress, subject, message, signal)
	<-signal
	fmt.Println("Email Send Successfully")
}

func init() {
	
	configuration, err := config.Configurations()
	Config = &configuration
	if err != nil {
		panic("error loading configurations")
	}
}