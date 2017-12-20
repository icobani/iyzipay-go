package iyzipay_go

import (
	"net/http"
	"time"
	"github.com/icobani/iyzipay-go/Request"
	"github.com/icobani/iyzipay-go/HashGenerator"
	"strings"
)

const AUTHORIZATION = "Authorization"
const RANDOM_HEADER_NAME = "x-iyzi-rnd"
const CLIENT_VERSION = "x-iyzi-client-version"
const IYZIWS_HEADER_NAME = "IYZWS "
const COLON = ":"

type Options struct {
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
	BaseURL   string `json:"baseUrl"`
}

func SetHttpHeader(request Request.BaseRequest, options Options, httpReq *http.Request) {
	t := time.Now()
	randomString := strings.Replace(t.Format("02012006150405.0000"), ".", "", -1)
	// karşılaştırma amaçlı
	//randomString = "061220171928250995"

	httpReq.Header.Add("Accept", "application/json")
	httpReq.Header.Add(RANDOM_HEADER_NAME, randomString)

	httpReq.Header.Add(CLIENT_VERSION, "iyzipay-dotnet-2.1.12")
	httpReq.Header.Add(AUTHORIZATION, PrepareAuthorizationString(request, options, randomString))

	httpReq.Header["Accept"] = []string{"application/json"}
	httpReq.Header[RANDOM_HEADER_NAME] = []string{randomString}
	httpReq.Header[CLIENT_VERSION] = []string{"iyzipay-dotnet-2.1.12"}
	httpReq.Header[AUTHORIZATION] = []string{PrepareAuthorizationString(request, options, randomString)}


	requestHeaders := make(map[string]string)
	requestHeaders["Accept"] = "application/json"
	requestHeaders[RANDOM_HEADER_NAME] = randomString
	requestHeaders[CLIENT_VERSION] = "iyzipay-dotnet-2.1.12"
	requestHeaders[AUTHORIZATION] = PrepareAuthorizationString(request, options, randomString)
	// httpReq.Header = requestHeaders




}

func PrepareAuthorizationString(request Request.BaseRequest, options Options, randomString string) string {
	hash := HashGenerator.GenerateHash(options.ApiKey, options.SecretKey, randomString, request)
	// q0/NoRvDVCsUZbDkuZub1nw4ZQE=
	// q0/NoRvDVCsUZbDkuZub1nw4ZQE=
	return IYZIWS_HEADER_NAME + options.ApiKey + COLON + hash
}
