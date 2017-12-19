package iyzypay_Samples

import (
	"testing"
	"github.com/icobani/iyzipay-go"
	"github.com/icobani/iyzipay-go/HealthCheck"
	"github.com/icobani/iyzipay-go/Request"
	"github.com/icobani/iyzipay-go/Locale"
	"github.com/icobani/iyzipay-go/Installment"
)

var apiOptions = iyzipay_go.Options{
	ApiKey:    "sandbox-lwggMKTv5mSwC1Z4e6Zay4n6HKUOhqOX",
	SecretKey: "sandbox-Ok8ZKrca4DWyEjGX0GHq8zwOCpne0rEQ",
	BaseURL:   "https://sandbox-api.iyzipay.com",
}

/*
Function for health Check
 */
func TestHealthCheck(t *testing.T) {
	res := HealthCheck.Retrive(apiOptions)
	if res.Status != "success" {
		t.Error("Api Error", res)
	} else {

		t.Log("Health Check is ", res.Status, " System Date : ", res.SystemTimeC)
	}
}

/*
	Taksit sorgulama(installment), ilgili kart bir kredi kartı ise ve bir taksit programına dahil ise,
	kaç taksit yapılabileceği ve taksit komisyon oranı konusunda bilgi vermektedir.
 */
func TestRetrieveInstallments(t *testing.T) {

	rec := Request.RetrieveInstallmentInfoRequest{
		BinNumber:      "589004",
		Price:          "200.0",
		ConversationId: "123123123",
		Locale:         Locale.TR,
	}

	t.Log(Installment.Info{}.Retrieve(rec, apiOptions))

}
