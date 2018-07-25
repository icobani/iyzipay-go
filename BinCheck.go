package Iyzipay

import (
	"github.com/icobani/iyzipay-go/Request"
	"encoding/json"
	"strings"
	"net/http"
	"io/ioutil"
	"github.com/icobani/iyzipay-go/Response"
)

func (i Iyzipay) BinCheck(obj *Request.BinCheckRequest) (*Response.BinCheckResponse, error) {
	requestBody, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	payload := strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", i.URI+Request.BinCheckUrl, payload)
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

	result := &Response.BinCheckResponse{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
