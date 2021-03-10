/*
   © 2021 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 25.01.2021  12:04
   Notes   :
*/
package Utils

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type ApiMethods struct {
	RETRIEVE string
	CREATE   string
	DELETE   string
	UPDATE   string
}

//var APIMethod = &ApiMethods{"retrieve", "create", "delete", "update"}

// GenerateAuthorizationHeader : Creting authorization header for request...
func GenerateAuthorizationHeader(apiKey, hashedPki string) string {
	return fmt.Sprintf("IYZWS %s:%s", apiKey, hashedPki)
}

// Hashed PKI String
func HashedPKIString(pki string) string {
	hasher := sha1.New()
	hasher.Write([]byte(pki))
	sha := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

// RandomString : Generating romdom string
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
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

// PKI String (apiKey + x-iyzi-rnd + secretKey + requestString)
func PKIString(apiKey string, randomString string, secretKey string, requestString string) string {
	return apiKey + randomString + secretKey + requestString
}

// Request String...
func RequestString(v interface{}) (res string) {
	rv := reflect.ValueOf(v)
	num := rv.NumField()
	for i := 0; i < num; i++ {
		fv := rv.Field(i)
		st := rv.Type().Field(i)
		jsonName, err := st.Tag.Lookup("json")
		if strings.Contains(jsonName, ",omitempty") {
			continue
		}
		if err != true {
			res += st.Name + "="
		} else {
			res += fmt.Sprintf("%s=", jsonName)
		}
		switch fv.Kind() {
		case reflect.String:
			res += fmt.Sprintf("%s", fv.String())
		case reflect.Int:
			res += fmt.Sprint(fv.Int())
		case reflect.Struct:
			res += RequestString(fv.Interface())
		case reflect.Slice:
			res += "["
			for k := 0; k < fv.Len(); k++ {
				e := fv.Index(k)
				res += RequestString(e.Interface())
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
	res = "[" + res + "]"
	return res
}
