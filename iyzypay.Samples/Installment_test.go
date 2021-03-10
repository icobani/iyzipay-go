package iyzypay_Samples

import (
	"github.com/icobani/iyzipay-go"
	"github.com/icobani/iyzipay-go/Request"
	"log"
	"testing"
)

/*

	Taksit sorgulama(installment), ilgili kart bir kredi kartı ise ve bir taksit programına dahil ise,
kaç taksit yapılabileceği ve taksit komisyon oranı konusunda bilgi vermektedir.

*/
func TestInstallmentCheck(t *testing.T) {
	var Pay, _ = Iyzipay.Iyzipay{}.New(
		"TBrS4jBYv3HY63qaAOQ1eC0oSmidbhuu",
		"vUoZ6kCGhEnEDwu2qVVZOq5DPKfOFL2r",
	)

	instReq := Request.InstallmentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		BinNumber:      "540709",
		Price:          "5721.8",
	}

	insRes, _ := Pay.InstallmentCheck(&instReq)

	log.Printf("%#v\n", insRes)
}
