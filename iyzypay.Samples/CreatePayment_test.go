package iyzypay_Samples

import (
	"encoding/base64"
	"fmt"
	"github.com/icobani/iyzipay-go"
	"github.com/icobani/iyzipay-go/Request"
	"log"
	"testing"
)

func TestCreatePayment(t *testing.T) {

	var Pay, _ = Iyzipay.Iyzipay{}.New(
		"<api-key>",
		"<api-secret>",
	)

	Req := Request.CreatePaymentRequest{
		Locale:         "tr",
		ConversationId: "123456789",
		Price:          "1.0",
		PaidPrice:      "1.2",
		Currency:       "TRY",
		Installment:    1,
		PaymentChannel: "WEB",
		BasketId:       "B67832",
		PaymentGroup:   "PRODUCT",
		CallbackURL:    "https://www.mecrchantsite.com/callback",
		PaymentCard: Request.CreatePayment_PaymentCard{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			ExpireYear:     "2030",
			ExpireMonth:    "12",
			Cvc:            "123",
			// RegisterCard:   1,
		},
		Buyer: Request.CreatePayment_Buyer{
			Id:                  "7334",
			Name:                "John",
			Surname:             "ÇOBANİ",
			IdentityNumber:      "11924186810",
			Email:               "ibrahim@imaconsult.com",
			GsmNumber:           "+905325401194",
			RegistrationDate:    "2021-02-21 15:12:09",
			LastLoginDate:       "2021-03-05 12:43:35",
			RegistrationAddress: "Kiremit cd. No:42 Balat Fatih",
			ZipCode:             "34732",
			City:                "Istanbul",
			Country:             "Turkey",
			Ip:                  "85.34.78.112",
		},
		ShippingAddress: Request.CreatePayment_ShippingAddress{
			Address:     "Kiremit cd. No:42 Balat Fatih",
			ZipCode:     "34732",
			City:        "Istanbul",
			Country:     "Turkey",
			ContactName: "Jane Doe",
		},
		BillingAddress: Request.CreatePayment_BillingAddress{
			Address:     "Kiremit cd. No:42 Balat Fatih",
			ZipCode:     "34732",
			City:        "Istanbul",
			Country:     "Turkey",
			ContactName: "Jane Doe",
		},
	}

	Req.BasketItems = []Request.CreatePayment_BasketItems{
		{
			Id:        "BI101",
			Price:     "0.3",
			Name:      "Binocular",
			Category1: "Collectibles",
			Category2: "Accessories",
			//SubMerchantKey:   "o+G86/7xLFb9rc6/gVvcM/tGySk=",
			//SubMerchantPrice: "0.18",
			ItemType: "PHYSICAL",
		},
		{
			Id:        "BI102",
			Price:     "0.5",
			Name:      "Game code",
			Category1: "Game",
			Category2: "Online Game Items",
			//SubMerchantKey:   "o+G86/7xLFb9rc6/gVvcM/tGySk=",
			//SubMerchantPrice: "0.18",
			ItemType: "VIRTUAL",
		},
		{
			Id:        "BI103",
			Price:     "0.2",
			Name:      "Usb",
			Category1: "Electronics",
			Category2: "Usb / Cable",
			//SubMerchantKey:   "o+G86/7xLFb9rc6/gVvcM/tGySk=",
			//SubMerchantPrice: "0.18",
			ItemType: "PHYSICAL",
		},
	}

	resp, err := Pay.CreatePayment(&Req)
	if err != nil {
		fmt.Println(err)
	}
	data, _ := base64.StdEncoding.DecodeString(resp.ThreeDSHtmlContent)
	log.Printf("%+v\\n", string(data))
}
