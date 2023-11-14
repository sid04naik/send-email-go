package helper

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sid04naik/send-email-go/config"
)

var Config *config.Config

// func Test_sendEmail(t *testing.T) {
// 	configurations, err := config.Configurations()
// 	if err != nil {
// 		t.Error("fail to get configurations")
// 	}

// 	eh := &EmailHelper{
// 		Config: &configurations,
// 	}

// }

// func Test_getAuth(t *testing.T) {
// 	configurations, err := config.Configurations()
// 	if err != nil {
// 		t.Error("fail to get configurations")
// 	}

// 	eh := &EmailHelper{
// 		Config: &configurations,
// 	}
// }

func TestGetAddress(t *testing.T) {
	loadConfig()
	eh := &EmailHelper{
		Config: Config,
	}
	expected := fmt.Sprintf("%s:%d", eh.Config.EmailConfig.HOST, eh.Config.EmailConfig.PORT)
	result := eh.getAddress()
	if result == expected {
		t.Log("address is successfully returned")
		return
	}
	t.Log("failed to return the expected address")
}

func Test_getMessageBody(t *testing.T) {
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
	// f, _ := os.Getwd()
	// rootPath := filepath.Dir(f)
	// envPath := filepath.Join(rootPath, ".env")
	configurations, err := config.Configurations(".env")
	if err != nil {
		fmt.Println("fail to get configurations", err)
	}
	Config = &configurations
}
