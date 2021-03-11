package iyzypay_Samples

import (
	"fmt"
	"github.com/icobani/iyzipay-go"
	"github.com/icobani/iyzipay-go/Request"
	"log"
	"testing"
)

func TestBin(t *testing.T) {
	var Pay, _ = Iyzipay.Iyzipay{}.New(
		"<api-key>",
		"<api-secret>",
	)

	binReq :=
		Request.BinCheckRequest{
			Locale:         "tr",
			ConversationId: "123456789",
			BinNumber:      "540709",
		}

	binRes, _ := Pay.BinCheck(&binReq)
	fmt.Println(binRes)

	if binRes.BankName != "Garanti Bankası" {
		t.Errorf("It is supprise :) We are waiting %s", binRes.BankName)
	} else {
		log.Printf("%#v\n", binRes)
	}
}
