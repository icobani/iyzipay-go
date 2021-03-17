package Response

import (
	"strings"
)

type CreatePaymentResponse struct {
	Status             string `json:"status"`
	ErrorCode          string `json:"errorCode"`
	ErrorMessage       string `json:"errorMessage"`
	Locale             string `json:"locale"`
	SystemTime         int    `json:"systemTime"`
	ConversationId     string `json:"conversationId"`
	ThreeDSHtmlContent string `json:"threeDSHtmlContent"`
}

// threedsHtmlContent parmetresini ekrana bastırdığınızda bankanın 3ds ekranı görüntülenecektir.
// Kart sahibi cep telefonuna gelen şifreyi girdikten sonra otomatik olarak callbackUrl parametresinde belirttiğiniz adrese yönlenecek
// ve iyzico bu adrese aşağıdaki değerleri post edecektir.
type ThreeDPaymentResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner
	Status string `json:"status"`
	// İşlem success yani başarılı ise, iyzico tarafından o ödemeye verilen değerdir. 3D bitirme sorgusunda kullanılacaktır
	PaymentID string `json:"paymentId"`
	// İşlem success yani başarılı ise, iyzico tarafından o ödemeye verilen değerdir. Bütün durumlarda dolu gelmeyebilir. Eğer dolu gelirse bu parametrenin de 3D bitirme sorgusunda gönderilmesi gerekir.
	ConversationData string `json:"conversationData"`
	// Başlatma sorgusunda gönderilen conversationId değeridir.
	ConversationID string `json:"conversationId"`
	// Bilgilendirme amaçlı dönen mdStatus değeridir. Başarılı durumlar için 1 başarısız durumlar için ise 0,2,3,4,5,6,7,8 olarak dönebilir
	//
	// mdStatus = 0	3-D Secure imzası geçersiz veya doğrulama
	//
	// mdStatus = 2	Kart sahibi veya bankası sisteme kayıtlı değil
	//
	// mdStatus = 3	Kartın bankası sisteme kayıtlı değil
	//
	// mdStatus = 4	Doğrulama denemesi, kart sahibi sisteme daha sonra kayıt olmayı seçmiş
	//
	// mdStatus = 5	Doğrulama yapılamıyor
	//
	// mdStatus = 6	3-D Secure hatası
	//
	// mdStatus = 7	Sistem hatası
	//
	// mdStatus = 8	Bilinmeyen kart no
	MDStatus string `json:"mdStatus"`
}

// "status=success&paymentId=1490176961&conversationData=&conversationId=123456789&mdStatus=1"
func (t *ThreeDPaymentResponse) Bind(query string) {
	params := strings.Split(query, "&")
	m := make(map[string]string)
	for _, part := range params {
		vals := strings.Split(part, "=")
		m[vals[0]] = vals[1]
	}
	t.Status = m["status"]
	t.PaymentID = m["paymentId"]
	t.ConversationData = m["conversationData"]
	t.ConversationID = m["conversationId"]
	t.MDStatus = m["mdStatus"]
}

// Handshake Response
// Dönüş parametreleri ise aşağıdaki gibidir:
//
// status parametresi işlemin durumu hakkında bilgi verir. success ise işlem başarılı bir şekilde yapılmış ve para çekilmiş demektir. failure ise işlem başarısız olmuş ve sebebi ile alakalı olarak hata bilgilendirmesi yapılmış demektir.
// paymentid ve paymenttransactionid değerleri mutlaka saklanmalıdır.

type HandshakeResponse3D struct {
	AuthCode         string `json:"authCode"`
	BasketID         string `json:"basketId"`
	BinNumber        string `json:"binNumber"`
	CardAssociation  string `json:"cardAssociation"`
	CardFamily       string `json:"cardFamily"`
	CardType         string `json:"cardType"`
	ConversationID   string `json:"conversationId"`
	Currency         string `json:"currency"`
	FraudStatus      int64  `json:"fraudStatus"`
	HostReference    string `json:"hostReference"`
	Installment      int64  `json:"installment"`
	ItemTransactions []struct {
		BlockageRate                  float64 `json:"blockageRate"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
		BlockageResolvedDate          string  `json:"blockageResolvedDate"`
		ConvertedPayout               struct {
			BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
			BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
			Currency                      string  `json:"currency"`
			IyziCommissionFee             float64 `json:"iyziCommissionFee"`
			IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
			IyziConversionRate            float64 `json:"iyziConversionRate"`
			IyziConversionRateAmount      float64 `json:"iyziConversionRateAmount"`
			MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
			PaidPrice                     float64 `json:"paidPrice"`
			SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
		} `json:"convertedPayout"`
		ItemID                       string  `json:"itemId"`
		IyziCommissionFee            float64 `json:"iyziCommissionFee"`
		IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
		MerchantCommissionRate       float64 `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
		MerchantPayoutAmount         float64 `json:"merchantPayoutAmount"`
		PaidPrice                    float64 `json:"paidPrice"`
		PaymentTransactionID         string  `json:"paymentTransactionId"`
		Price                        float64 `json:"price"`
		SubMerchantPayoutAmount      float64 `json:"subMerchantPayoutAmount"`
		SubMerchantPayoutRate        float64 `json:"subMerchantPayoutRate"`
		SubMerchantPrice             float64 `json:"subMerchantPrice"`
		TransactionStatus            int64   `json:"transactionStatus"`
	} `json:"itemTransactions"`
	IyziCommissionFee            float64 `json:"iyziCommissionFee"`
	IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
	LastFourDigits               string  `json:"lastFourDigits"`
	Locale                       string  `json:"locale"`
	MdStatus                     int64   `json:"mdStatus"`
	MerchantCommissionRate       float64 `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
	PaidPrice                    float64 `json:"paidPrice"`
	PaymentID                    string  `json:"paymentId"`
	Phase                        string  `json:"phase"`
	Price                        float64 `json:"price"`
	Status                       string  `json:"status"`
	SystemTime                   int64   `json:"systemTime"`
}
