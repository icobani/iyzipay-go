/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 19:57
* Developer : ibrahimcobani
*
*******/
package Model

type BasicBkm struct {
	Token string `json:"token"`
	CallbackUrl string `json:"callback_url"`
	PaymentStatus string `json:"payment_status"`
}
