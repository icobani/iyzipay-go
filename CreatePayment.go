package Iyzipay

import (
	"encoding/json"
	"fmt"
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

func (i Iyzipay) ThreeDSPayment(obj *Request.ThreeDHandshakeReq) (Response.HandshakeResponse3D, error) {
	result := Response.HandshakeResponse3D{}
	if resp, err := i.Client.R().
		SetHeaders(i.GetHeaders(*obj)).
		SetBody(obj).
		Post(i.GetURI("/payment/3dsecure/initialize")); err == nil {
		fmt.Println(string(resp.Body()))
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	} else {
		return result, err
	}
}
