package Response

import "time"

/*
Servisin verdiği yanıta göre şu kurallar uygulanmalıdır:

    card_type parametresi DEBIT_CARD olarak döndüğünde bu kart ile yapılmak istenen ödeme işlemi
		3D ile ödeme kullanılarak yapılmalıdır.

    card_family alanındaki değerlere göre taksit seçenekleri ön yüzde gösterilebilir. Taksit yapabileceğiniz
		aileler: Bonus, World, Maximum, Axess, Cardfinans, Paraf ve Advantage'tır.

    force3ds(installmentDetails) değeri 1 olarak döner ise, işlem 3D olarak yapılmalıdır.
		0 ise işlem 3D olmadan yapılabilir ek olarak 3D de tercih edilebilir. Üye işyerine
		3D zorunlu olarak işaretlenmiş ise bu değer sürekli 1 olarak döner.

    Taksit sorgulama, iyzico kontrol panelinde, “Taksit ve Komisyon Yönetimi” bölümünde belirlenen oranları
		döndürmektedir. Bu bölümde yapılan değişiklikler otomatik olarak installment servisine yansıyacaktır.

    Eğer istek yaparken “bin” numarası gönderilmez ise, tüm taksit seçenekleri dönmektedir.
		“Bin” numarası gönderilir ise, sadece o karta özel taksit seçenekleri dönmektedir.

*/
type BinCheckResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Status string `json:"status"`

	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Locale string `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir GetSystemTime metodu ile time değeri alınır
	SystemTime      int64  `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir
	ConversationId  string `json:"conversationId"`

	// Kartın ilk 6 hanesidir.
	BinNumber       string `json:"binNumber"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu tipi.
	// Geçerli değerler: CREDIT_CARD, DEBIT_CARD, PREPAID_CARD
	CardType        string `json:"cardType"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu kuruluş. Geçerli değerler:
	// VISA, MASTER_CARD, AMERICAN_EXPRESS, TROY
	CardAssociation string `json:"cardAssociation"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu aile.
	// Geçerli değerler:
	// Bonus, Axess, World, Maximum, Paraf, CardFinans, Advantage
	CardFamily      string `json:"cardFamily"`

	// Banka Adı
	BankName        string `json:"bankName"`

	// Banka Kodu
	BankCode        int    `json:"bankCode"`


	Commercial      int    `json:"commercial"`

	// Varsa hata kodu
	ErrorCode       string `json:"errorCode"`

	// Varsa hata mesajı
	ErrorMessage    string `json:"errorMessage"`
}
// Gelen Unix SystemTime değerini time olarak döndürür
func (b BinCheckResponse) GetSystemTime() time.Time {
	return time.Unix(0, b.SystemTime*int64(time.Millisecond))
}
