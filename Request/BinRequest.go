package Request
/*
	Taksit sorgulama(installment), ilgili kart bir kredi kartı ise ve bir taksit programına dahil ise,
	kaç taksit yapılabileceği ve taksit komisyon oranı konusunda bilgi vermektedir.
*/
type BinCheckRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır. Varsayılan değeri tr’dir.
	Locale         string `json:"locale"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer,
	// request/response eşleşmesi yapmak için kullanılabilir.
	ConversationId string `json:"conversationId"`

	// Kartın İlk 6 hanesi
	BinNumber      string `json:"binNumber"`
}

// api adresi
var BinCheckUrl string = "/payment/bin/check"