package Request

/*
	Ödeme sorgusu, müşterinin kredi kartından para çekme işlemini yapar. İlgili servis kullanıldığında iyzico
işlemin başarılı olup olmadığı ile alakalı olarak yanıtı anında döner.

    * Üzerinde TROY, MASTERCARD, VISA, AMEX logolu kartlardan işlem yapmaya olanak verir.
    * Üzerinde BONUS, WORLD, MAXIMUM, AXESS, CARDFINANS, PARAF, ADVANTAGE taksit programlarına dahil
		kartlar için 2, 3, 6, 9, 12 taksit seçenekleri bu servis ile kullanılabilir.
    * Başarılı işlemler yeşil, başarısız işlemler ise kırmızı renkte panel üzerinde görüntülenebilir.
    * Üye işyeri tarafında sipariş numarası olarak conversationId ve basketId parametreleri kullanılabilir.

Servise gönderilmesi gereken parametreler şu şekildedir:
*/
type CreatePaymentRequest struct {
	// iyzico istek sonucunda dönen metinlerin dilini ayarlamak için kullanılır.
	// Varsayılan değeri tr’dir. en olarak kullanılabilir.
	Locale string `json:"locale"`

	// İstek esnasında gönderip, sonuçta alabileceğiniz bir değer, request/response eşleşmesi yapmak için
	// kullanılabilir. En yaygın kullanış biçimi üye iş yerinin sipariş numarasıdır.
	ConversationId string `json:"conversationId"`

	// Ödeme sepet tutarı. Kırılım (sepetin içerisindeki ürünler) tutarlar toplamı sepet tutarına eşit olmalı.
	Price string `json:"price"`

	// İndirim vade farkı vs. hesaplanmış POS’tan geçecek nihai tutar. Price değerinden küçük,
	// büyük veya eşit olabilir.
	PaidPrice string `json:"paidPrice"`

	// Taksit bilgisi, tek çekim için 1 gönderilmelidir. Geçerli değerler: 1, 2, 3, 6, 9, 12
	Installment int `json:"installment"`

	// Ödeme kanalı. Geçerli değerler enum içinde sunulmaktadır:
	// * WEB,
	// * MOBILE,
	// * MOBILE_WEB,
	// * MOBILE_IOS,
	// * MOBILE_ANDROID,
	// * MOBILE_WINDOWS,
	// * MOBILE_TABLET,
	// * MOBILE_PHONE
	PaymentChannel string `json:"paymentChannel"`

	// Üye işyeri tarafından ilgili ödemenin sepetini tanımlamak için kullanılan id'dir.
	// Sipariş numarası ya da anlamlı bir değer olabilir.
	BasketId string `json:"basketId"`

	// Ödeme grubu, varsayılan PRODUCT. Geçerli değerler enum içinde sunulmaktadır:
	// * PRODUCT,
	// * LISTING,
	// * SUBSCRIPTION,
	// * OTHER.
	PaymentGroup string `json:"paymentGroup"`

	// Ödeme Kartı
	PaymentCard CreatePayment_PaymentCard `json:"paymentCard"`

	// Alıcı
	Buyer CreatePayment_Buyer `json:"buyer"`

	// Sevk Adresi
	ShippingAddress CreatePayment_ShippingAddress `json:"shippingAddress"`

	// Fatura Adresi
	BillingAddress CreatePayment_BillingAddress `json:"billingAddress"`

	// Sepetteki ürünler
	BasketItems []CreatePayment_BasketItems `json:"basketItems"`

	// Döviz Kodu
	Currency string `json:"currency"`

	// callbackUrl
	// 3D Secure ödeme akışında üye işyerine başarılı ve hatalı sonucu bildirmek üzere alınan URL adresi.
	// Sadece 3D Secure ödemenin init3DS metodunda zorunludur.
	CallbackURL string `json:"callbackUrl"`

	// PaymentSource string `json:"paymentSource"`
}

type CreatePayment_PaymentCard struct {
	// Ödemenin alınacağı kart sahibinin adı soyadı. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur.
	CardHolderName string `json:"cardHolderName"`

	// Ödemenin alınacağı kart numarası. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur.
	CardNumber string `json:"cardNumber"`

	// Ödemenin alınacağı kartın son kullanma tarihi yılı. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur.
	ExpireYear string `json:"expireYear"`

	// Ödemenin alınacağı kartın son kullanma tarihi ayı. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur.
	ExpireMonth string `json:"expireMonth"`

	// Ödeme esnasında kartın kaydedilip kaydedilmeyeceğini belirleyen parametre.
	// Varsayılan değeri 0 olup, geçerli değerler 0 ve 1’dir.
	RegisterCard *int `json:"registerCard,omitempty"`

	// Ödemenin alınacağı kartın güvenlik kodu. Eğer saklı kart ile ödeme yapılmıyorsa zorunludur.
	// Saklı kartla ödeme yapılırken gönderilirse aynen bankaya iletilir.
	Cvc string `json:"cvc"`
}

