package Response

type CreatePaymentResponse struct {
	Status                       string             `json:"status"`
	Locale                       string             `json:"locale"`
	SystemTime                   int                `json:"systemTime"`
	ConversationId               string             `json:"conversationId"`
	Price                        float32            `json:"price"`
	PaidPrice                    float32            `json:"paidPrice"`
	Installment                  int                `json:"installment"`
	PaymentId                    string             `json:"paymentId"`
	FraudStatus                  int                `json:"fraudStatus"`
	MerchantCommissionRate       float32            `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float32            `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float32            `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float32            `json:"iyziCommissionFee"`
	CardType                     string             `json:"cardType"`
	CardAssociation              string             `json:"cardAssociation"`
	CardFamily                   string             `json:"cardFamily"`
	BinNumber                    string             `json:"binNumber"`
	LastFourDigits               string             `json:"lastFourDigits"`
	BasketId                     string             `json:"basketId"`
	Currency                     string             `json:"currency"`
	Token                        string             `json:"token"`
	PaymentStatus                string             `json:"paymentStatus"`
	ItemTransactions             []ItemTransactions `json:"itemTransactions"`
	CardToken                    string             `json:"cardToken"`
	CardUserKey                  string             `json:"cardUserKey"`
	AuthCode                     string             `json:"authCode"`
	Phase                        string             `json:"phase"`
	ErrorCode                    string             `json:"errorCode"`
	ErrorMessage                 string             `json:"errorMessage"`
	ErrorGroup                   string             `json:"errorGroup"`
}

type ItemTransactions struct {
	ItemId                        string          `json:"itemId"`
	PaymentTransactionId          string          `json:"paymentTransactionId"`
	TransactionStatus             int             `json:"transactionStatus"`
	Price                         float32         `json:"price"`
	PaidPrice                     float32         `json:"paidPrice"`
	MerchantCommissionRate        float32         `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount  float32         `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount      float32         `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float32         `json:"iyziCommissionFee"`
	BlockageRate                  float32         `json:"blockageRate"`
	BlockageRateAmountMerchant    float32         `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float32         `json:"blockageRateAmountSubMerchant"`
	BlockageResolvedDate          string          `json:"blockageResolvedDate"`
	SubMerchantPrice              float32         `json:"subMerchantPrice"`
	SubMerchantPayoutRate         float32         `json:"subMerchantPayoutRate"`
	SubMerchantPayoutAmount       float32         `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float32         `json:"merchantPayoutAmount"`
	ConvertedPayout               ConvertedPayout `json:"convertedPayout"`
}

type ConvertedPayout struct {
	PaidPrice                     float32 `json:"paidPrice"`
	IyziCommissionRateAmount      float32 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee             float32 `json:"iyziCommissionFee"`
	BlockageRateAmountMerchant    float32 `json:"blockageRateAmountMerchant"`
	BlockageRateAmountSubMerchant float32 `json:"blockageRateAmountSubMerchant"`
	SubMerchantPayoutAmount       float32 `json:"subMerchantPayoutAmount"`
	MerchantPayoutAmount          float32 `json:"merchantPayoutAmount"`
	IyziConversionRate            float32 `json:"iyziConversionRate"`
	IyziConversionRateAmount      float32 `json:"iyziConversionRateAmount"`
	Currency                      string  `json:"currency"`
}
