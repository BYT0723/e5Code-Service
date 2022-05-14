package mailx

import (
	"log"
	"testing"
)

func TestEmail(t *testing.T) {
	dialer := NewDialer()
	mes := NewMessage(Admin, "1151713064@qq.com", VerifyTitle, GenBody(VerifyTemplate, "402010"))
	if err := dialer.DialAndSend(mes); err != nil {
		log.Fatal(err)
	}
}
