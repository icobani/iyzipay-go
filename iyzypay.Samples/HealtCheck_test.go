package iyzypay_Samples

import (
	"fmt"
	"github.com/icobani/iyzipay-go"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	var Pay, _ = Iyzipay.Iyzipay{}.New(
		"<api-key>",
		"<api-secret>",
	)
	healtResp, _ := Pay.HealtCheck()

	if healtResp.Status != "success" {
		t.Errorf("Cannot connect the system. Check the internet connections %s", healtResp.Status)
	} else {
		fmt.Printf("%#v\n", healtResp)
	}
}
