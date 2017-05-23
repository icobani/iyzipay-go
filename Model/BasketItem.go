/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:03
* Developer : ibrahimcobani
*
*******/
package Model

type BasketItem struct {
	Id               string `json:"id"`
	Price            string `json:"price"`
	Name             string `json:"name"`
	Category1        string `json:"category_1"`
	Category2        string `json:"category_2"`
	ItemType         string `json:"item_type"`
	SubMerchantKey   string `json:"sub_merchant_key"`
	SubMerchantPrice string `json:"sub_merchant_price"`
}
