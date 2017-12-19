package iyzipay_go

import "time"

type IyzipayResource struct {
	Status string `json:"status"`
	ErrorCode string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	ErrorGroup string `json:"error_group"`
	Locale string `json:"locale"`
	SystemTime int64 `json:"systemTime"`
	SystemTimeC time.Time `json:"systemTimeC"`
	ConversationId string `json:"conversation_id"`
}
