/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:22
* Developer : ibrahimcobani
*
*******/
package Model

type CardList struct {
	CardUserKey string `json:"card_user_key"`
	CardDetails [] Card `json:"card_details"`
}
