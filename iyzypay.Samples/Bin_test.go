package iyzypay_Samples

import (
	"testing"
	"github.com/icobani/iyzipay-go"
	"github.com/icobani/iyzipay-go/Request"
	"log"
)

func TestBin(t *testing.T) {
	var Pay, _ = Iyzipay.Iyzipay{}.New(
		"sandbox-lwggMKTv5mSwC1Z4e6Zay4n6HKUOhqOX",
		"sandbox-Ok8ZKrca4DWyEjGX0GHq8zwOCpne0rEQ",
		"https://sandbox-api.iyzipay.com",
	)

	binReq :=
		Request.BinCheckRequest{
			Locale:         "tr",
			ConversationId: "123456789",
			BinNumber:      "540709",
		}

	binRes, _ := Pay.BinCheck(&binReq)


	if binRes.BankName != "Garanti Bankası" {
		t.Errorf("It is supprise :) We are waiting %s",binRes.BankName)
	} else {
		log.Println("Status", binRes.Status)
		log.Println("Locale", binRes.Locale)
		log.Println("System Time", binRes.SystemTime)
		log.Println("Card Type", binRes.CardType)
		log.Println("Card Association", binRes.CardAssociation)
		log.Println("Card Family", binRes.CardFamily)
		log.Println("Bank Name", binRes.BankName)
		log.Println("Bank Code", binRes.BankCode)
		log.Println("Commercial", binRes.Commercial)
	}
}