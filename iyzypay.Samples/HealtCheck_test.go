package iyzypay_Samples

import (
	"testing"
	"github.com/icobani/iyzipay-go"
	"log"
)

func TestHealthCheck(t *testing.T) {
	var Pay, _ = Iyzipay.Iyzipay{}.New(
		"sandbox-lwggMKTv5mSwC1Z4e6Zay4n6HKUOhqOX",
		"sandbox-Ok8ZKrca4DWyEjGX0GHq8zwOCpne0rEQ",
		"https://sandbox-api.iyzipay.com",
	)
	healtResp, _ := Pay.HealtCheck()

	if healtResp.Status != "success" {
		t.Errorf("Cannot connect the system. Check the internet connections %s",healtResp.Status)
	} else {
		log.Println("Status",healtResp.Status)
	}
}
