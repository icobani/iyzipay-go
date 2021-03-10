package Request

type InstallmentRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır.
	// Varsayılan değeri tr’dir.
	Locale string `json:"locale"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer,
	// request/response eşleşmesi yapmak için kullanılabilir.
	ConversationId string `json:"conversationId"`

	// Kartın ilk 6 hanesidir.
	BinNumber string `json:"binNumber"`

	// Taksitlendirilmek istenen tutar bilgisi.
	Price string `json:"price"`
}
