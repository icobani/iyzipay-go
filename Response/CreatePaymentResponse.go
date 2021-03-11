package Response

type CreatePaymentResponse struct {
	Status             string `json:"status"`
	ErrorCode          string `json:"errorCode"`
	ErrorMessage       string `json:"errorMessage"`
	Locale             string `json:"locale"`
	SystemTime         int    `json:"systemTime"`
	ConversationId     string `json:"conversationId"`
	ThreeDSHtmlContent string `json:"threeDSHtmlContent"`
}

// Handshake Response
// Dönüş parametreleri ise aşağıdaki gibidir:
//
// status parametresi işlemin durumu hakkında bilgi verir. success ise işlem başarılı bir şekilde yapılmış ve para çekilmiş demektir. failure ise işlem başarısız olmuş ve sebebi ile alakalı olarak hata bilgilendirmesi yapılmış demektir.
// paymentid ve paymenttransactionid değerleri mutlaka saklanmalıdır.
type HandshakeResponse3D struct {
	Status                       string            `json:"status"`
	ErrorCode                    string            `json:"errorCode"`
	ErrorMessage                 string            `json:"errorMessage"`
	Locale                       string            `json:"locale"`
	SystemTime                   int               `json:"systemTime"`
	ConversationId               string            `json:"conversationId"`
	Price                        float32           `json:"price"`
	PaidPrice                    float32           `json:"paidPrice"`
	Currency                     string            `json:"currency"`
	Installment                  int               `json:"installment"`
	BasketID                     int               `json:"basketId"`
	BinNumber                    string            `json:"binNumber"`
	CardAssociation              string            `json:"cardAssociation"`
	CardFamily                   string            `json:"cardFamily"`
	CardType                     string            `json:"cardType"`
	FraudStatus                  int               `json:"fraudStatus"`
	PaymentID                    string            `json:"paymentId"`
	IyziCommissionRateAmount     float32           `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float32           `json:"iyziCommissionFee"`
	MerchantCommissionRate       float32           `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float32           `json:"merchantCommissionRateAmount"`
	ItemTransactions             []ItemTransaction `json:"ItemTransactions"`
}

type ItemTransaction struct {
	PaymentTransactionId         string          `json:"paymentTransactionId"`
	ItemId                       string          `json:"itemId"`
	Price                        float64         `json:"price"`
	PaidPrice                    float64         `json:"paidPrice"`
	TransactionStatus            int             `json:"transactionStatus"`
	BlockageRate                 float32         `json:"blockageRate"`
	BlockageRateAmountMerchant   float32         `json:"blockageRateAmountMerchant"`
	BlockageResolvedDate         string          `json:"blockageResolvedDate"`
	IyziCommissionFee            float64         `json:"iyziCommissionFee"`
	IyziCommissionRateAmount     float64         `json:"iyziCommissionRateAmount"`
	MerchantCommissionRate       float64         `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64         `json:"merchantCommissionRateAmount"`
	MerchantPayoutAmount         float32         `json:"merchantPayoutAmount"`
	ConvertedPayout              ConvertedPayout `json:"convertedPayout"`
}

type ConvertedPayout struct {
	PaidPrice                  float32 `json:"paidPrice"`
	IyziCommissionFee          float32 `json:"iyziCommissionFee"`
	IyziCommissionRateAmount   float32 `json:"iyziCommissionRateAmount"`
	BlockageRateAmountMerchant float32 `json:"blockageRateAmountMerchant"`
	MerchantPayoutAmount       float32 `json:"merchantPayoutAmount"`
	IyziConversionRate         float32 `json:"iyziConversionRate"`
	IyziConversionRateAmount   float32 `json:"iyziConversionRateAmount"`
	Currency                   string  `json:"currency"`
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
	ConversationID uint64 `json:"conversationId"`
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
	MSStatus string `json:"mdStatus"`
}
