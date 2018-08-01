package Response
type InstallmentResponse struct {
	// Yapılan isteğin sonucunu bildirir. İşlem başarılı ise success, hatalı ise failure döner.
	Status             string              `json:"status"`

	// İstekte belirtilen locale değeri geri dönülür, varsayılan değeri trdir.
	Locale             string              `json:"locale"`

	// Dönen sonucun o anki unix timestamp değeridir.
	SystemTime         int                 `json:"systemTime"`

	// İstek esnasında gönderilmişse, sonuçta aynen geri iletilir.
	ConversationId     string              `json:"conversationId"`

	//
	InstallmentDetails []InstallmentDetail `json:"installmentDetails"`

	// İşlem hatalıysa, bu hataya dair belirtilen koddur.
	ErrorCode          string              `json:"errorCode"`

	// İşlem hatalıysa, bu hataya dair belirtilen mesajdır. Locale parametresine göre dil desteği sunar.
	ErrorMessage       string              `json:"errorMessage"`
}

type InstallmentDetail struct {
	// Kartın ilk 6 hanesidir.
	BinNumber         string              `json:"binNumber"`

	// Toplam tutar.
	Price             float32            `json:"price"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu tipi.
	// Geçerli değerler:
	// * CREDIT_CARD,
	// * DEBIT_CARD,
	// * PREPAID_CARD
	CardType          string             `json:"cardType"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu kuruluş.
	// Geçerli değerler:
	// * TROY,
	// * VISA,
	// * MASTER_CARD,
	// * AMERICAN_EXPRESS
	CardAssociation   string             `json:"cardAssociation"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu aile.
	// Geçerli değerler:
	// * Bonus,
	// * Axess,
	// * World,
	// * Maximum,
	// * Paraf,
	// * CardFinans,
	// * Advantage
	CardFamilyName    string             `json:"cardFamilyName"`

	// İşlemin 3ds yapılmasına gerek olup olmadığını gösterir.
	// 1 veya 0 değerlerini alır.
	// 1 ise işlem 3ds ile yapılmalıdır.
	Force3ds          int                `json:"force3ds"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu banka kodu.
	BankCode          int                `json:"bankCode"`

	// Eğer ödeme yapılan kart yerel bir kart ise, kartın ait olduğu banka adı.
	BankName          string             `json:"bankName"`

	//
	ForceCvc          int                `json:"forceCvc"`


	Commercial        int                `json:"commercial"`

	InstallmentPrices []InstallmentPrice `json:"installmentPrices"`
}

type InstallmentPrice struct {
	//  	Taksit başına düşen tutar.
	InstallmentPrice  float32 `json:"installmentPrice"`

	// Toplam taksitli tutar.
	TotalPrice        float32 `json:"totalPrice"`

	// Taksit sayısı.
	InstallmentNumber float32 `json:"installmentNumber"`
}