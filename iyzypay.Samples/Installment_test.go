package iyzypay_Samples

import (
	"testing"
	"github.com/icobani/iyzipay-go"
	"github.com/icobani/iyzipay-go/Request"
	"log"
)

/*

	Taksit sorgulama(installment), ilgili kart bir kredi kartı ise ve bir taksit programına dahil ise,
kaç taksit yapılabileceği ve taksit komisyon oranı konusunda bilgi vermektedir.

 */
func TestInstallmentCheck(t *testing.T) {
	var Pay, _ = Iyzipay.Iyzipay{}.New(
		"sandbox-lwggMKTv5mSwC1Z4e6Zay4n6HKUOhqOX",
		"sandbox-Ok8ZKrca4DWyEjGX0GHq8zwOCpne0rEQ",
		"https://sandbox-api.iyzipay.com",
	)

	instReq := Request.InstallmentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		BinNumber:      "540709",
		Price:          "5721.8",
	}

	insRes, _ := Pay.InstallmentCheck(&instReq)

	log.Println("Status", insRes.Status)
	log.Println("Error Code", insRes.ErrorCode)
	log.Println("Error Message", insRes.ErrorMessage)
	for _, item := range insRes.InstallmentDetails {
		log.Println("Bank Name", item.BankName)
		log.Println("Force 3DS", item.Force3ds)
		log.Println("Force Cvc", item.ForceCvc)
		log.Println("Card Type", item.CardType)
		log.Println("Card Family", item.CardFamilyName)
		log.Println("Card Association", item.CardAssociation)
		for _, installments := range item.InstallmentPrices {
			log.Println("Installemnts Number : ", installments.InstallmentNumber,
				"Installment Price", installments.InstallmentPrice,
				"Total Price", installments.TotalPrice)
		}
	}
}
