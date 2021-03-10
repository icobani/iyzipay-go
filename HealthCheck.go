package Iyzipay

import (
	"encoding/json"
	"github.com/icobani/iyzipay-go/Response"
)

func (i Iyzipay) HealtCheck() (*Response.HealtCheckResponse, error) {
	result := &Response.HealtCheckResponse{}
	if resp, err := i.Client.R().
		Get(i.GetURI("/payment/test")); err == nil {
		err = json.Unmarshal(resp.Body(), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	} else {
		return result, err
	}
}
