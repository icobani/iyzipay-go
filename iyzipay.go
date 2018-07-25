package Iyzipay

import (
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"time"
	"reflect"
	"fmt"
	"net/http"
)

type Iyzipay struct {
	APIKey    string
	APISecret string
	URI       string
}

type ErrorStruct struct {
	Error string
}

func (i Iyzipay) New(APIKey, APISecret, URI string) (*Iyzipay, *ErrorStruct) {
	if APIKey == "" {
		return nil, &ErrorStruct{"APIKey not found."}
	}
	if APISecret == "" {
		return nil, &ErrorStruct{"APISecret not found."}
	}
	if URI == "" {
		return nil, &ErrorStruct{"URI not found."}
	}
	return &Iyzipay{APIKey, APISecret, URI}, nil
}

// RandomString : Generating romdom string
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randomString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// PKI string generator...
func pKIStringify(v interface{}) (res string) {
	rv := reflect.ValueOf(v)
	num := rv.NumField()
	for i := 0; i < num; i++ {
		fv := rv.Field(i)
		st := rv.Type().Field(i)
		jsonName, err := st.Tag.Lookup("json")
		if err != true {
			res += st.Name + "="
		} else {
			res += jsonName + "="
		}
		switch fv.Kind() {
		case reflect.String:
			res += fv.String()
		case reflect.Int:
			res += fmt.Sprint(fv.Int())
		case reflect.Struct:
			res += pKIStringify(fv.Interface())
		case reflect.Slice:
			res += "["
			for k := 0; k < fv.Len(); k++ {
				e := fv.Index(k)
				res += pKIStringify(e.Interface())
				if k != fv.Len()-1 {
					res += ", "
				}
			}
			res += "]"
		}
		if i != num-1 {
			res += ","
		}
	}
	return "[" + res + "]"
}

// generateAuthorizationHeader : Creting authorization header for request...
func generateAuthorizationHeader(iyziWsHeaderName, apiKey, separator, secretKey, body, randomString string) string {
	return iyziWsHeaderName + " " + apiKey + separator + generateHash(apiKey, randomString, secretKey, body)
}

// generateHash : Generating hash
func generateHash(apiKey, randomString, secretKey, body string) string {
	hasher := sha1.New()
	hasher.Write([]byte(apiKey + randomString + secretKey + body))
	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func (i Iyzipay) InitHttpRequest(req *http.Request, obj interface{}) {

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	RandomString := randomString(8)
	pki := pKIStringify(obj)

	authorization := generateAuthorizationHeader("IYZWS", i.APIKey, ":", i.APISecret, pki, RandomString)
	req.Header.Add("authorization", authorization)
	req.Header.Add("x-iyzi-rnd", RandomString)
	req.Header.Add("cache-control", "no-cache")
}
