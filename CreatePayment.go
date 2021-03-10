package Iyzipay

import (
	"encoding/json"
	"github.com/icobani/iyzipay-go/Request"
	"github.com/icobani/iyzipay-go/Response"
)

func (i Iyzipay) CreatePayment(obj *Request.CreatePaymentRequest) (Response.CreatePaymentResponse, error) {
	result := Response.CreatePaymentResponse{}
	if resp, err := i.Client.R().
		SetHeaders(i.GetHeaders(*obj)).
		SetBody(obj).
		Post(i.GetURI("/payment/3dsecure/initialize")); err == nil {
		//fmt.Println(string(resp.Body()))
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	} else {
		return result, err
	}
}