type CreatePayment_Buyer struct {
	// Üye işyeri tarafındaki alıcıya ait id.
	Id string `json:"id"`

	// Üye işyeri tarafındaki alıcıya ait ad.
	Name string `json:"name"`

	// Üye işyeri tarafındaki alıcıya ait soyad.
	Surname string `json:"surname"`

	// Üye işyeri tarafındaki alıcıya ait kimlik (TCKN) numarası.
	IdentityNumber string `json:"identityNumber"`

	// Üye işyeri tarafındaki alıcıya ait e-posta bilgisi. E-posta adresi alıcıya ait geçerli ve erişilebilir bir
	// adres olmalıdır.
	Email string `json:"email"`

	// Üye işyeri tarafındaki alıcıya ait GSM numarası.
	GsmNumber string `json:"gsmNumber"`

	// Üye işyeri tarafındaki alıcıya ait kayıt tarihi. Tarih formatı 2015-09-17 23:45:06 şeklinde olmalıdır.
	RegistrationDate string `json:"registrationDate"`

	// Üye işyeri tarafındaki alıcıya ait son giriş tarihi. Tarih formatı 2015-09-17 23:45:06 şeklinde olmalıdır.
	LastLoginDate string `json:"lastLoginDate"`

	// Üye işyeri tarafındaki alıcıya ait kayıt adresi.
	RegistrationAddress string `json:"registrationAddress"`

	// Üye işyeri tarafındaki alıcıya ait kayıt adresi.
	City string `json:"city"`

	// Üye işyeri tarafındaki fatura adresi ülke bilgisi.
	Country string `json:"country"`

	// Üye işyeri tarafındaki fatura adresi posta kodu.
	ZipCode string `json:"zipCode"`

	// Üye işyeri tarafındaki alıcıya ait IP adresi.
	Ip string `json:"ip"`
}

type CreatePayment_ShippingAddress struct {
	// Üye işyeri tarafındaki teslimat adresi. Sepetteki ürünlerden en az 1 tanesi fiziksel ürün
	// (itemType=PHYSICAL) ise zorunludur.
	Address string `json:"address"`

	// Üye işyeri tarafındaki teslimat adresi posta kodu.
	ZipCode string `json:"zipCode"`

	// Üye işyeri tarafındaki teslimat adresi, ad soyad bilgisi. Sepetteki ürünlerden en az 1 tanesi fiziksel ürün
	// (itemType=PHYSICAL) ise zorunludur.
	ContactName string `json:"contactName"`

	// Üye işyeri tarafındaki teslimat adresi şehir bilgisi. Sepetteki ürünlerden en az 1 tanesi fiziksel ürün
	// (itemType=PHYSICAL) ise zorunludur.
	City string `json:"city"`

	// Üye işyeri tarafındaki teslimat adresi ülke bilgisi. Sepetteki ürünlerden en az 1 tanesi fiziksel ürün
	// (itemType=PHYSICAL) ise zorunludur.
	Country string `json:"country"`
}

type CreatePayment_BillingAddress struct {
	// Üye işyeri tarafındaki fatura adresi ülke bilgisi.
	Address string `json:"address"`

	// Üye işyeri tarafındaki fatura adresi posta kodu.
	ZipCode string `json:"zipCode"`

	// Üye işyeri tarafındaki fatura adresi, ad soyad bilgisi.
	ContactName string `json:"contactName"`

	// Üye işyeri tarafındaki fatura adresi şehir bilgisi.
	City string `json:"city"`

	// Üye işyeri tarafındaki fatura adresi ülke bilgisi.
	Country string `json:"country"`
}

type CreatePayment_BasketItems struct {
	// Üye işyeri tarafındaki sepetteki ürüne ait id.
	Id string `json:"id"`

	// Üye işyeri tarafındaki sepetteki ürüne ait tutar. 0 ve 0’dan küçük olamaz; tutarlar toplamı sepet
	// tutarına (price) eşit olmalıdır.
	Price string `json:"price"`

	// Üye işyeri tarafındaki sepetteki ürüne ait isim.
	Name string `json:"name"`

	// Üye işyeri tarafındaki sepetteki ürüne ait kategori 1.
	Category1 string `json:"category1"`

	// Üye işyeri tarafındaki sepetteki ürüne ait kategori 2.
	Category2 string `json:"category2"`

	// Üye işyeri tarafındaki sepetteki ürüne ait tip. Geçerli enum değerler:
	//
	// PHYSICAL ve
	//
	// VIRTUAL.
	ItemType string `json:"itemType"`

	// SubMerchantKey
	SubMerchantKey string `json:"subMerchantKey,omitempty"`

	// SubMerchantPrice
	SubMerchantPrice string `json:"subMerchantPrice,omitempty"`
}
