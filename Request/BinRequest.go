package Request

type BinCheckRequest struct {
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
	BinNumber      string `json:"binNumber"`
}


var BinCheckUrl string = "/payment/bin/check"