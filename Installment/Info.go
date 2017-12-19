package Installment

import (
	"github.com/icobani/iyzipay-go/Request"
	iyzipayGo "github.com/icobani/iyzipay-go"
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net"
)

type Info struct {
	InstallmentDetails string `json:"installment_details"`
}

func (i Info) Retrieve(insReq Request.RetrieveInstallmentInfoRequest, options iyzipayGo.Options) string {

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var client = &http.Client{
		Timeout: time.Second * 10,
		Transport: netTransport,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment/iyzipos/installment", options.BaseURL), nil)

	if err != nil {
		// err
	}



	iyzipayGo.SetHttpHeader(insReq, options, req)



	// log.Println("Header  : ",req.Header)
	log.Println("x-iyzi-client-version : ",req.Header["x-iyzi-client-version"])
	log.Println("Authorization : ",req.Header["Authorization"])
	log.Println("x-iyzi-rndn : ",req.Header["x-iyzi-rnd"])
	log.Println("Accept : ",req.Header["Accept"])

	response, err := client.Do(req)

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			log.Println("Errpr:", err)
		}
		bodyString := string(bodyBytes)
		return bodyString
		//json.Unmarshal([]byte(bodyString), &res)
	}
	return ""
	//return res
}
