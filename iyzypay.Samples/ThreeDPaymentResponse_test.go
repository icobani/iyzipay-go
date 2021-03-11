/*
   © 2021 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 11.03.2021  16:42
   Notes   :
*/
package iyzypay_Samples

import (
	"github.com/icobani/iyzipay-go/Response"
	"testing"
)

func TestThreeDPaymentResp(t *testing.T) {
	body := "status=success&paymentId=1490176961&conversationData=&conversationId=123456789&mdStatus=1"
	var tD = Response.ThreeDPaymentResponse{}
	tD.Bind(body)
}
