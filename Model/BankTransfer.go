/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 19:41
* Developer : ibrahimcobani
*
*******/
package Model

type BankTransfer struct {
	SubMerchantKey string `json:"sub_merchant_key"`
	Iban string `json:"iban"`
	ContactName string `json:"contact_name"`
	ContactSurname string `json:"contact_surname"`
	LegalCompanyTitle string `json:"legal_company_title"`
	MarketplaceSubMerchantType string `json:"marketplace_sub_merchant_type"`
}
