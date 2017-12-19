package HashGenerator

import (
	"github.com/icobani/iyzipay-go/Request"
	"io"
	"crypto/sha1"
	"encoding/base64"
	"log"
)

func GenerateHash(apiKey string, secretKey string, randomString string, request Request.BaseRequest) string {
	h := sha1.New()
	// api key051220170822507398secret key[locale=tr,conversationId=123456789,binNumber=554960,price=100.0]
	hashStr := apiKey + randomString + secretKey + request.ToPKIRequestString()
	log.Println(hashStr)
	io.WriteString(h, hashStr)
	data :=  []byte(h.Sum(nil))
	str := base64.StdEncoding.EncodeToString(data)

	return str
}
