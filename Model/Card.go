/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:19
* Developer : ibrahimcobani
*
*******/
package Model

type Card struct {
	ExternalId      string `json:"external_id"`
	Email           string `json:"email"`
	CardUserKey     string `json:"card_user_key"`
	CardToken       string `json:"card_token"`
	CardAlias       string `json:"card_alias"`
	BinNumber       string `json:"bin_number"`
	CardType        string `json:"card_type"`
	CardAssociation string `json:"card_association"`
	CardFamily      string `json:"card_family"`
	CardBankCode    int32 `json:"card_bank_code"`
	CardBankName    string `json:"card_bank_name"`
}
