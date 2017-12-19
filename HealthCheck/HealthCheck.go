package HealthCheck

import (
	"github.com/icobani/iyzipay-go"
	"net/http"
	"time"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"math"

)

func Retrive(Options iyzipay_go.Options) iyzipay_go.IyzipayResource {
	var res iyzipay_go.IyzipayResource
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := netClient.Get(fmt.Sprintf("%s/payment/test", Options.BaseURL))

	if err != nil {
		// err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			log.Println("Errpr:", err)
		}
		bodyString := string(bodyBytes)

		json.Unmarshal([]byte(bodyString), &res)
		st := float64(res.SystemTime/1000)

		res.SystemTimeC = time.Unix(
			int64(math.Ceil(st)),
			0)

	}
	return res
}
