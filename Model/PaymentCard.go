/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:40
* Developer : ibrahimcobani
*
*******/
package Model

type PaymentCard struct {
	CardHolderName string `json:"card_holder_name"`
	CardNumber     string `json:"card_number"`
	ExpireYear     string `json:"expire_year"`
	ExpireMonth    string `json:"expire_month"`
	Cvc            string `json:"cvc"`
	RegisterCard   int `json:"register_card"`
	CardAlias      string `json:"card_alias"`
	CardToken      string `json:"card_token"`
	CardUserKey    string `json:"card_user_key"`
}
