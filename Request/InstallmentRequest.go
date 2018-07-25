package Request

type InstallmentRequest struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`
	Price          string `json:"price"`
}
var InstallmentCheckURI string = "/payment/iyzipos/installment"