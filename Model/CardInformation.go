/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:21
* Developer : ibrahimcobani
*
*******/
package Model

type CardInformation struct {
	CardAlias      string `json:"card_alias"`
	CardNumber     string `json:"card_number"`
	ExpireYear     string `json:"expire_year"`
	ExpireMonth    string `json:"expire_month"`
	CardHolderName string `json:"card_holder_name"`
}
