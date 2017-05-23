/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 21:04
* Developer : ibrahimcobani
*
*******/
package Model

type SubMerchant struct {
	Name                  string `json:"name"`
	Email                 string `json:"email"`
	GsmNumber             string `json:"gsm_number"`
	Address               string `json:"address"`
	Iban                  string `json:"iban"`
	SwiftCode             string `json:"swift_code"`
	Currency              string `json:"currency"`
	TaxOffice             string `json:"tax_office"`
	ContactName           string `json:"contact_name"`
	ContactSurname        string `json:"contact_surname"`
	LegalCompanyTitle     string `json:"legal_company_title"`
	SubMerchantExternalId string `json:"sub_merchant_external_id"`
	IdentityNumber        string `json:"identity_number"`
	TaxNumber             string `json:"tax_number"`
	SubMerchantType       string `json:"sub_merchant_type"`
	SubMerchantKey        string `json:"sub_merchant_key"`
}
