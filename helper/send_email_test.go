package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sid04naik/send-email-go/config"
)

var Config = config.Config{}

func TestGetAddress(t *testing.T) {
	loadConfig()
	eh := &EmailHelper{
		Config: &Config,
	}
	expected := fmt.Sprintf("%s:%d", eh.Config.EmailConfig.HOST, eh.Config.EmailConfig.PORT)
	result := eh.getAddress()
	if result == expected {
		t.Log("address is successfully returned")
		return
	}
	t.Log("failed to return the expected address")
}

func TestSendMail(t *testing.T) {
	loadConfig()
	// eh := &EmailHelper{
	// 	Config: &Config,
	// }
	// 	toEmailAddress := []string{
	// 	"gotester@mailinator.com",
	// }
	// var message string = "This is test email"
	// var subject string = "This is email Subject"
	// fromEmail := emailConfig.AUTH.USER
	// eh.SendEmail()
}

func TestGetAuth(t *testing.T) {
	loadConfig()
	eh := &EmailHelper{
		Config: &Config,
	}

	auth := eh.getAuth()
	if auth != nil {
		t.Logf("auth received")
	}
}

func TestGetMessageBody(t *testing.T) {
	var inputParams = map[string]string{
		"subject": "test subject",
		"body":    "test body",
	}

	byteData := getMessageBody(inputParams["subject"], inputParams["body"], []string{"testmail@malinator.com"})
	expectedType := reflect.TypeOf([]byte{}).Kind().String()

	if reflect.TypeOf(byteData).Kind().String() == expectedType {
		t.Logf("result and expected type is same")
		return
	}
	t.Errorf("result and expected type is not same")

}

func loadConfig() {
	f, _ := os.Getwd()
	rootPath := filepath.Dir(f)
	envPath := filepath.Join(rootPath, ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println("fail to load env", err)
	}
	Config.EmailConfig.HOST = os.Getenv("EMAIL_HOST")
	Config.EmailConfig.PORT, _ = strconv.Atoi(os.Getenv("EMAIL_PORT"))
	Config.EmailConfig.AUTH.USER = os.Getenv("EMAIL_USERNAME")
	Config.EmailConfig.AUTH.PASSWORD = os.Getenv("EMAIL_PASSWORD")
}
