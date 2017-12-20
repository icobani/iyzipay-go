package Request

import "fmt"

type RetrieveInstallmentInfoRequest struct {
	BinNumber      string `json:"binNumber"`
	Price          string `json:"price"`
	Locale         string `json:"locale"`
	ConversationId string `json:"conversationId"`
}

func (this RetrieveInstallmentInfoRequest) ToPKIRequestString() string {
	return fmt.Sprintf("[locale=%s,conversationId=%s,binNumber=%s,price=%s]", this.Locale, this.ConversationId, this.BinNumber, this.Price)
}
