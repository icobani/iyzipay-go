package Iyzipay

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/icobani/iyzipay-go/Response"
	"github.com/icobani/iyzipay-go/Request"
)

func (i Iyzipay) HealtCheck() (*Response.HealtCheckResponse, error) {

	req, err := http.NewRequest("GET", i.URI+Request.HealthCheckURI, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := &Response.HealtCheckResponse{}

	err = json.Unmarshal(body,result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
