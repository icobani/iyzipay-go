/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 23/05/2017    
* Time      : 20:25
* Developer : ibrahimcobani
*
*******/
package Model

type CheckoutForm struct {
	Token       string `json:"token"`
	CallbackUrl string `json:"callback_url"`
}
