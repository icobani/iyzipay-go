package Response
type InstallmentResponse struct {
	Status             string              `json:"status"`
	Locale             string              `json:"locale"`
	SystemTime         int                 `json:"systemTime"`
	ConversationId     string              `json:"conversationId"`
	InstallmentDetails []InstallmentDetail `json:"installmentDetails"`
	ErrorCode          string              `json:"errorCode"`
	ErrorMessage       string              `json:"errorMessage"`
}

type InstallmentDetail struct {
	BinNumber         string              `json:"binNumber"`
	Price             float32            `json:"price"`
	CardType          string             `json:"cardType"`
	CardAssociation   string             `json:"cardAssociation"`
	CardFamilyName    string             `json:"cardFamilyName"`
	Force3ds          int                `json:"force3ds"`
	BankCode          int                `json:"bankCode"`
	BankName          string             `json:"bankName"`
	ForceCvc          int                `json:"forceCvc"`
	Commercial        int                `json:"commercial"`
	InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
}

type InstallmentPrice struct {
	InstallmentPrice  float32 `json:"installmentPrice"`
	TotalPrice        float32 `json:"totalPrice"`
	InstallmentNumber float32 `json:"installmentNumber"`
}