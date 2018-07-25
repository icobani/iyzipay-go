package Response

type BinCheckResponse struct {
	Status          string `json:"status"`
	Locale          string `json:"locale"`
	SystemTime      int    `json:"systemTime"`
	ConversationId  string `json:"conversationId"`
	BinNumber       string `json:"binNumber"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	BankName        string `json:"bankName"`
	BankCode        int    `json:"bankCode"`
	Commercial      int    `json:"commercial"`
	ErrorCode       string `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
}
