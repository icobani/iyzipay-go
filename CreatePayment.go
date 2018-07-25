package Iyzipay

import (
	"github.com/icobani/iyzipay-go/Request"
	"github.com/icobani/iyzipay-go/Response"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)



func (i Iyzipay) CreatePayment(obj *Request.CreatePaymentRequest) (*Response.CreatePaymentResponse, error) {
	requestBody, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", i.URI+Request.CreatePaymentEndpoint, payload)
	if err != nil {
		return nil, err
	}
	i.InitHttpRequest(req, *obj)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	result := &Response.CreatePaymentResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
