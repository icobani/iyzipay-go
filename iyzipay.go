package Iyzipay

import (
	"github.com/go-resty/resty/v2"
	"github.com/icobani/iyzipay-go/Utils"
	"strings"
)

type Iyzipay struct {
	APIKey    string
	APISecret string
	URI       string
	Client    *resty.Client
}

type ErrorStruct struct {
	Error string
}

func (i Iyzipay) New(APIKey, APISecret string) (*Iyzipay, *ErrorStruct) {
	if APIKey == "" {
		return nil, &ErrorStruct{"APIKey not found."}
	}
	if APISecret == "" {
		return nil, &ErrorStruct{"APISecret not found."}
	}
	URI := "https://api.iyzipay.com"
	if strings.Contains(APIKey, "sandbox-") {
		URI = "https://sandbox-api.iyzipay.com"
	}
	return &Iyzipay{APIKey, APISecret, URI, resty.New()}, nil
}

func (i Iyzipay) GetURI(uri string) string {
	return i.URI + uri
}

func (i Iyzipay) GetHeaders(obj interface{}) map[string]string {
	header := map[string]string{}
	header["accept"] = "application/json"
	header["content-type"] = "application/json"
	RandomString := Utils.RandomString(8)
	requestString := Utils.RequestString(obj)
	pki := Utils.PKIString(i.APIKey, RandomString, i.APISecret, requestString)
	hashedPki := Utils.HashedPKIString(pki)
	authorization := Utils.GenerateAuthorizationHeader(i.APIKey, hashedPki)
	header["authorization"] = authorization
	header["x-iyzi-rnd"] = RandomString
	header["cache-control"] = "no-cache"

	return header
}
