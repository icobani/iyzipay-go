package Iyzipay

import (
	"encoding/json"
	"github.com/icobani/iyzipay-go/Request"
	"github.com/icobani/iyzipay-go/Response"
)

func (i Iyzipay) BinCheck(obj *Request.BinCheckRequest) (Response.BinCheckResponse, error) {
	// minnak kurbaa
	result := Response.BinCheckResponse{}
	if resp, err := i.Client.R().
		SetHeaders(i.GetHeaders(*obj)).
		SetBody(obj).
		Post(i.GetURI("/payment/bin/check")); err == nil {
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	} else {
		return result, err
	}
}
