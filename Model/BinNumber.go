/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:09
* Developer : ibrahimcobani
*
*******/
package Model

type BinNumber struct {
	Bin             string `json:"bin"`
	CardType        string `json:"card_type"`
	CardAssociation string `json:"card_association"`
	CardFamily      string `json:"card_family"`
	BankName        string `json:"bank_name"`
	BankCode        int32 `json:"bank_code"`
}
